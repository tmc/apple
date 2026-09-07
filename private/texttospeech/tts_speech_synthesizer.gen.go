// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[TTSSpeechSynthesizer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

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
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
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
	StartSpeakingStringToURLWithLanguageCodeError(string_ objectivec.IObject, url foundation.NSURL, code objectivec.IObject) (bool, error)
	StartSpeakingStringToURLWithLanguageCodeRequestError(string_ objectivec.IObject, url foundation.NSURL, code objectivec.IObject, request []objectivec.IObject) (bool, error)
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
	rv := objc.SendIfResponds[TTSSpeechSynthesizer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSSpeechSynthesizer) Autorelease() TTSSpeechSynthesizer {
	rv := objc.SendIfResponds[TTSSpeechSynthesizer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSSpeechSynthesizer creates a new TTSSpeechSynthesizer instance.
func NewTTSSpeechSynthesizer() TTSSpeechSynthesizer {
	class := getTTSSpeechSynthesizerClass()
	rv := objc.SendIfResponds[TTSSpeechSynthesizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

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
func (t TTSSpeechSynthesizer) _makeRequestForVoiceAndLanguageCode(voice objectivec.IObject, code objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("_makeRequestForVoice:andLanguageCode:"), voice, code)
	return objectivec.Object{ID: rv}
}

// MakeRequestForVoiceAndLanguageCode is an exported wrapper for the private method _makeRequestForVoiceAndLanguageCode.
func (t TTSSpeechSynthesizer) MakeRequestForVoiceAndLanguageCode(voice objectivec.IObject, code objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_makeRequestForVoice:andLanguageCode:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_makeRequestForVoice:andLanguageCode:"}
		return nil, err
	}
	return t._makeRequestForVoiceAndLanguageCode(voice, code), nil
}

// CanMakeRequestForVoiceAndLanguageCode reports whether the receiver responds to the private selector _makeRequestForVoice:andLanguageCode:.
func (t TTSSpeechSynthesizer) CanMakeRequestForVoiceAndLanguageCode() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_makeRequestForVoice:andLanguageCode:"))
}
func (t TTSSpeechSynthesizer) _mediaServicesDied() {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("_mediaServicesDied"))
}

// MediaServicesDied is an exported wrapper for the private method _mediaServicesDied.
func (t TTSSpeechSynthesizer) MediaServicesDied() error {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_mediaServicesDied")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_mediaServicesDied"}
		return err
	}
	t._mediaServicesDied()
	return nil
}

// CanMediaServicesDied reports whether the receiver responds to the private selector _mediaServicesDied.
func (t TTSSpeechSynthesizer) CanMediaServicesDied() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_mediaServicesDied"))
}
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
func (t TTSSpeechSynthesizer) _preprocessTextLanguageCode(text objectivec.IObject, code objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("_preprocessText:languageCode:"), text, code)
	return objectivec.Object{ID: rv}
}

// PreprocessTextLanguageCode is an exported wrapper for the private method _preprocessTextLanguageCode.
func (t TTSSpeechSynthesizer) PreprocessTextLanguageCode(text objectivec.IObject, code objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_preprocessText:languageCode:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_preprocessText:languageCode:"}
		return nil, err
	}
	return t._preprocessTextLanguageCode(text, code), nil
}

// CanPreprocessTextLanguageCode reports whether the receiver responds to the private selector _preprocessText:languageCode:.
func (t TTSSpeechSynthesizer) CanPreprocessTextLanguageCode() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_preprocessText:languageCode:"))
}
func (t TTSSpeechSynthesizer) _processMarkerForRequest(marker objectivec.IObject, request objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("_processMarker:forRequest:"), marker, request)
	return objectivec.Object{ID: rv}
}

// ProcessMarkerForRequest is an exported wrapper for the private method _processMarkerForRequest.
func (t TTSSpeechSynthesizer) ProcessMarkerForRequest(marker objectivec.IObject, request objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_processMarker:forRequest:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_processMarker:forRequest:"}
		return nil, err
	}
	return t._processMarkerForRequest(marker, request), nil
}

// CanProcessMarkerForRequest reports whether the receiver responds to the private selector _processMarker:forRequest:.
func (t TTSSpeechSynthesizer) CanProcessMarkerForRequest() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_processMarker:forRequest:"))
}
func (t TTSSpeechSynthesizer) _resolveVoiceForLanguage(language objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("_resolveVoiceForLanguage:"), language)
	return objectivec.Object{ID: rv}
}

// ResolveVoiceForLanguage is an exported wrapper for the private method _resolveVoiceForLanguage.
func (t TTSSpeechSynthesizer) ResolveVoiceForLanguage(language objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_resolveVoiceForLanguage:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_resolveVoiceForLanguage:"}
		return nil, err
	}
	return t._resolveVoiceForLanguage(language), nil
}

// CanResolveVoiceForLanguage reports whether the receiver responds to the private selector _resolveVoiceForLanguage:.
func (t TTSSpeechSynthesizer) CanResolveVoiceForLanguage() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_resolveVoiceForLanguage:"))
}
func (t TTSSpeechSynthesizer) _setDelegate(delegate objectivec.IObject) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("_setDelegate:"), delegate)
}
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
	if !objc.RespondsToSelector(t.ID, objc.Sel("_startSpeakingString:orSSMLString:withLanguageCode:jobId:request:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_startSpeakingString:orSSMLString:withLanguageCode:jobId:request:error:"}
		return false, err
	}
	return t._startSpeakingStringOrSSMLStringWithLanguageCodeJobIdRequestError(string_, sSMLString, code, id, request)
}

// CanStartSpeakingStringOrSSMLStringWithLanguageCodeJobIdRequestError reports whether the receiver responds to the private selector _startSpeakingString:orSSMLString:withLanguageCode:jobId:request:error:.
func (t TTSSpeechSynthesizer) CanStartSpeakingStringOrSSMLStringWithLanguageCodeJobIdRequestError() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_startSpeakingString:orSSMLString:withLanguageCode:jobId:request:error:"))
}
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
func (t TTSSpeechSynthesizer) Footprint() int64 {
	rv := objc.SendIfResponds[int64](t.ID, objc.Sel("footprint"))
	return rv
}
func (t TTSSpeechSynthesizer) GetPerVoiceSettings() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("getPerVoiceSettings"))
	return objectivec.Object{ID: rv}
}
func (t TTSSpeechSynthesizer) IsSpeaking() bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("isSpeaking"))
	return rv
}
func (t TTSSpeechSynthesizer) MaximumRate() float32 {
	rv := objc.SendIfResponds[float32](t.ID, objc.Sel("maximumRate"))
	return rv
}
func (t TTSSpeechSynthesizer) MinimumRate() float32 {
	rv := objc.SendIfResponds[float32](t.ID, objc.Sel("minimumRate"))
	return rv
}
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
func (t TTSSpeechSynthesizer) ResolvedVoiceIdentifierForLanguageCode(code objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("resolvedVoiceIdentifierForLanguageCode:"), code)
	return objectivec.Object{ID: rv}
}

var _ttsspeechsynthesizer_setaudiobuffercallback_p0_key byte

func (t TTSSpeechSynthesizer) SetAudioBufferCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("setAudioBufferCallback:"), _block0)
}
func (t TTSSpeechSynthesizer) SetFootprint(footprint int64) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("setFootprint:"), footprint)
}
func (t TTSSpeechSynthesizer) SpeechRequestDidStopWithSuccessPhonemesSpokenError(request objectivec.IObject, success bool, spoken objectivec.IObject, error_ objectivec.IObject) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("speechRequest:didStopWithSuccess:phonemesSpoken:error:"), request, success, spoken, error_)
}
func (t TTSSpeechSynthesizer) SpeechRequestWithMarker(request objectivec.IObject, marker objectivec.IObject) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("speechRequest:withMarker:"), request, marker)
}
func (t TTSSpeechSynthesizer) SpeechRequestDidContinue(continue_ objectivec.IObject) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("speechRequestDidContinue:"), continue_)
}
func (t TTSSpeechSynthesizer) SpeechRequestDidPause(pause objectivec.IObject) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("speechRequestDidPause:"), pause)
}
func (t TTSSpeechSynthesizer) SpeechRequestDidStart(start objectivec.IObject) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("speechRequestDidStart:"), start)
}
func (t TTSSpeechSynthesizer) SpeechString() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("speechString"))
	return objectivec.Object{ID: rv}
}
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
func (t TTSSpeechSynthesizer) StartSpeakingStringToURLWithLanguageCodeError(string_ objectivec.IObject, url foundation.NSURL, code objectivec.IObject) (bool, error) {
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
func (t TTSSpeechSynthesizer) StartSpeakingStringToURLWithLanguageCodeRequestError(string_ objectivec.IObject, url foundation.NSURL, code objectivec.IObject, request []objectivec.IObject) (bool, error) {
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
func (t TTSSpeechSynthesizer) TestingLastRuleConversion() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("testingLastRuleConversion"))
	return objectivec.Object{ID: rv}
}
func (t TTSSpeechSynthesizer) TestingSetLastRuleConversionReplacement(conversion objectivec.IObject, replacement objectivec.IObject) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("testingSetLastRuleConversion:replacement:"), conversion, replacement)
}
func (t TTSSpeechSynthesizer) UpdateCoreSynthSubstitutions() {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("updateCoreSynthSubstitutions"))
}
func (t TTSSpeechSynthesizer) UseSpecificAudioSession(session uint32) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("useSpecificAudioSession:"), session)
}
func (t TTSSpeechSynthesizer) VoiceResolver() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("voiceResolver"))
	return objectivec.Object{ID: rv}
}

func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) _speechVoiceForIdentifierLanguageFootprint(identifier objectivec.IObject, language objectivec.IObject, footprint int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("_speechVoiceForIdentifier:language:footprint:"), identifier, language, footprint)
	return objectivec.Object{ID: rv}
}

// SpeechVoiceForIdentifierLanguageFootprint is an exported wrapper for the private method _speechVoiceForIdentifierLanguageFootprint.
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) SpeechVoiceForIdentifierLanguageFootprint(identifier objectivec.IObject, language objectivec.IObject, footprint int64) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("_speechVoiceForIdentifier:language:footprint:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_speechVoiceForIdentifier:language:footprint:"}
		return nil, err
	}
	return _TTSSpeechSynthesizerClass._speechVoiceForIdentifierLanguageFootprint(identifier, language, footprint), nil
}

// CanSpeechVoiceForIdentifierLanguageFootprint reports whether the receiver responds to the private selector _speechVoiceForIdentifier:language:footprint:.
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) CanSpeechVoiceForIdentifierLanguageFootprint() bool {
	return objc.RespondsToSelector(objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("_speechVoiceForIdentifier:language:footprint:"))
}
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) AudioFileSettingsForVoice(voice objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("audioFileSettingsForVoice:"), voice)
	return objectivec.Object{ID: rv}
}
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) AvailableLanguageCodes() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("availableLanguageCodes"))
	return objectivec.Object{ID: rv}
}
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) CombinedProsodyMarkupForIdentifierStringRatePitchVolume(identifier objectivec.IObject, string_ objectivec.IObject, rate objectivec.IObject, pitch objectivec.IObject, volume objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("combinedProsodyMarkupForIdentifier:string:rate:pitch:volume:"), identifier, string_, rate, pitch, volume)
	return objectivec.Object{ID: rv}
}
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) EmploySpeechMarkupForTypeIdentifierWithLanguage(type_ int64, identifier objectivec.IObject, language objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("employSpeechMarkupForType:identifier:withLanguage:"), type_, identifier, language)
	return rv
}
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) GenericMarkMarkupForIdentifierName(identifier objectivec.IObject, name objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("genericMarkMarkupForIdentifier:name:"), identifier, name)
	return objectivec.Object{ID: rv}
}
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) IsSystemVoice(voice objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("isSystemVoice:"), voice)
	return rv
}
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) RemapVoiceIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("remapVoiceIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) SpeechMarkupStringForTypeForIdentifierString(type_ int64, identifier objectivec.IObject, string_ objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("speechMarkupStringForType:forIdentifier:string:"), type_, identifier, string_)
	return objectivec.Object{ID: rv}
}
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) SupportedIPAPhonemeLanguages() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("supportedIPAPhonemeLanguages"))
	return objectivec.Object{ID: rv}
}
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) UnavailableVoiceIdentifiers() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("unavailableVoiceIdentifiers"))
	return objectivec.Object{ID: rv}
}
func (_TTSSpeechSynthesizerClass TTSSpeechSynthesizerClass) VoiceForIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_TTSSpeechSynthesizerClass.class), objc.Sel("voiceForIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}

func (t TTSSpeechSynthesizer) AudioDeviceId() uint32 {
	rv := objc.SendIfResponds[uint32](t.ID, objc.Sel("audioDeviceId"))
	return rv
}
func (t TTSSpeechSynthesizer) SetAudioDeviceId(value uint32) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setAudioDeviceId:"), value)
}
func (t TTSSpeechSynthesizer) AudioEffects() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("audioEffects"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSSpeechSynthesizer) SetAudioEffects(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setAudioEffects:"), value)
}
func (t TTSSpeechSynthesizer) AudioQueueFlags() uint32 {
	rv := objc.SendIfResponds[uint32](t.ID, objc.Sel("audioQueueFlags"))
	return rv
}
func (t TTSSpeechSynthesizer) SetAudioQueueFlags(value uint32) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setAudioQueueFlags:"), value)
}
func (t TTSSpeechSynthesizer) BundleIdentifier() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("bundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechSynthesizer) SetBundleIdentifier(value string) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setBundleIdentifier:"), objc.String(value))
}
func (t TTSSpeechSynthesizer) CoreSynth() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](t.ID, objc.Sel("coreSynth"))
	return rv
}
func (t TTSSpeechSynthesizer) SetCoreSynth(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setCoreSynth:"), value)
}
func (t TTSSpeechSynthesizer) Delegate() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](t.ID, objc.Sel("delegate"))
	return rv
}
func (t TTSSpeechSynthesizer) SetDelegate(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setDelegate:"), value)
}
func (t TTSSpeechSynthesizer) DelegateTargetQueue() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("delegateTargetQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (t TTSSpeechSynthesizer) SetDelegateTargetQueue(value objectivec.Object) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setDelegateTargetQueue:"), value)
}
func (t TTSSpeechSynthesizer) IgnoreSubstitutions() bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("ignoreSubstitutions"))
	return rv
}
func (t TTSSpeechSynthesizer) SetIgnoreSubstitutions(value bool) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setIgnoreSubstitutions:"), value)
}
func (t TTSSpeechSynthesizer) NormalizedRate() float32 {
	rv := objc.SendIfResponds[float32](t.ID, objc.Sel("normalizedRate"))
	return rv
}
func (t TTSSpeechSynthesizer) SetNormalizedRate(value float32) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setNormalizedRate:"), value)
}
func (t TTSSpeechSynthesizer) OutputChannels() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("outputChannels"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSSpeechSynthesizer) SetOutputChannels(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setOutputChannels:"), value)
}
func (t TTSSpeechSynthesizer) PerVoiceSettings() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("perVoiceSettings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t TTSSpeechSynthesizer) SetPerVoiceSettings(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setPerVoiceSettings:"), value)
}
func (t TTSSpeechSynthesizer) PhonemeSubstitutions() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("phonemeSubstitutions"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSSpeechSynthesizer) SetPhonemeSubstitutions(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setPhonemeSubstitutions:"), value)
}
func (t TTSSpeechSynthesizer) Pitch() float32 {
	rv := objc.SendIfResponds[float32](t.ID, objc.Sel("pitch"))
	return rv
}
func (t TTSSpeechSynthesizer) SetPitch(value float32) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setPitch:"), value)
}
func (t TTSSpeechSynthesizer) Rate() float32 {
	rv := objc.SendIfResponds[float32](t.ID, objc.Sel("rate"))
	return rv
}
func (t TTSSpeechSynthesizer) SetRate(value float32) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setRate:"), value)
}
func (t TTSSpeechSynthesizer) RequestClientIdentifier() uint64 {
	rv := objc.SendIfResponds[uint64](t.ID, objc.Sel("requestClientIdentifier"))
	return rv
}
func (t TTSSpeechSynthesizer) SetRequestClientIdentifier(value uint64) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setRequestClientIdentifier:"), value)
}
func (t TTSSpeechSynthesizer) ResolvedVoiceIdentifier() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("resolvedVoiceIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechSynthesizer) SkipLuthorRules() bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("skipLuthorRules"))
	return rv
}
func (t TTSSpeechSynthesizer) SetSkipLuthorRules(value bool) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setSkipLuthorRules:"), value)
}
func (t TTSSpeechSynthesizer) SpeakingRequestClientContext() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](t.ID, objc.Sel("speakingRequestClientContext"))
	return rv
}
func (t TTSSpeechSynthesizer) SetSpeakingRequestClientContext(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setSpeakingRequestClientContext:"), value)
}
func (t TTSSpeechSynthesizer) SpeechSource() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("speechSource"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechSynthesizer) SetSpeechSource(value string) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setSpeechSource:"), objc.String(value))
}
func (t TTSSpeechSynthesizer) SupportsAccurateWordCallbacks() bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("supportsAccurateWordCallbacks"))
	return rv
}
func (t TTSSpeechSynthesizer) SetSupportsAccurateWordCallbacks(value bool) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setSupportsAccurateWordCallbacks:"), value)
}
func (t TTSSpeechSynthesizer) SynthesizeSilently() bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("synthesizeSilently"))
	return rv
}
func (t TTSSpeechSynthesizer) SetSynthesizeSilently(value bool) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setSynthesizeSilently:"), value)
}
func (t TTSSpeechSynthesizer) UserSubstitutions() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("userSubstitutions"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSSpeechSynthesizer) SetUserSubstitutions(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setUserSubstitutions:"), value)
}
func (t TTSSpeechSynthesizer) VoiceIdentifier() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("voiceIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechSynthesizer) SetVoiceIdentifier(value string) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setVoiceIdentifier:"), objc.String(value))
}
func (t TTSSpeechSynthesizer) Volume() float32 {
	rv := objc.SendIfResponds[float32](t.ID, objc.Sel("volume"))
	return rv
}
func (t TTSSpeechSynthesizer) SetVolume(value float32) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setVolume:"), value)
}
func (t TTSSpeechSynthesizer) Voucher() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("voucher"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (t TTSSpeechSynthesizer) SetVoucher(value objectivec.Object) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setVoucher:"), value)
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
