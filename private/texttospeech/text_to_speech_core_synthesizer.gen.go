// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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

// Class returns the underlying Objective-C class pointer.
func (tc TextToSpeechCoreSynthesizerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TextToSpeechCoreSynthesizerClass) Alloc() TextToSpeechCoreSynthesizer {
	rv := objc.SendIfResponds[TextToSpeechCoreSynthesizer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

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
//   - [TextToSpeechCoreSynthesizer.PauseSpeakingAt]
//   - [TextToSpeechCoreSynthesizer.SetLegacySubstitutions]
//   - [TextToSpeechCoreSynthesizer.SpeakSynth]
//   - [TextToSpeechCoreSynthesizer.SpeakWithRequestLanguageSynthesizerCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.StopSpeakingAt]
//   - [TextToSpeechCoreSynthesizer.StopWithCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.VoiceResolver]
//   - [TextToSpeechCoreSynthesizer.SetVoiceResolver]
//   - [TextToSpeechCoreSynthesizer.VoiceWithIdentifierCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.VoiceWithLocaleCompletionHandler]
//   - [TextToSpeechCoreSynthesizer.WriteToBufferCallbackSynth]
//   - [TextToSpeechCoreSynthesizer.WriteToBufferCallbackToMarkerCallbackSynth]
//   - [TextToSpeechCoreSynthesizer.WriteWithSpeechPhraseToAudioFileWithAudioSettingsCompletionHandler]
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
//   - [ITextToSpeechCoreSynthesizer.PauseSpeakingAt]
//   - [ITextToSpeechCoreSynthesizer.SetLegacySubstitutions]
//   - [ITextToSpeechCoreSynthesizer.SpeakSynth]
//   - [ITextToSpeechCoreSynthesizer.SpeakWithRequestLanguageSynthesizerCompletionHandler]
//   - [ITextToSpeechCoreSynthesizer.StopSpeakingAt]
//   - [ITextToSpeechCoreSynthesizer.StopWithCompletionHandler]
//   - [ITextToSpeechCoreSynthesizer.VoiceResolver]
//   - [ITextToSpeechCoreSynthesizer.SetVoiceResolver]
//   - [ITextToSpeechCoreSynthesizer.VoiceWithIdentifierCompletionHandler]
//   - [ITextToSpeechCoreSynthesizer.VoiceWithLocaleCompletionHandler]
//   - [ITextToSpeechCoreSynthesizer.WriteToBufferCallbackSynth]
//   - [ITextToSpeechCoreSynthesizer.WriteToBufferCallbackToMarkerCallbackSynth]
//   - [ITextToSpeechCoreSynthesizer.WriteWithSpeechPhraseToAudioFileWithAudioSettingsCompletionHandler]
type ITextToSpeechCoreSynthesizer interface {
	objectivec.IObject

	// Topic: Methods

	_audioQueue() unsafe.Pointer
	Set_audioQueue(value unsafe.Pointer)
	_bundleIdentifier() string
	Set_bundleIdentifier(value string)
	_effects() foundation.INSArray
	Set_effects(value foundation.INSArray)
	_voiceResolver() unsafe.Pointer
	Set_voiceResolver(value unsafe.Pointer)
	AudioDevice() uint32
	SetAudioDevice(value uint32)
	AudioQueue() unsafe.Pointer
	SetAudioQueue(value unsafe.Pointer)
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
	PauseSpeakingAt(at int64) bool
	SetLegacySubstitutions(substitutions objectivec.IObject)
	SpeakSynth(speak objectivec.IObject, synth objectivec.IObject)
	SpeakWithRequestLanguageSynthesizerCompletionHandler(request objectivec.IObject, language objectivec.IObject, synthesizer objectivec.IObject, handler ErrorHandler)
	StopSpeakingAt(at int64) bool
	StopWithCompletionHandler(handler ErrorHandler)
	VoiceResolver() unsafe.Pointer
	SetVoiceResolver(value unsafe.Pointer)
	VoiceWithIdentifierCompletionHandler(identifier string, handler ErrorHandler)
	VoiceWithLocaleCompletionHandler(locale foundation.NSLocale, handler ErrorHandler)
	WriteToBufferCallbackSynth(write objectivec.IObject, callback VoidHandler, synth objectivec.IObject)
	WriteToBufferCallbackToMarkerCallbackSynth(write objectivec.IObject, callback VoidHandler, callback2 VoidHandler, synth objectivec.IObject)
	WriteWithSpeechPhraseToAudioFileWithAudioSettingsCompletionHandler(phrase string, file foundation.NSURL, settings foundation.INSDictionary, handler ErrorHandler)
}

// Init initializes the instance.
func (t TextToSpeechCoreSynthesizer) Init() TextToSpeechCoreSynthesizer {
	rv := objc.SendIfResponds[TextToSpeechCoreSynthesizer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TextToSpeechCoreSynthesizer) Autorelease() TextToSpeechCoreSynthesizer {
	rv := objc.SendIfResponds[TextToSpeechCoreSynthesizer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextToSpeechCoreSynthesizer creates a new TextToSpeechCoreSynthesizer instance.
func NewTextToSpeechCoreSynthesizer() TextToSpeechCoreSynthesizer {
	class := getTextToSpeechCoreSynthesizerClass()
	rv := objc.SendIfResponds[TextToSpeechCoreSynthesizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (t TextToSpeechCoreSynthesizer) ContinueSpeaking() bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("continueSpeaking"))
	return rv
}
func (t TextToSpeechCoreSynthesizer) PauseSpeakingAt(at int64) bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("pauseSpeakingAt:"), at)
	return rv
}
func (t TextToSpeechCoreSynthesizer) SetLegacySubstitutions(substitutions objectivec.IObject) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("setLegacySubstitutions:"), substitutions)
}
func (t TextToSpeechCoreSynthesizer) SpeakSynth(speak objectivec.IObject, synth objectivec.IObject) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("speak:synth:"), speak, synth)
}
func (t TextToSpeechCoreSynthesizer) SpeakWithRequestLanguageSynthesizerCompletionHandler(request objectivec.IObject, language objectivec.IObject, synthesizer objectivec.IObject, handler ErrorHandler) {
	_block3, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("speakWithRequest:language:synthesizer:completionHandler:"), request, language, synthesizer, _block3)
}
func (t TextToSpeechCoreSynthesizer) StopSpeakingAt(at int64) bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("stopSpeakingAt:"), at)
	return rv
}
func (t TextToSpeechCoreSynthesizer) StopWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("stopWithCompletionHandler:"), _block0)
}
func (t TextToSpeechCoreSynthesizer) VoiceWithIdentifierCompletionHandler(identifier string, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("voiceWithIdentifier:completionHandler:"), objc.String(identifier), _block1)
}
func (t TextToSpeechCoreSynthesizer) VoiceWithLocaleCompletionHandler(locale foundation.NSLocale, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("voiceWithLocale:completionHandler:"), locale, _block1)
}
func (t TextToSpeechCoreSynthesizer) WriteToBufferCallbackSynth(write objectivec.IObject, callback VoidHandler, synth objectivec.IObject) {
	_block1, _ := NewVoidBlock(callback)
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("write:toBufferCallback:synth:"), write, _block1, synth)
}
func (t TextToSpeechCoreSynthesizer) WriteToBufferCallbackToMarkerCallbackSynth(write objectivec.IObject, callback VoidHandler, callback2 VoidHandler, synth objectivec.IObject) {
	_block1, _ := NewVoidBlock(callback)
	_block2, _ := NewVoidBlock(callback2)
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("write:toBufferCallback:toMarkerCallback:synth:"), write, _block1, _block2, synth)
}
func (t TextToSpeechCoreSynthesizer) WriteWithSpeechPhraseToAudioFileWithAudioSettingsCompletionHandler(phrase string, file foundation.NSURL, settings foundation.INSDictionary, handler ErrorHandler) {
	_block3, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("writeWithSpeechPhrase:toAudioFile:withAudioSettings:completionHandler:"), objc.String(phrase), file, settings, _block3)
}

func (t TextToSpeechCoreSynthesizer) _audioQueue() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](t.ID, objc.Sel("_audioQueue"))
	return rv
}
func (t TextToSpeechCoreSynthesizer) Set_audioQueue(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("set_audioQueue:"), value)
}
func (t TextToSpeechCoreSynthesizer) _bundleIdentifier() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("_bundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (t TextToSpeechCoreSynthesizer) Set_bundleIdentifier(value string) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("set_bundleIdentifier:"), objc.String(value))
}
func (t TextToSpeechCoreSynthesizer) _effects() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("_effects"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TextToSpeechCoreSynthesizer) Set_effects(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("set_effects:"), value)
}
func (t TextToSpeechCoreSynthesizer) _voiceResolver() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](t.ID, objc.Sel("_voiceResolver"))
	return rv
}
func (t TextToSpeechCoreSynthesizer) Set_voiceResolver(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("set_voiceResolver:"), value)
}
func (t TextToSpeechCoreSynthesizer) AudioDevice() uint32 {
	rv := objc.SendIfResponds[uint32](t.ID, objc.Sel("audioDevice"))
	return rv
}
func (t TextToSpeechCoreSynthesizer) SetAudioDevice(value uint32) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setAudioDevice:"), value)
}
func (t TextToSpeechCoreSynthesizer) AudioQueue() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](t.ID, objc.Sel("audioQueue"))
	return rv
}
func (t TextToSpeechCoreSynthesizer) SetAudioQueue(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setAudioQueue:"), value)
}
func (t TextToSpeechCoreSynthesizer) AudioQueueFlags() uint32 {
	rv := objc.SendIfResponds[uint32](t.ID, objc.Sel("audioQueueFlags"))
	return rv
}
func (t TextToSpeechCoreSynthesizer) SetAudioQueueFlags(value uint32) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setAudioQueueFlags:"), value)
}
func (t TextToSpeechCoreSynthesizer) BundleIdentifier() string {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("bundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (t TextToSpeechCoreSynthesizer) SetBundleIdentifier(value string) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setBundleIdentifier:"), objc.String(value))
}
func (t TextToSpeechCoreSynthesizer) Effects() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("effects"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TextToSpeechCoreSynthesizer) SetEffects(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setEffects:"), value)
}
func (t TextToSpeechCoreSynthesizer) IsPaused() bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("isPaused"))
	return rv
}
func (t TextToSpeechCoreSynthesizer) IsSpeaking() bool {
	rv := objc.SendIfResponds[bool](t.ID, objc.Sel("isSpeaking"))
	return rv
}
func (t TextToSpeechCoreSynthesizer) OfflineChain() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("offlineChain"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TextToSpeechCoreSynthesizer) SetOfflineChain(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setOfflineChain:"), value)
}
func (t TextToSpeechCoreSynthesizer) VoiceResolver() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](t.ID, objc.Sel("voiceResolver"))
	return rv
}
func (t TextToSpeechCoreSynthesizer) SetVoiceResolver(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setVoiceResolver:"), value)
}

// SpeakWithRequestLanguageSynthesizer is a synchronous wrapper around [TextToSpeechCoreSynthesizer.SpeakWithRequestLanguageSynthesizerCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesizer) SpeakWithRequestLanguageSynthesizer(ctx context.Context, request objectivec.IObject, language objectivec.IObject, synthesizer objectivec.IObject) error {
	done := make(chan error, 1)
	t.SpeakWithRequestLanguageSynthesizerCompletionHandler(request, language, synthesizer, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Stop is a synchronous wrapper around [TextToSpeechCoreSynthesizer.StopWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesizer) Stop(ctx context.Context) error {
	done := make(chan error, 1)
	t.StopWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// VoiceWithIdentifier is a synchronous wrapper around [TextToSpeechCoreSynthesizer.VoiceWithIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesizer) VoiceWithIdentifier(ctx context.Context, identifier string) error {
	done := make(chan error, 1)
	t.VoiceWithIdentifierCompletionHandler(identifier, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// VoiceWithLocale is a synchronous wrapper around [TextToSpeechCoreSynthesizer.VoiceWithLocaleCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesizer) VoiceWithLocale(ctx context.Context, locale foundation.NSLocale) error {
	done := make(chan error, 1)
	t.VoiceWithLocaleCompletionHandler(locale, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// WriteWithSpeechPhraseToAudioFileWithAudioSettings is a synchronous wrapper around [TextToSpeechCoreSynthesizer.WriteWithSpeechPhraseToAudioFileWithAudioSettingsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TextToSpeechCoreSynthesizer) WriteWithSpeechPhraseToAudioFileWithAudioSettings(ctx context.Context, phrase string, file foundation.NSURL, settings foundation.INSDictionary) error {
	done := make(chan error, 1)
	t.WriteWithSpeechPhraseToAudioFileWithAudioSettingsCompletionHandler(phrase, file, settings, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
