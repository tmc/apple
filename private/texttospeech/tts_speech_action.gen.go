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

// The class instance for the [TTSSpeechAction] class.
var (
	_TTSSpeechActionClass     TTSSpeechActionClass
	_TTSSpeechActionClassOnce sync.Once
)

func getTTSSpeechActionClass() TTSSpeechActionClass {
	_TTSSpeechActionClassOnce.Do(func() {
		_TTSSpeechActionClass = TTSSpeechActionClass{class: objc.GetClass("TTSSpeechAction")}
	})
	return _TTSSpeechActionClass
}

// GetTTSSpeechActionClass returns the class object for TTSSpeechAction.
func GetTTSSpeechActionClass() TTSSpeechActionClass {
	return getTTSSpeechActionClass()
}

type TTSSpeechActionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSSpeechActionClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSSpeechActionClass) Alloc() TTSSpeechAction {
	rv := objc.Send[TTSSpeechAction](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TTSSpeechAction._detectLanguageFromContent]
//   - [TTSSpeechAction.AttributedString]
//   - [TTSSpeechAction.SetAttributedString]
//   - [TTSSpeechAction.CannotInterrupt]
//   - [TTSSpeechAction.SetCannotInterrupt]
//   - [TTSSpeechAction.CurrentSSMLSnippetIndex]
//   - [TTSSpeechAction.SetCurrentSSMLSnippetIndex]
//   - [TTSSpeechAction.EmojiRangeReplacements]
//   - [TTSSpeechAction.SetEmojiRangeReplacements]
//   - [TTSSpeechAction.FinalSpokenString]
//   - [TTSSpeechAction.SetFinalSpokenString]
//   - [TTSSpeechAction.IgnoreCustomSubstitutions]
//   - [TTSSpeechAction.SetIgnoreCustomSubstitutions]
//   - [TTSSpeechAction.Language]
//   - [TTSSpeechAction.SetLanguage]
//   - [TTSSpeechAction.Pitch]
//   - [TTSSpeechAction.SetPitch]
//   - [TTSSpeechAction.PreprocessAction]
//   - [TTSSpeechAction.ProcessedString]
//   - [TTSSpeechAction.SetProcessedString]
//   - [TTSSpeechAction.SetAudioBufferCallback]
//   - [TTSSpeechAction.SetCompletionCallback]
//   - [TTSSpeechAction.SetMarkerCallback]
//   - [TTSSpeechAction.SetOnMarkerCallback]
//   - [TTSSpeechAction.SetOnPauseCallback]
//   - [TTSSpeechAction.SetOnResumeCallback]
//   - [TTSSpeechAction.SetOnSpeechStartCallback]
//   - [TTSSpeechAction.SetOnWillSpeakRangeCallback]
//   - [TTSSpeechAction.ShouldDetectLanguage]
//   - [TTSSpeechAction.SetShouldDetectLanguage]
//   - [TTSSpeechAction.ShouldPrecomposeString]
//   - [TTSSpeechAction.SetShouldPrecomposeString]
//   - [TTSSpeechAction.ShouldProcessEmoji]
//   - [TTSSpeechAction.SetShouldProcessEmoji]
//   - [TTSSpeechAction.ShouldProcessEmoticons]
//   - [TTSSpeechAction.SetShouldProcessEmoticons]
//   - [TTSSpeechAction.ShouldQueue]
//   - [TTSSpeechAction.SetShouldQueue]
//   - [TTSSpeechAction.SpeakingRate]
//   - [TTSSpeechAction.SetSpeakingRate]
//   - [TTSSpeechAction.SsmlRepresentation]
//   - [TTSSpeechAction.SetSsmlRepresentation]
//   - [TTSSpeechAction.State]
//   - [TTSSpeechAction.SetState]
//   - [TTSSpeechAction.String]
//   - [TTSSpeechAction.SetString]
//   - [TTSSpeechAction.SynthesizeSilently]
//   - [TTSSpeechAction.SetSynthesizeSilently]
//   - [TTSSpeechAction.TaggedSSML]
//   - [TTSSpeechAction.SetTaggedSSML]
//   - [TTSSpeechAction.Utterance]
//   - [TTSSpeechAction.SetUtterance]
//   - [TTSSpeechAction.VoiceIdentifier]
//   - [TTSSpeechAction.SetVoiceIdentifier]
//   - [TTSSpeechAction.VoiceSelection]
//   - [TTSSpeechAction.SetVoiceSelection]
//   - [TTSSpeechAction.Volume]
//   - [TTSSpeechAction.SetVolume]
//   - [TTSSpeechAction.WordCallbackPostProcessedOffset]
//   - [TTSSpeechAction.SetWordCallbackPostProcessedOffset]
type TTSSpeechAction struct {
	objectivec.Object
}

// TTSSpeechActionFromID constructs a [TTSSpeechAction] from an objc.ID.
func TTSSpeechActionFromID(id objc.ID) TTSSpeechAction {
	return TTSSpeechAction{objectivec.Object{ID: id}}
}

// Ensure TTSSpeechAction implements ITTSSpeechAction.
var _ ITTSSpeechAction = TTSSpeechAction{}

// An interface definition for the [TTSSpeechAction] class.
//
// # Methods
//
//   - [ITTSSpeechAction._detectLanguageFromContent]
//   - [ITTSSpeechAction.AttributedString]
//   - [ITTSSpeechAction.SetAttributedString]
//   - [ITTSSpeechAction.CannotInterrupt]
//   - [ITTSSpeechAction.SetCannotInterrupt]
//   - [ITTSSpeechAction.CurrentSSMLSnippetIndex]
//   - [ITTSSpeechAction.SetCurrentSSMLSnippetIndex]
//   - [ITTSSpeechAction.EmojiRangeReplacements]
//   - [ITTSSpeechAction.SetEmojiRangeReplacements]
//   - [ITTSSpeechAction.FinalSpokenString]
//   - [ITTSSpeechAction.SetFinalSpokenString]
//   - [ITTSSpeechAction.IgnoreCustomSubstitutions]
//   - [ITTSSpeechAction.SetIgnoreCustomSubstitutions]
//   - [ITTSSpeechAction.Language]
//   - [ITTSSpeechAction.SetLanguage]
//   - [ITTSSpeechAction.Pitch]
//   - [ITTSSpeechAction.SetPitch]
//   - [ITTSSpeechAction.PreprocessAction]
//   - [ITTSSpeechAction.ProcessedString]
//   - [ITTSSpeechAction.SetProcessedString]
//   - [ITTSSpeechAction.SetAudioBufferCallback]
//   - [ITTSSpeechAction.SetCompletionCallback]
//   - [ITTSSpeechAction.SetMarkerCallback]
//   - [ITTSSpeechAction.SetOnMarkerCallback]
//   - [ITTSSpeechAction.SetOnPauseCallback]
//   - [ITTSSpeechAction.SetOnResumeCallback]
//   - [ITTSSpeechAction.SetOnSpeechStartCallback]
//   - [ITTSSpeechAction.SetOnWillSpeakRangeCallback]
//   - [ITTSSpeechAction.ShouldDetectLanguage]
//   - [ITTSSpeechAction.SetShouldDetectLanguage]
//   - [ITTSSpeechAction.ShouldPrecomposeString]
//   - [ITTSSpeechAction.SetShouldPrecomposeString]
//   - [ITTSSpeechAction.ShouldProcessEmoji]
//   - [ITTSSpeechAction.SetShouldProcessEmoji]
//   - [ITTSSpeechAction.ShouldProcessEmoticons]
//   - [ITTSSpeechAction.SetShouldProcessEmoticons]
//   - [ITTSSpeechAction.ShouldQueue]
//   - [ITTSSpeechAction.SetShouldQueue]
//   - [ITTSSpeechAction.SpeakingRate]
//   - [ITTSSpeechAction.SetSpeakingRate]
//   - [ITTSSpeechAction.SsmlRepresentation]
//   - [ITTSSpeechAction.SetSsmlRepresentation]
//   - [ITTSSpeechAction.State]
//   - [ITTSSpeechAction.SetState]
//   - [ITTSSpeechAction.String]
//   - [ITTSSpeechAction.SetString]
//   - [ITTSSpeechAction.SynthesizeSilently]
//   - [ITTSSpeechAction.SetSynthesizeSilently]
//   - [ITTSSpeechAction.TaggedSSML]
//   - [ITTSSpeechAction.SetTaggedSSML]
//   - [ITTSSpeechAction.Utterance]
//   - [ITTSSpeechAction.SetUtterance]
//   - [ITTSSpeechAction.VoiceIdentifier]
//   - [ITTSSpeechAction.SetVoiceIdentifier]
//   - [ITTSSpeechAction.VoiceSelection]
//   - [ITTSSpeechAction.SetVoiceSelection]
//   - [ITTSSpeechAction.Volume]
//   - [ITTSSpeechAction.SetVolume]
//   - [ITTSSpeechAction.WordCallbackPostProcessedOffset]
//   - [ITTSSpeechAction.SetWordCallbackPostProcessedOffset]
type ITTSSpeechAction interface {
	objectivec.IObject

	// Topic: Methods

	_detectLanguageFromContent() objectivec.IObject
	AttributedString() foundation.NSAttributedString
	SetAttributedString(value foundation.NSAttributedString)
	CannotInterrupt() bool
	SetCannotInterrupt(value bool)
	CurrentSSMLSnippetIndex() uint64
	SetCurrentSSMLSnippetIndex(value uint64)
	EmojiRangeReplacements() foundation.INSArray
	SetEmojiRangeReplacements(value foundation.INSArray)
	FinalSpokenString() string
	SetFinalSpokenString(value string)
	IgnoreCustomSubstitutions() bool
	SetIgnoreCustomSubstitutions(value bool)
	Language() string
	SetLanguage(value string)
	Pitch() float64
	SetPitch(value float64)
	PreprocessAction()
	ProcessedString() string
	SetProcessedString(value string)
	SetAudioBufferCallback(callback VoidHandler)
	SetCompletionCallback(callback VoidHandler)
	SetMarkerCallback(callback VoidHandler)
	SetOnMarkerCallback(callback VoidHandler)
	SetOnPauseCallback(callback VoidHandler)
	SetOnResumeCallback(callback VoidHandler)
	SetOnSpeechStartCallback(callback VoidHandler)
	SetOnWillSpeakRangeCallback(callback VoidHandler)
	ShouldDetectLanguage() bool
	SetShouldDetectLanguage(value bool)
	ShouldPrecomposeString() bool
	SetShouldPrecomposeString(value bool)
	ShouldProcessEmoji() bool
	SetShouldProcessEmoji(value bool)
	ShouldProcessEmoticons() bool
	SetShouldProcessEmoticons(value bool)
	ShouldQueue() bool
	SetShouldQueue(value bool)
	SpeakingRate() float64
	SetSpeakingRate(value float64)
	SsmlRepresentation() string
	SetSsmlRepresentation(value string)
	State() int64
	SetState(value int64)
	String() string
	SetString(value string)
	SynthesizeSilently() bool
	SetSynthesizeSilently(value bool)
	TaggedSSML() unsafe.Pointer
	SetTaggedSSML(value unsafe.Pointer)
	Utterance() avfaudio.AVSpeechUtterance
	SetUtterance(value avfaudio.AVSpeechUtterance)
	VoiceIdentifier() string
	SetVoiceIdentifier(value string)
	VoiceSelection() unsafe.Pointer
	SetVoiceSelection(value unsafe.Pointer)
	Volume() float64
	SetVolume(value float64)
	WordCallbackPostProcessedOffset() int64
	SetWordCallbackPostProcessedOffset(value int64)
}

// Init initializes the instance.
func (t TTSSpeechAction) Init() TTSSpeechAction {
	rv := objc.Send[TTSSpeechAction](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSSpeechAction) Autorelease() TTSSpeechAction {
	rv := objc.Send[TTSSpeechAction](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSSpeechAction creates a new TTSSpeechAction instance.
func NewTTSSpeechAction() TTSSpeechAction {
	class := getTTSSpeechActionClass()
	rv := objc.Send[TTSSpeechAction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (t TTSSpeechAction) _detectLanguageFromContent() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("_detectLanguageFromContent"))
	return objectivec.Object{ID: rv}
}

// DetectLanguageFromContent is an exported wrapper for the private method _detectLanguageFromContent.
func (t TTSSpeechAction) DetectLanguageFromContent() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(t.ID, objc.Sel("_detectLanguageFromContent")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_detectLanguageFromContent"}
		return nil, err
	}
	return t._detectLanguageFromContent(), nil
}

// CanDetectLanguageFromContent reports whether the receiver responds to the private selector _detectLanguageFromContent.
func (t TTSSpeechAction) CanDetectLanguageFromContent() bool {
	return objc.RespondsToSelector(t.ID, objc.Sel("_detectLanguageFromContent"))
}
func (t TTSSpeechAction) PreprocessAction() {
	objc.Send[objc.ID](t.ID, objc.Sel("preprocessAction"))
}
func (t TTSSpeechAction) SetAudioBufferCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.Send[objc.ID](t.ID, objc.Sel("setAudioBufferCallback:"), _block0)
}
func (t TTSSpeechAction) SetCompletionCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.Send[objc.ID](t.ID, objc.Sel("setCompletionCallback:"), _block0)
}
func (t TTSSpeechAction) SetMarkerCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.Send[objc.ID](t.ID, objc.Sel("setMarkerCallback:"), _block0)
}
func (t TTSSpeechAction) SetOnMarkerCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.Send[objc.ID](t.ID, objc.Sel("setOnMarkerCallback:"), _block0)
}
func (t TTSSpeechAction) SetOnPauseCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.Send[objc.ID](t.ID, objc.Sel("setOnPauseCallback:"), _block0)
}
func (t TTSSpeechAction) SetOnResumeCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.Send[objc.ID](t.ID, objc.Sel("setOnResumeCallback:"), _block0)
}
func (t TTSSpeechAction) SetOnSpeechStartCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.Send[objc.ID](t.ID, objc.Sel("setOnSpeechStartCallback:"), _block0)
}
func (t TTSSpeechAction) SetOnWillSpeakRangeCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.Send[objc.ID](t.ID, objc.Sel("setOnWillSpeakRangeCallback:"), _block0)
}

func (_TTSSpeechActionClass TTSSpeechActionClass) ActionWithAttributedStringShouldQueue(string_ objectivec.IObject, queue bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechActionClass.class), objc.Sel("actionWithAttributedString:shouldQueue:"), string_, queue)
	return objectivec.Object{ID: rv}
}
func (_TTSSpeechActionClass TTSSpeechActionClass) ActionWithSSMLRepresentationShouldQueue(sSMLRepresentation objectivec.IObject, queue bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechActionClass.class), objc.Sel("actionWithSSMLRepresentation:shouldQueue:"), sSMLRepresentation, queue)
	return objectivec.Object{ID: rv}
}
func (_TTSSpeechActionClass TTSSpeechActionClass) ActionWithStringShouldQueue(string_ objectivec.IObject, queue bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechActionClass.class), objc.Sel("actionWithString:shouldQueue:"), string_, queue)
	return objectivec.Object{ID: rv}
}
func (_TTSSpeechActionClass TTSSpeechActionClass) Test_setUseMaxSpeechRate(rate bool) {
	objc.Send[objc.ID](objc.ID(_TTSSpeechActionClass.class), objc.Sel("test_setUseMaxSpeechRate:"), rate)
}

func (t TTSSpeechAction) AttributedString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("attributedString"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (t TTSSpeechAction) SetAttributedString(value foundation.NSAttributedString) {
	objc.Send[struct{}](t.ID, objc.Sel("setAttributedString:"), value)
}
func (t TTSSpeechAction) CannotInterrupt() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("cannotInterrupt"))
	return rv
}
func (t TTSSpeechAction) SetCannotInterrupt(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setCannotInterrupt:"), value)
}
func (t TTSSpeechAction) CurrentSSMLSnippetIndex() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("currentSSMLSnippetIndex"))
	return rv
}
func (t TTSSpeechAction) SetCurrentSSMLSnippetIndex(value uint64) {
	objc.Send[struct{}](t.ID, objc.Sel("setCurrentSSMLSnippetIndex:"), value)
}
func (t TTSSpeechAction) EmojiRangeReplacements() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("emojiRangeReplacements"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSSpeechAction) SetEmojiRangeReplacements(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setEmojiRangeReplacements:"), value)
}
func (t TTSSpeechAction) FinalSpokenString() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("finalSpokenString"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechAction) SetFinalSpokenString(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setFinalSpokenString:"), objc.String(value))
}
func (t TTSSpeechAction) IgnoreCustomSubstitutions() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("ignoreCustomSubstitutions"))
	return rv
}
func (t TTSSpeechAction) SetIgnoreCustomSubstitutions(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setIgnoreCustomSubstitutions:"), value)
}
func (t TTSSpeechAction) Language() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("language"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechAction) SetLanguage(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setLanguage:"), objc.String(value))
}
func (t TTSSpeechAction) Pitch() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("pitch"))
	return rv
}
func (t TTSSpeechAction) SetPitch(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setPitch:"), value)
}
func (t TTSSpeechAction) ProcessedString() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("processedString"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechAction) SetProcessedString(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setProcessedString:"), objc.String(value))
}
func (t TTSSpeechAction) ShouldDetectLanguage() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("shouldDetectLanguage"))
	return rv
}
func (t TTSSpeechAction) SetShouldDetectLanguage(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setShouldDetectLanguage:"), value)
}
func (t TTSSpeechAction) ShouldPrecomposeString() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("shouldPrecomposeString"))
	return rv
}
func (t TTSSpeechAction) SetShouldPrecomposeString(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setShouldPrecomposeString:"), value)
}
func (t TTSSpeechAction) ShouldProcessEmoji() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("shouldProcessEmoji"))
	return rv
}
func (t TTSSpeechAction) SetShouldProcessEmoji(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setShouldProcessEmoji:"), value)
}
func (t TTSSpeechAction) ShouldProcessEmoticons() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("shouldProcessEmoticons"))
	return rv
}
func (t TTSSpeechAction) SetShouldProcessEmoticons(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setShouldProcessEmoticons:"), value)
}
func (t TTSSpeechAction) ShouldQueue() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("shouldQueue"))
	return rv
}
func (t TTSSpeechAction) SetShouldQueue(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setShouldQueue:"), value)
}
func (t TTSSpeechAction) SpeakingRate() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("speakingRate"))
	return rv
}
func (t TTSSpeechAction) SetSpeakingRate(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setSpeakingRate:"), value)
}
func (t TTSSpeechAction) SsmlRepresentation() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("ssmlRepresentation"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechAction) SetSsmlRepresentation(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setSsmlRepresentation:"), objc.String(value))
}
func (t TTSSpeechAction) State() int64 {
	rv := objc.Send[int64](t.ID, objc.Sel("state"))
	return rv
}
func (t TTSSpeechAction) SetState(value int64) {
	objc.Send[struct{}](t.ID, objc.Sel("setState:"), value)
}
func (t TTSSpeechAction) String() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("string"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechAction) SetString(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setString:"), objc.String(value))
}
func (t TTSSpeechAction) SynthesizeSilently() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("synthesizeSilently"))
	return rv
}
func (t TTSSpeechAction) SetSynthesizeSilently(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setSynthesizeSilently:"), value)
}
func (t TTSSpeechAction) TaggedSSML() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t.ID, objc.Sel("taggedSSML"))
	return rv
}
func (t TTSSpeechAction) SetTaggedSSML(value unsafe.Pointer) {
	objc.Send[struct{}](t.ID, objc.Sel("setTaggedSSML:"), value)
}
func (t TTSSpeechAction) Utterance() avfaudio.AVSpeechUtterance {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("utterance"))
	return avfaudio.AVSpeechUtteranceFromID(objc.ID(rv))
}
func (t TTSSpeechAction) SetUtterance(value avfaudio.AVSpeechUtterance) {
	objc.Send[struct{}](t.ID, objc.Sel("setUtterance:"), value)
}
func (t TTSSpeechAction) VoiceIdentifier() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("voiceIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechAction) SetVoiceIdentifier(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setVoiceIdentifier:"), objc.String(value))
}
func (t TTSSpeechAction) VoiceSelection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t.ID, objc.Sel("voiceSelection"))
	return rv
}
func (t TTSSpeechAction) SetVoiceSelection(value unsafe.Pointer) {
	objc.Send[struct{}](t.ID, objc.Sel("setVoiceSelection:"), value)
}
func (t TTSSpeechAction) Volume() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("volume"))
	return rv
}
func (t TTSSpeechAction) SetVolume(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setVolume:"), value)
}
func (t TTSSpeechAction) WordCallbackPostProcessedOffset() int64 {
	rv := objc.Send[int64](t.ID, objc.Sel("wordCallbackPostProcessedOffset"))
	return rv
}
func (t TTSSpeechAction) SetWordCallbackPostProcessedOffset(value int64) {
	objc.Send[struct{}](t.ID, objc.Sel("setWordCallbackPostProcessedOffset:"), value)
}

// SetAudioBufferCallbackSync is a synchronous wrapper around [TTSSpeechAction.SetAudioBufferCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSSpeechAction) SetAudioBufferCallbackSync(ctx context.Context) error {
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

// SetCompletionCallbackSync is a synchronous wrapper around [TTSSpeechAction.SetCompletionCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSSpeechAction) SetCompletionCallbackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.SetCompletionCallback(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetMarkerCallbackSync is a synchronous wrapper around [TTSSpeechAction.SetMarkerCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSSpeechAction) SetMarkerCallbackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.SetMarkerCallback(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetOnMarkerCallbackSync is a synchronous wrapper around [TTSSpeechAction.SetOnMarkerCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSSpeechAction) SetOnMarkerCallbackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.SetOnMarkerCallback(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetOnPauseCallbackSync is a synchronous wrapper around [TTSSpeechAction.SetOnPauseCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSSpeechAction) SetOnPauseCallbackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.SetOnPauseCallback(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetOnResumeCallbackSync is a synchronous wrapper around [TTSSpeechAction.SetOnResumeCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSSpeechAction) SetOnResumeCallbackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.SetOnResumeCallback(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetOnSpeechStartCallbackSync is a synchronous wrapper around [TTSSpeechAction.SetOnSpeechStartCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSSpeechAction) SetOnSpeechStartCallbackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.SetOnSpeechStartCallback(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetOnWillSpeakRangeCallbackSync is a synchronous wrapper around [TTSSpeechAction.SetOnWillSpeakRangeCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSSpeechAction) SetOnWillSpeakRangeCallbackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.SetOnWillSpeakRangeCallback(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
