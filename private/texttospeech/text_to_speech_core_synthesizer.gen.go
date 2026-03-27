// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/avfaudio"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TextToSpeechCoreSynthesizer] class.
var (
	_TextToSpeechCoreSynthesizerClass     TextToSpeechCoreSynthesizerClass
	_TextToSpeechCoreSynthesizerClassOnce sync.Once
)

func getTextToSpeechCoreSynthesizerClass() TextToSpeechCoreSynthesizerClass {
	_TextToSpeechCoreSynthesizerClassOnce.Do(func() {
		_TextToSpeechCoreSynthesizerClass = TextToSpeechCoreSynthesizerClass{class: objc.GetClass("TextToSpeech.CoreSynthesizer")}
	})
	return _TextToSpeechCoreSynthesizerClass
}

// GetTextToSpeechCoreSynthesizerClass returns the class object for TextToSpeech.CoreSynthesizer.
func GetTextToSpeechCoreSynthesizerClass() TextToSpeechCoreSynthesizerClass {
	return getTextToSpeechCoreSynthesizerClass()
}

type TextToSpeechCoreSynthesizerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechCoreSynthesizerClass) Alloc() TextToSpeechCoreSynthesizer {
	rv := objc.Send[TextToSpeechCoreSynthesizer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TextToSpeechCoreSynthesizer._audioQueue]
//   - [TextToSpeechCoreSynthesizer.Set_audioQueue]
//   - [TextToSpeechCoreSynthesizer._bundleIdentifier]
//   - [TextToSpeechCoreSynthesizer.Set_bundleIdentifier]
//   - [TextToSpeechCoreSynthesizer._effects]
//   - [TextToSpeechCoreSynthesizer.Set_effects]
//   - [TextToSpeechCoreSynthesizer._voiceResolver]
//   - [TextToSpeechCoreSynthesizer.Set_voiceResolver]
//   - [TextToSpeechCoreSynthesizer.AudioDevice]
//   - [TextToSpeechCoreSynthesizer.SetAudioDevice]
//   - [TextToSpeechCoreSynthesizer.AudioQueue]
//   - [TextToSpeechCoreSynthesizer.SetAudioQueue]
//   - [TextToSpeechCoreSynthesizer.AudioQueueFlags]
//   - [TextToSpeechCoreSynthesizer.SetAudioQueueFlags]
//   - [TextToSpeechCoreSynthesizer.BundleIdentifier]
//   - [TextToSpeechCoreSynthesizer.SetBundleIdentifier]
//   - [TextToSpeechCoreSynthesizer.ContinueSpeaking]
//   - [TextToSpeechCoreSynthesizer.Effects]
//   - [TextToSpeechCoreSynthesizer.SetEffects]
//   - [TextToSpeechCoreSynthesizer.IsPaused]
//   - [TextToSpeechCoreSynthesizer.IsSpeaking]
//   - [TextToSpeechCoreSynthesizer.OfflineChain]
//   - [TextToSpeechCoreSynthesizer.SetOfflineChain]
//   - [TextToSpeechCoreSynthesizer.PauseSpeakingAtCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.SetLegacySubstitutions]
//   - [TextToSpeechCoreSynthesizer.SpeakSynthCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.SpeakWithRequestLanguageSynthesizerCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.StopSpeakingAtCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.StopWithCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.VoiceResolver]
//   - [TextToSpeechCoreSynthesizer.SetVoiceResolver]
//   - [TextToSpeechCoreSynthesizer.VoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.VoiceWithLocaleCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.WriteToBufferCallbackSynthCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.WriteToBufferCallbackToMarkerCallbackSynthCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.WriteWithSpeechPhraseToAudioFileWithAudioSettingsCompletionHandler]
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer
type TextToSpeechCoreSynthesizer struct {
	objectivec.Object
}

// TextToSpeechCoreSynthesizerFromID constructs a [TextToSpeechCoreSynthesizer] from an objc.ID.
func TextToSpeechCoreSynthesizerFromID(id objc.ID) TextToSpeechCoreSynthesizer {
	return TextToSpeechCoreSynthesizer{objectivec.Object{ID: id}}
}
// Ensure TextToSpeechCoreSynthesizer implements ITextToSpeechCoreSynthesizer.
var _ ITextToSpeechCoreSynthesizer = TextToSpeechCoreSynthesizer{}

// An interface definition for the [TextToSpeechCoreSynthesizer] class.
//
// # Methods
//
//   - [ITextToSpeechCoreSynthesizer._audioQueue]
//   - [ITextToSpeechCoreSynthesizer.Set_audioQueue]
//   - [ITextToSpeechCoreSynthesizer._bundleIdentifier]
//   - [ITextToSpeechCoreSynthesizer.Set_bundleIdentifier]
//   - [ITextToSpeechCoreSynthesizer._effects]
//   - [ITextToSpeechCoreSynthesizer.Set_effects]
//   - [ITextToSpeechCoreSynthesizer._voiceResolver]
//   - [ITextToSpeechCoreSynthesizer.Set_voiceResolver]
//   - [ITextToSpeechCoreSynthesizer.AudioDevice]
//   - [ITextToSpeechCoreSynthesizer.SetAudioDevice]
//   - [ITextToSpeechCoreSynthesizer.AudioQueue]
//   - [ITextToSpeechCoreSynthesizer.SetAudioQueue]
//   - [ITextToSpeechCoreSynthesizer.AudioQueueFlags]
//   - [ITextToSpeechCoreSynthesizer.SetAudioQueueFlags]
//   - [ITextToSpeechCoreSynthesizer.BundleIdentifier]
//   - [ITextToSpeechCoreSynthesizer.SetBundleIdentifier]
//   - [ITextToSpeechCoreSynthesizer.ContinueSpeaking]
//   - [ITextToSpeechCoreSynthesizer.Effects]
//   - [ITextToSpeechCoreSynthesizer.SetEffects]
//   - [ITextToSpeechCoreSynthesizer.IsPaused]
//   - [ITextToSpeechCoreSynthesizer.IsSpeaking]
//   - [ITextToSpeechCoreSynthesizer.OfflineChain]
//   - [ITextToSpeechCoreSynthesizer.SetOfflineChain]
//   - [ITextToSpeechCoreSynthesizer.PauseSpeakingAtCompletionHandler]
//   - [ITextToSpeechCoreSynthesizer.SetLegacySubstitutions]
//   - [ITextToSpeechCoreSynthesizer.SpeakSynthCompletionHandler]
//   - [ITextToSpeechCoreSynthesizer.SpeakWithRequestLanguageSynthesizerCompletionHandler]
//   - [ITextToSpeechCoreSynthesizer.StopSpeakingAtCompletionHandler]
//   - [ITextToSpeechCoreSynthesizer.StopWithCompletionHandler]
//   - [ITextToSpeechCoreSynthesizer.VoiceResolver]
//   - [ITextToSpeechCoreSynthesizer.SetVoiceResolver]
//   - [ITextToSpeechCoreSynthesizer.VoiceWithIdentifierCompletionHandler]
//   - [ITextToSpeechCoreSynthesizer.VoiceWithLocaleCompletionHandler]
//   - [ITextToSpeechCoreSynthesizer.WriteToBufferCallbackSynthCompletionHandler]
//   - [ITextToSpeechCoreSynthesizer.WriteToBufferCallbackToMarkerCallbackSynthCompletionHandler]
//   - [ITextToSpeechCoreSynthesizer.WriteWithSpeechPhraseToAudioFileWithAudioSettingsCompletionHandler]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer
type ITextToSpeechCoreSynthesizer interface {
	objectivec.IObject

	// Topic: Methods

	_audioQueue() ITTSWrappedAudioQueue
	Set_audioQueue(value ITTSWrappedAudioQueue)
	_bundleIdentifier() string
	Set_bundleIdentifier(value string)
	_effects() foundation.INSArray
	Set_effects(value foundation.INSArray)
	_voiceResolver() unsafe.Pointer
	Set_voiceResolver(value unsafe.Pointer)
	AudioDevice() uint32
	SetAudioDevice(value uint32)
	AudioQueue() ITTSWrappedAudioQueue
	SetAudioQueue(value ITTSWrappedAudioQueue)
	AudioQueueFlags() uint32
	SetAudioQueueFlags(value uint32)
	BundleIdentifier() string
	SetBundleIdentifier(value string)
	ContinueSpeaking() bool
	Effects() foundation.INSArray
	SetEffects(value foundation.INSArray)
	IsPaused() bool
	IsSpeaking() bool
	OfflineChain() foundation.INSArray
	SetOfflineChain(value foundation.INSArray)
	PauseSpeakingAtCompletionHandler(at int64, handler BoolErrorHandler)
	SetLegacySubstitutions(substitutions objectivec.IObject)
	SpeakSynthCompletionHandler(speak avfaudio.AVSpeechUtterance, synth avfaudio.AVSpeechSynthesizer, handler BoolErrorHandler)
	SpeakWithRequestLanguageSynthesizerCompletionHandler(request objectivec.IObject, language objectivec.IObject, synthesizer objectivec.IObject, handler BoolErrorHandler)
	StopSpeakingAtCompletionHandler(at int64, handler BoolErrorHandler)
	StopWithCompletionHandler(handler BoolErrorHandler)
	VoiceResolver() unsafe.Pointer
	SetVoiceResolver(value unsafe.Pointer)
	VoiceWithIdentifierCompletionHandler(identifier string, handler BoolErrorHandler)
	VoiceWithLocaleCompletionHandler(locale foundation.NSLocale, handler BoolErrorHandler)
	WriteToBufferCallbackSynthCompletionHandler(write avfaudio.AVSpeechUtterance, callback BoolErrorHandler, synth objectivec.IObject, handler ErrorHandler)
	WriteToBufferCallbackToMarkerCallbackSynthCompletionHandler(write avfaudio.AVSpeechUtterance, callback BoolErrorHandler, callback2 objectivec.IObject, synth objectivec.IObject, handler BoolErrorHandler)
	WriteWithSpeechPhraseToAudioFileWithAudioSettingsCompletionHandler(phrase string, file foundation.INSURL, settings foundation.INSDictionary, handler BoolErrorHandler)
}

// Init initializes the instance.
func (t TextToSpeechCoreSynthesizer) Init() TextToSpeechCoreSynthesizer {
	rv := objc.Send[TextToSpeechCoreSynthesizer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechCoreSynthesizer) Autorelease() TextToSpeechCoreSynthesizer {
	rv := objc.Send[TextToSpeechCoreSynthesizer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechCoreSynthesizer creates a new TextToSpeechCoreSynthesizer instance.
func NewTextToSpeechCoreSynthesizer() TextToSpeechCoreSynthesizer {
	class := getTextToSpeechCoreSynthesizerClass()
	rv := objc.Send[TextToSpeechCoreSynthesizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/continueSpeaking
func (t TextToSpeechCoreSynthesizer) ContinueSpeaking() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("continueSpeaking"))
	return rv
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/pauseSpeakingAt:completionHandler:
func (t TextToSpeechCoreSynthesizer) PauseSpeakingAtCompletionHandler(at int64, handler BoolErrorHandler) {
_block1, _cleanup1 := NewBoolErrorBlock(handler)
	defer _cleanup1()
	objc.Send[objc.ID](t.ID, objc.Sel("pauseSpeakingAt:completionHandler:"), at, _block1)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/setLegacySubstitutions:
func (t TextToSpeechCoreSynthesizer) SetLegacySubstitutions(substitutions objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("setLegacySubstitutions:"), substitutions)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/speak:synth:completionHandler:
func (t TextToSpeechCoreSynthesizer) SpeakSynthCompletionHandler(speak avfaudio.AVSpeechUtterance, synth avfaudio.AVSpeechSynthesizer, handler BoolErrorHandler) {
_block2, _cleanup2 := NewBoolErrorBlock(handler)
	defer _cleanup2()
	objc.Send[objc.ID](t.ID, objc.Sel("speak:synth:completionHandler:"), speak, synth, _block2)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/speakWithRequest:language:synthesizer:completionHandler:
func (t TextToSpeechCoreSynthesizer) SpeakWithRequestLanguageSynthesizerCompletionHandler(request objectivec.IObject, language objectivec.IObject, synthesizer objectivec.IObject, handler BoolErrorHandler) {
_block3, _cleanup3 := NewBoolErrorBlock(handler)
	defer _cleanup3()
	objc.Send[objc.ID](t.ID, objc.Sel("speakWithRequest:language:synthesizer:completionHandler:"), request, language, synthesizer, _block3)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/stopSpeakingAt:completionHandler:
func (t TextToSpeechCoreSynthesizer) StopSpeakingAtCompletionHandler(at int64, handler BoolErrorHandler) {
_block1, _cleanup1 := NewBoolErrorBlock(handler)
	defer _cleanup1()
	objc.Send[objc.ID](t.ID, objc.Sel("stopSpeakingAt:completionHandler:"), at, _block1)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/stopWithCompletionHandler:
func (t TextToSpeechCoreSynthesizer) StopWithCompletionHandler(handler BoolErrorHandler) {
_block0, _cleanup0 := NewBoolErrorBlock(handler)
	defer _cleanup0()
	objc.Send[objc.ID](t.ID, objc.Sel("stopWithCompletionHandler:"), _block0)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/voiceWithIdentifier:completionHandler:
func (t TextToSpeechCoreSynthesizer) VoiceWithIdentifierCompletionHandler(identifier string, handler BoolErrorHandler) {
_block1, _cleanup1 := NewBoolErrorBlock(handler)
	defer _cleanup1()
	objc.Send[objc.ID](t.ID, objc.Sel("voiceWithIdentifier:completionHandler:"), objc.String(identifier), _block1)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/voiceWithLocale:completionHandler:
func (t TextToSpeechCoreSynthesizer) VoiceWithLocaleCompletionHandler(locale foundation.NSLocale, handler BoolErrorHandler) {
_block1, _cleanup1 := NewBoolErrorBlock(handler)
	defer _cleanup1()
	objc.Send[objc.ID](t.ID, objc.Sel("voiceWithLocale:completionHandler:"), locale, _block1)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/write:toBufferCallback:synth:completionHandler:
func (t TextToSpeechCoreSynthesizer) WriteToBufferCallbackSynthCompletionHandler(write avfaudio.AVSpeechUtterance, callback BoolErrorHandler, synth objectivec.IObject, handler ErrorHandler) {
_block1, _cleanup1 := NewBoolErrorBlock(callback)
	defer _cleanup1()
	_block3, _cleanup3 := NewErrorBlock(handler)
	defer _cleanup3()
	objc.Send[objc.ID](t.ID, objc.Sel("write:toBufferCallback:synth:completionHandler:"), write, _block1, synth, _block3)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/write:toBufferCallback:toMarkerCallback:synth:completionHandler:
func (t TextToSpeechCoreSynthesizer) WriteToBufferCallbackToMarkerCallbackSynthCompletionHandler(write avfaudio.AVSpeechUtterance, callback BoolErrorHandler, callback2 objectivec.IObject, synth objectivec.IObject, handler BoolErrorHandler) {
_block1, _cleanup1 := NewBoolErrorBlock(callback)
	defer _cleanup1()
	_block4, _cleanup4 := NewBoolErrorBlock(handler)
	defer _cleanup4()
	objc.Send[objc.ID](t.ID, objc.Sel("write:toBufferCallback:toMarkerCallback:synth:completionHandler:"), write, _block1, callback2, synth, _block4)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/writeWithSpeechPhrase:toAudioFile:withAudioSettings:completionHandler:
func (t TextToSpeechCoreSynthesizer) WriteWithSpeechPhraseToAudioFileWithAudioSettingsCompletionHandler(phrase string, file foundation.INSURL, settings foundation.INSDictionary, handler BoolErrorHandler) {
_block3, _cleanup3 := NewBoolErrorBlock(handler)
	defer _cleanup3()
	objc.Send[objc.ID](t.ID, objc.Sel("writeWithSpeechPhrase:toAudioFile:withAudioSettings:completionHandler:"), objc.String(phrase), file, settings, _block3)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/_audioQueue
func (t TextToSpeechCoreSynthesizer) _audioQueue() ITTSWrappedAudioQueue {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("_audioQueue"))
	return TTSWrappedAudioQueueFromID(objc.ID(rv))
}
func (t TextToSpeechCoreSynthesizer) Set_audioQueue(value ITTSWrappedAudioQueue) {
	objc.Send[struct{}](t.ID, objc.Sel("set_audioQueue:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/_bundleIdentifier
func (t TextToSpeechCoreSynthesizer) _bundleIdentifier() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("_bundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (t TextToSpeechCoreSynthesizer) Set_bundleIdentifier(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("set_bundleIdentifier:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/_effects
func (t TextToSpeechCoreSynthesizer) _effects() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("_effects"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TextToSpeechCoreSynthesizer) Set_effects(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("set_effects:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/_voiceResolver
func (t TextToSpeechCoreSynthesizer) _voiceResolver() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t.ID, objc.Sel("_voiceResolver"))
	return rv
}
func (t TextToSpeechCoreSynthesizer) Set_voiceResolver(value unsafe.Pointer) {
	objc.Send[struct{}](t.ID, objc.Sel("set_voiceResolver:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/audioDevice
func (t TextToSpeechCoreSynthesizer) AudioDevice() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("audioDevice"))
	return rv
}
func (t TextToSpeechCoreSynthesizer) SetAudioDevice(value uint32) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioDevice:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/audioQueue
func (t TextToSpeechCoreSynthesizer) AudioQueue() ITTSWrappedAudioQueue {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("audioQueue"))
	return TTSWrappedAudioQueueFromID(objc.ID(rv))
}
func (t TextToSpeechCoreSynthesizer) SetAudioQueue(value ITTSWrappedAudioQueue) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioQueue:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/audioQueueFlags
func (t TextToSpeechCoreSynthesizer) AudioQueueFlags() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("audioQueueFlags"))
	return rv
}
func (t TextToSpeechCoreSynthesizer) SetAudioQueueFlags(value uint32) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioQueueFlags:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/bundleIdentifier
func (t TextToSpeechCoreSynthesizer) BundleIdentifier() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("bundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (t TextToSpeechCoreSynthesizer) SetBundleIdentifier(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setBundleIdentifier:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/effects
func (t TextToSpeechCoreSynthesizer) Effects() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("effects"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TextToSpeechCoreSynthesizer) SetEffects(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setEffects:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/isPaused
func (t TextToSpeechCoreSynthesizer) IsPaused() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isPaused"))
	return rv
}
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/isSpeaking
func (t TextToSpeechCoreSynthesizer) IsSpeaking() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isSpeaking"))
	return rv
}
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/offlineChain
func (t TextToSpeechCoreSynthesizer) OfflineChain() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("offlineChain"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TextToSpeechCoreSynthesizer) SetOfflineChain(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setOfflineChain:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TextToSpeech.CoreSynthesizer/voiceResolver
func (t TextToSpeechCoreSynthesizer) VoiceResolver() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t.ID, objc.Sel("voiceResolver"))
	return rv
}
func (t TextToSpeechCoreSynthesizer) SetVoiceResolver(value unsafe.Pointer) {
	objc.Send[struct{}](t.ID, objc.Sel("setVoiceResolver:"), value)
}

// PauseSpeakingAt is a synchronous wrapper around [TextToSpeechCoreSynthesizer.PauseSpeakingAtCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesizer) PauseSpeakingAt(ctx context.Context, at int64) (bool, error) {
	type result struct {
		val bool
		err error
	}
	done := make(chan result, 1)
	t.PauseSpeakingAtCompletionHandler(at, func(val bool, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// SpeakSynth is a synchronous wrapper around [TextToSpeechCoreSynthesizer.SpeakSynthCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesizer) SpeakSynth(ctx context.Context, speak avfaudio.AVSpeechUtterance, synth avfaudio.AVSpeechSynthesizer) (bool, error) {
	type result struct {
		val bool
		err error
	}
	done := make(chan result, 1)
	t.SpeakSynthCompletionHandler(speak, synth, func(val bool, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// SpeakWithRequestLanguageSynthesizer is a synchronous wrapper around [TextToSpeechCoreSynthesizer.SpeakWithRequestLanguageSynthesizerCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesizer) SpeakWithRequestLanguageSynthesizer(ctx context.Context, request objectivec.IObject, language objectivec.IObject, synthesizer objectivec.IObject) (bool, error) {
	type result struct {
		val bool
		err error
	}
	done := make(chan result, 1)
	t.SpeakWithRequestLanguageSynthesizerCompletionHandler(request, language, synthesizer, func(val bool, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// StopSpeakingAt is a synchronous wrapper around [TextToSpeechCoreSynthesizer.StopSpeakingAtCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesizer) StopSpeakingAt(ctx context.Context, at int64) (bool, error) {
	type result struct {
		val bool
		err error
	}
	done := make(chan result, 1)
	t.StopSpeakingAtCompletionHandler(at, func(val bool, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// Stop is a synchronous wrapper around [TextToSpeechCoreSynthesizer.StopWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesizer) Stop(ctx context.Context) (bool, error) {
	type result struct {
		val bool
		err error
	}
	done := make(chan result, 1)
	t.StopWithCompletionHandler(func(val bool, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// VoiceWithIdentifier is a synchronous wrapper around [TextToSpeechCoreSynthesizer.VoiceWithIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesizer) VoiceWithIdentifier(ctx context.Context, identifier string) (bool, error) {
	type result struct {
		val bool
		err error
	}
	done := make(chan result, 1)
	t.VoiceWithIdentifierCompletionHandler(identifier, func(val bool, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// VoiceWithLocale is a synchronous wrapper around [TextToSpeechCoreSynthesizer.VoiceWithLocaleCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesizer) VoiceWithLocale(ctx context.Context, locale foundation.NSLocale) (bool, error) {
	type result struct {
		val bool
		err error
	}
	done := make(chan result, 1)
	t.VoiceWithLocaleCompletionHandler(locale, func(val bool, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// WriteToBufferCallbackToMarkerCallbackSynth is a synchronous wrapper around [TextToSpeechCoreSynthesizer.WriteToBufferCallbackToMarkerCallbackSynthCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesizer) WriteToBufferCallbackToMarkerCallbackSynth(ctx context.Context, write avfaudio.AVSpeechUtterance, callback BoolErrorHandler, callback2 objectivec.IObject, synth objectivec.IObject) (bool, error) {
	type result struct {
		val bool
		err error
	}
	done := make(chan result, 1)
	t.WriteToBufferCallbackToMarkerCallbackSynthCompletionHandler(write, callback, callback2, synth, func(val bool, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// WriteWithSpeechPhraseToAudioFileWithAudioSettings is a synchronous wrapper around [TextToSpeechCoreSynthesizer.WriteWithSpeechPhraseToAudioFileWithAudioSettingsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesizer) WriteWithSpeechPhraseToAudioFileWithAudioSettings(ctx context.Context, phrase string, file foundation.INSURL, settings foundation.INSDictionary) (bool, error) {
	type result struct {
		val bool
		err error
	}
	done := make(chan result, 1)
	t.WriteWithSpeechPhraseToAudioFileWithAudioSettingsCompletionHandler(phrase, file, settings, func(val bool, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

