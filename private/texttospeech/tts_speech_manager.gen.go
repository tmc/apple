// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSSpeechManager] class.
var (
	_TTSSpeechManagerClass     TTSSpeechManagerClass
	_TTSSpeechManagerClassOnce sync.Once
)

func getTTSSpeechManagerClass() TTSSpeechManagerClass {
	_TTSSpeechManagerClassOnce.Do(func() {
		_TTSSpeechManagerClass = TTSSpeechManagerClass{class: objc.GetClass("TTSSpeechManager")}
	})
	return _TTSSpeechManagerClass
}

// GetTTSSpeechManagerClass returns the class object for TTSSpeechManager.
func GetTTSSpeechManagerClass() TTSSpeechManagerClass {
	return getTTSSpeechManagerClass()
}

type TTSSpeechManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSSpeechManagerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSSpeechManagerClass) Alloc() TTSSpeechManager {
	rv := objc.Send[TTSSpeechManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSSpeechManager.__speechJobFinished]
//   - [TTSSpeechManager._clearSpeechQueue]
//   - [TTSSpeechManager._continueSpeaking]
//   - [TTSSpeechManager._didBeginInterruption]
//   - [TTSSpeechManager._didEndInterruption]
//   - [TTSSpeechManager._dispatchSpeechAction]
//   - [TTSSpeechManager._enqueueSelectorOnSpeechThreadObjectWaitUntilDone]
//   - [TTSSpeechManager._handleAudioInterruption]
//   - [TTSSpeechManager._handleMediaServicesWereLost]
//   - [TTSSpeechManager._handleMediaServicesWereReset]
//   - [TTSSpeechManager._initialize]
//   - [TTSSpeechManager._isSpeaking]
//   - [TTSSpeechManager._pauseSpeaking]
//   - [TTSSpeechManager._phonemeSubstitutionsForAction]
//   - [TTSSpeechManager._processAudioBufferCallback]
//   - [TTSSpeechManager._processDidContinueCallback]
//   - [TTSSpeechManager._processDidEncounterMarker]
//   - [TTSSpeechManager._processDidPauseCallback]
//   - [TTSSpeechManager._processDidStartCallback]
//   - [TTSSpeechManager._processWillSpeechRange]
//   - [TTSSpeechManager._resetInterruptionTracking]
//   - [TTSSpeechManager._setVoiceForActionSnippet]
//   - [TTSSpeechManager._speechJobFinishedAction]
//   - [TTSSpeechManager._startNextSpeechJob]
//   - [TTSSpeechManager._stopSpeaking]
//   - [TTSSpeechManager._tearDown]
//   - [TTSSpeechManager._updateAudioSessionProperties]
//   - [TTSSpeechManager._updateAuxiliarySession]
//   - [TTSSpeechManager._updateUserSubstitutions]
//   - [TTSSpeechManager.AudioDeactivatorTimer]
//   - [TTSSpeechManager.SetAudioDeactivatorTimer]
//   - [TTSSpeechManager.AudioDeviceId]
//   - [TTSSpeechManager.SetAudioDeviceId]
//   - [TTSSpeechManager.AudioInterruptionStartedTime]
//   - [TTSSpeechManager.SetAudioInterruptionStartedTime]
//   - [TTSSpeechManager.AudioOperationQueue]
//   - [TTSSpeechManager.SetAudioOperationQueue]
//   - [TTSSpeechManager.AudioQueueFlags]
//   - [TTSSpeechManager.SetAudioQueueFlags]
//   - [TTSSpeechManager.AudioSession]
//   - [TTSSpeechManager.SetAudioSession]
//   - [TTSSpeechManager.AudioSessionCategory]
//   - [TTSSpeechManager.SetAudioSessionCategory]
//   - [TTSSpeechManager.AudioSessionCategoryOptions]
//   - [TTSSpeechManager.SetAudioSessionCategoryOptions]
//   - [TTSSpeechManager.AudioSessionInactiveTimeout]
//   - [TTSSpeechManager.SetAudioSessionInactiveTimeout]
//   - [TTSSpeechManager.ClearSpeechQueue]
//   - [TTSSpeechManager.ContinueSpeaking]
//   - [TTSSpeechManager.DidRequestPauseSpeakingDuringAudioInterruption]
//   - [TTSSpeechManager.SetDidRequestPauseSpeakingDuringAudioInterruption]
//   - [TTSSpeechManager.DidRequestResumeSpeakingDuringAudioInterruption]
//   - [TTSSpeechManager.SetDidRequestResumeSpeakingDuringAudioInterruption]
//   - [TTSSpeechManager.DidRequestStartSpeakingDuringAudioInterruption]
//   - [TTSSpeechManager.SetDidRequestStartSpeakingDuringAudioInterruption]
//   - [TTSSpeechManager.DispatchSpeechAction]
//   - [TTSSpeechManager.ExternalVoiceIdentifierUsedForLanguage]
//   - [TTSSpeechManager.HandleAudioInterruption]
//   - [TTSSpeechManager.HandleAudioSessionObservers]
//   - [TTSSpeechManager.HandleMediaServicesWereLost]
//   - [TTSSpeechManager.HandleMediaServicesWereReset]
//   - [TTSSpeechManager.IsInAudioInterruption]
//   - [TTSSpeechManager.SetIsInAudioInterruption]
//   - [TTSSpeechManager.IsPaused]
//   - [TTSSpeechManager.SetIsPaused]
//   - [TTSSpeechManager.IsSpeaking]
//   - [TTSSpeechManager.OriginalSpeechRateForJobOverride]
//   - [TTSSpeechManager.SetOriginalSpeechRateForJobOverride]
//   - [TTSSpeechManager.OutputChannels]
//   - [TTSSpeechManager.SetOutputChannels]
//   - [TTSSpeechManager.PauseSpeaking]
//   - [TTSSpeechManager.RequestedActionDuringAudioInterruption]
//   - [TTSSpeechManager.SetRequestedActionDuringAudioInterruption]
//   - [TTSSpeechManager.SetActiveOptions]
//   - [TTSSpeechManager.SetSetActiveOptions]
//   - [TTSSpeechManager.SetRequestWillStart]
//   - [TTSSpeechManager.ShouldHandleAudioInterruptions]
//   - [TTSSpeechManager.SetShouldHandleAudioInterruptions]
//   - [TTSSpeechManager.ShowControlCenterControls]
//   - [TTSSpeechManager.SpeechEnabled]
//   - [TTSSpeechManager.SetSpeechEnabled]
//   - [TTSSpeechManager.SpeechSource]
//   - [TTSSpeechManager.SetSpeechSource]
//   - [TTSSpeechManager.SpeechSynthesizerDidContinueSpeakingRequest]
//   - [TTSSpeechManager.SpeechSynthesizerDidEncounterMarkerForRequest]
//   - [TTSSpeechManager.SpeechSynthesizerDidFinishSpeakingRequestSuccessfullyWithError]
//   - [TTSSpeechManager.SpeechSynthesizerDidPauseSpeakingRequest]
//   - [TTSSpeechManager.SpeechSynthesizerDidStartSpeakingRequest]
//   - [TTSSpeechManager.StopSpeaking]
//   - [TTSSpeechManager.StopSpeakingWithSpeaking]
//   - [TTSSpeechManager.TearDown]
//   - [TTSSpeechManager.UsesAuxiliarySession]
//   - [TTSSpeechManager.SetUsesAuxiliarySession]
//   - [TTSSpeechManager.VoiceIdentifierUsedForLanguage]
//   - [TTSSpeechManager.WasSpeakingBeforeAudioInterruption]
//   - [TTSSpeechManager.SetWasSpeakingBeforeAudioInterruption]
//   - [TTSSpeechManager.DebugDescription]
//   - [TTSSpeechManager.Description]
//   - [TTSSpeechManager.Hash]
//   - [TTSSpeechManager.Superclass]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager
type TTSSpeechManager struct {
	objectivec.Object
}

// TTSSpeechManagerFromID constructs a [TTSSpeechManager] from an objc.ID.
func TTSSpeechManagerFromID(id objc.ID) TTSSpeechManager {
	return TTSSpeechManager{objectivec.Object{ID: id}}
}
// Ensure TTSSpeechManager implements ITTSSpeechManager.
var _ ITTSSpeechManager = TTSSpeechManager{}

// An interface definition for the [TTSSpeechManager] class.
//
// # Methods
//
//   - [ITTSSpeechManager.__speechJobFinished]
//   - [ITTSSpeechManager._clearSpeechQueue]
//   - [ITTSSpeechManager._continueSpeaking]
//   - [ITTSSpeechManager._didBeginInterruption]
//   - [ITTSSpeechManager._didEndInterruption]
//   - [ITTSSpeechManager._dispatchSpeechAction]
//   - [ITTSSpeechManager._enqueueSelectorOnSpeechThreadObjectWaitUntilDone]
//   - [ITTSSpeechManager._handleAudioInterruption]
//   - [ITTSSpeechManager._handleMediaServicesWereLost]
//   - [ITTSSpeechManager._handleMediaServicesWereReset]
//   - [ITTSSpeechManager._initialize]
//   - [ITTSSpeechManager._isSpeaking]
//   - [ITTSSpeechManager._pauseSpeaking]
//   - [ITTSSpeechManager._phonemeSubstitutionsForAction]
//   - [ITTSSpeechManager._processAudioBufferCallback]
//   - [ITTSSpeechManager._processDidContinueCallback]
//   - [ITTSSpeechManager._processDidEncounterMarker]
//   - [ITTSSpeechManager._processDidPauseCallback]
//   - [ITTSSpeechManager._processDidStartCallback]
//   - [ITTSSpeechManager._processWillSpeechRange]
//   - [ITTSSpeechManager._resetInterruptionTracking]
//   - [ITTSSpeechManager._setVoiceForActionSnippet]
//   - [ITTSSpeechManager._speechJobFinishedAction]
//   - [ITTSSpeechManager._startNextSpeechJob]
//   - [ITTSSpeechManager._stopSpeaking]
//   - [ITTSSpeechManager._tearDown]
//   - [ITTSSpeechManager._updateAudioSessionProperties]
//   - [ITTSSpeechManager._updateAuxiliarySession]
//   - [ITTSSpeechManager._updateUserSubstitutions]
//   - [ITTSSpeechManager.AudioDeactivatorTimer]
//   - [ITTSSpeechManager.SetAudioDeactivatorTimer]
//   - [ITTSSpeechManager.AudioDeviceId]
//   - [ITTSSpeechManager.SetAudioDeviceId]
//   - [ITTSSpeechManager.AudioInterruptionStartedTime]
//   - [ITTSSpeechManager.SetAudioInterruptionStartedTime]
//   - [ITTSSpeechManager.AudioOperationQueue]
//   - [ITTSSpeechManager.SetAudioOperationQueue]
//   - [ITTSSpeechManager.AudioQueueFlags]
//   - [ITTSSpeechManager.SetAudioQueueFlags]
//   - [ITTSSpeechManager.AudioSession]
//   - [ITTSSpeechManager.SetAudioSession]
//   - [ITTSSpeechManager.AudioSessionCategory]
//   - [ITTSSpeechManager.SetAudioSessionCategory]
//   - [ITTSSpeechManager.AudioSessionCategoryOptions]
//   - [ITTSSpeechManager.SetAudioSessionCategoryOptions]
//   - [ITTSSpeechManager.AudioSessionInactiveTimeout]
//   - [ITTSSpeechManager.SetAudioSessionInactiveTimeout]
//   - [ITTSSpeechManager.ClearSpeechQueue]
//   - [ITTSSpeechManager.ContinueSpeaking]
//   - [ITTSSpeechManager.DidRequestPauseSpeakingDuringAudioInterruption]
//   - [ITTSSpeechManager.SetDidRequestPauseSpeakingDuringAudioInterruption]
//   - [ITTSSpeechManager.DidRequestResumeSpeakingDuringAudioInterruption]
//   - [ITTSSpeechManager.SetDidRequestResumeSpeakingDuringAudioInterruption]
//   - [ITTSSpeechManager.DidRequestStartSpeakingDuringAudioInterruption]
//   - [ITTSSpeechManager.SetDidRequestStartSpeakingDuringAudioInterruption]
//   - [ITTSSpeechManager.DispatchSpeechAction]
//   - [ITTSSpeechManager.ExternalVoiceIdentifierUsedForLanguage]
//   - [ITTSSpeechManager.HandleAudioInterruption]
//   - [ITTSSpeechManager.HandleAudioSessionObservers]
//   - [ITTSSpeechManager.HandleMediaServicesWereLost]
//   - [ITTSSpeechManager.HandleMediaServicesWereReset]
//   - [ITTSSpeechManager.IsInAudioInterruption]
//   - [ITTSSpeechManager.SetIsInAudioInterruption]
//   - [ITTSSpeechManager.IsPaused]
//   - [ITTSSpeechManager.SetIsPaused]
//   - [ITTSSpeechManager.IsSpeaking]
//   - [ITTSSpeechManager.OriginalSpeechRateForJobOverride]
//   - [ITTSSpeechManager.SetOriginalSpeechRateForJobOverride]
//   - [ITTSSpeechManager.OutputChannels]
//   - [ITTSSpeechManager.SetOutputChannels]
//   - [ITTSSpeechManager.PauseSpeaking]
//   - [ITTSSpeechManager.RequestedActionDuringAudioInterruption]
//   - [ITTSSpeechManager.SetRequestedActionDuringAudioInterruption]
//   - [ITTSSpeechManager.SetActiveOptions]
//   - [ITTSSpeechManager.SetSetActiveOptions]
//   - [ITTSSpeechManager.SetRequestWillStart]
//   - [ITTSSpeechManager.ShouldHandleAudioInterruptions]
//   - [ITTSSpeechManager.SetShouldHandleAudioInterruptions]
//   - [ITTSSpeechManager.ShowControlCenterControls]
//   - [ITTSSpeechManager.SpeechEnabled]
//   - [ITTSSpeechManager.SetSpeechEnabled]
//   - [ITTSSpeechManager.SpeechSource]
//   - [ITTSSpeechManager.SetSpeechSource]
//   - [ITTSSpeechManager.SpeechSynthesizerDidContinueSpeakingRequest]
//   - [ITTSSpeechManager.SpeechSynthesizerDidEncounterMarkerForRequest]
//   - [ITTSSpeechManager.SpeechSynthesizerDidFinishSpeakingRequestSuccessfullyWithError]
//   - [ITTSSpeechManager.SpeechSynthesizerDidPauseSpeakingRequest]
//   - [ITTSSpeechManager.SpeechSynthesizerDidStartSpeakingRequest]
//   - [ITTSSpeechManager.StopSpeaking]
//   - [ITTSSpeechManager.StopSpeakingWithSpeaking]
//   - [ITTSSpeechManager.TearDown]
//   - [ITTSSpeechManager.UsesAuxiliarySession]
//   - [ITTSSpeechManager.SetUsesAuxiliarySession]
//   - [ITTSSpeechManager.VoiceIdentifierUsedForLanguage]
//   - [ITTSSpeechManager.WasSpeakingBeforeAudioInterruption]
//   - [ITTSSpeechManager.SetWasSpeakingBeforeAudioInterruption]
//   - [ITTSSpeechManager.DebugDescription]
//   - [ITTSSpeechManager.Description]
//   - [ITTSSpeechManager.Hash]
//   - [ITTSSpeechManager.Superclass]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager
type ITTSSpeechManager interface {
	objectivec.IObject

	// Topic: Methods

	__speechJobFinished(finished objectivec.IObject)
	_clearSpeechQueue()
	_continueSpeaking()
	_didBeginInterruption()
	_didEndInterruption()
	_dispatchSpeechAction(action objectivec.IObject)
	_enqueueSelectorOnSpeechThreadObjectWaitUntilDone(thread objc.SEL, object objectivec.IObject, done bool) bool
	_handleAudioInterruption(interruption objectivec.IObject)
	_handleMediaServicesWereLost(lost objectivec.IObject)
	_handleMediaServicesWereReset(reset objectivec.IObject)
	_initialize()
	_isSpeaking(speaking objectivec.IObject)
	_pauseSpeaking(speaking objectivec.IObject)
	_phonemeSubstitutionsForAction(action objectivec.IObject) objectivec.IObject
	_processAudioBufferCallback(callback objectivec.IObject)
	_processDidContinueCallback(callback objectivec.IObject)
	_processDidEncounterMarker(marker objectivec.IObject)
	_processDidPauseCallback(callback objectivec.IObject)
	_processDidStartCallback(callback objectivec.IObject)
	_processWillSpeechRange(range_ objectivec.IObject)
	_resetInterruptionTracking()
	_setVoiceForActionSnippet(action objectivec.IObject, snippet objectivec.IObject)
	_speechJobFinishedAction(finished bool, action objectivec.IObject)
	_startNextSpeechJob()
	_stopSpeaking(speaking objectivec.IObject)
	_tearDown()
	_updateAudioSessionProperties()
	_updateAuxiliarySession()
	_updateUserSubstitutions()
	AudioDeactivatorTimer() unsafe.Pointer
	SetAudioDeactivatorTimer(value unsafe.Pointer)
	AudioDeviceId() uint32
	SetAudioDeviceId(value uint32)
	AudioInterruptionStartedTime() float64
	SetAudioInterruptionStartedTime(value float64)
	AudioOperationQueue() objectivec.Object
	SetAudioOperationQueue(value objectivec.Object)
	AudioQueueFlags() uint32
	SetAudioQueueFlags(value uint32)
	AudioSession() objc.ID
	SetAudioSession(value objc.ID)
	AudioSessionCategory() string
	SetAudioSessionCategory(value string)
	AudioSessionCategoryOptions() uint64
	SetAudioSessionCategoryOptions(value uint64)
	AudioSessionInactiveTimeout() float64
	SetAudioSessionInactiveTimeout(value float64)
	ClearSpeechQueue()
	ContinueSpeaking()
	DidRequestPauseSpeakingDuringAudioInterruption() bool
	SetDidRequestPauseSpeakingDuringAudioInterruption(value bool)
	DidRequestResumeSpeakingDuringAudioInterruption() bool
	SetDidRequestResumeSpeakingDuringAudioInterruption(value bool)
	DidRequestStartSpeakingDuringAudioInterruption() bool
	SetDidRequestStartSpeakingDuringAudioInterruption(value bool)
	DispatchSpeechAction(action objectivec.IObject)
	ExternalVoiceIdentifierUsedForLanguage(language objectivec.IObject) objectivec.IObject
	HandleAudioInterruption(interruption objectivec.IObject)
	HandleAudioSessionObservers(observers bool)
	HandleMediaServicesWereLost(lost objectivec.IObject)
	HandleMediaServicesWereReset(reset objectivec.IObject)
	IsInAudioInterruption() bool
	SetIsInAudioInterruption(value bool)
	IsPaused() bool
	SetIsPaused(value bool)
	IsSpeaking() bool
	OriginalSpeechRateForJobOverride() foundation.NSNumber
	SetOriginalSpeechRateForJobOverride(value foundation.NSNumber)
	OutputChannels() foundation.INSArray
	SetOutputChannels(value foundation.INSArray)
	PauseSpeaking(speaking int64)
	RequestedActionDuringAudioInterruption() ITTSSpeechAction
	SetRequestedActionDuringAudioInterruption(value ITTSSpeechAction)
	SetActiveOptions() uint64
	SetSetActiveOptions(value uint64)
	SetRequestWillStart(start VoidHandler)
	ShouldHandleAudioInterruptions() bool
	SetShouldHandleAudioInterruptions(value bool)
	ShowControlCenterControls() bool
	SpeechEnabled() bool
	SetSpeechEnabled(value bool)
	SpeechSource() string
	SetSpeechSource(value string)
	SpeechSynthesizerDidContinueSpeakingRequest(synthesizer objectivec.IObject, request objectivec.IObject)
	SpeechSynthesizerDidEncounterMarkerForRequest(synthesizer objectivec.IObject, marker objectivec.IObject, request objectivec.IObject)
	SpeechSynthesizerDidFinishSpeakingRequestSuccessfullyWithError(synthesizer objectivec.IObject, request objectivec.IObject, successfully bool, error_ objectivec.IObject)
	SpeechSynthesizerDidPauseSpeakingRequest(synthesizer objectivec.IObject, request objectivec.IObject)
	SpeechSynthesizerDidStartSpeakingRequest(synthesizer objectivec.IObject, request objectivec.IObject)
	StopSpeaking()
	StopSpeakingWithSpeaking(speaking int64)
	TearDown()
	UsesAuxiliarySession() bool
	SetUsesAuxiliarySession(value bool)
	VoiceIdentifierUsedForLanguage(language objectivec.IObject) objectivec.IObject
	WasSpeakingBeforeAudioInterruption() bool
	SetWasSpeakingBeforeAudioInterruption(value bool)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (t TTSSpeechManager) Init() TTSSpeechManager {
	rv := objc.Send[TTSSpeechManager](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSSpeechManager) Autorelease() TTSSpeechManager {
	rv := objc.Send[TTSSpeechManager](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSSpeechManager creates a new TTSSpeechManager instance.
func NewTTSSpeechManager() TTSSpeechManager {
	class := getTTSSpeechManagerClass()
	rv := objc.Send[TTSSpeechManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/__speechJobFinished:
func (t TTSSpeechManager) __speechJobFinished(finished objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("__speechJobFinished:"), finished)
}

// SpeechJobFinished is an exported wrapper for the private method __speechJobFinished.
func (t TTSSpeechManager) SpeechJobFinished(finished objectivec.IObject) {
	t.__speechJobFinished(finished)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_clearSpeechQueue
func (t TTSSpeechManager) _clearSpeechQueue() {
	objc.Send[objc.ID](t.ID, objc.Sel("_clearSpeechQueue"))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_continueSpeaking
func (t TTSSpeechManager) _continueSpeaking() {
	objc.Send[objc.ID](t.ID, objc.Sel("_continueSpeaking"))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_didBeginInterruption
func (t TTSSpeechManager) _didBeginInterruption() {
	objc.Send[objc.ID](t.ID, objc.Sel("_didBeginInterruption"))
}

// DidBeginInterruption is an exported wrapper for the private method _didBeginInterruption.
func (t TTSSpeechManager) DidBeginInterruption() {
	t._didBeginInterruption()
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_didEndInterruption
func (t TTSSpeechManager) _didEndInterruption() {
	objc.Send[objc.ID](t.ID, objc.Sel("_didEndInterruption"))
}

// DidEndInterruption is an exported wrapper for the private method _didEndInterruption.
func (t TTSSpeechManager) DidEndInterruption() {
	t._didEndInterruption()
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_dispatchSpeechAction:
func (t TTSSpeechManager) _dispatchSpeechAction(action objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("_dispatchSpeechAction:"), action)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_enqueueSelectorOnSpeechThread:object:waitUntilDone:
func (t TTSSpeechManager) _enqueueSelectorOnSpeechThreadObjectWaitUntilDone(thread objc.SEL, object objectivec.IObject, done bool) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("_enqueueSelectorOnSpeechThread:object:waitUntilDone:"), thread, object, done)
	return rv
}

// EnqueueSelectorOnSpeechThreadObjectWaitUntilDone is an exported wrapper for the private method _enqueueSelectorOnSpeechThreadObjectWaitUntilDone.
func (t TTSSpeechManager) EnqueueSelectorOnSpeechThreadObjectWaitUntilDone(thread objc.SEL, object objectivec.IObject, done bool) bool {
	return t._enqueueSelectorOnSpeechThreadObjectWaitUntilDone(thread, object, done)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_handleAudioInterruption:
func (t TTSSpeechManager) _handleAudioInterruption(interruption objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("_handleAudioInterruption:"), interruption)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_handleMediaServicesWereLost:
func (t TTSSpeechManager) _handleMediaServicesWereLost(lost objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("_handleMediaServicesWereLost:"), lost)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_handleMediaServicesWereReset:
func (t TTSSpeechManager) _handleMediaServicesWereReset(reset objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("_handleMediaServicesWereReset:"), reset)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_initialize
func (t TTSSpeechManager) _initialize() {
	objc.Send[objc.ID](t.ID, objc.Sel("_initialize"))
}

// Initialize is an exported wrapper for the private method _initialize.
func (t TTSSpeechManager) Initialize() {
	t._initialize()
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_isSpeaking:
func (t TTSSpeechManager) _isSpeaking(speaking objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("_isSpeaking:"), speaking)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_pauseSpeaking:
func (t TTSSpeechManager) _pauseSpeaking(speaking objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("_pauseSpeaking:"), speaking)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_phonemeSubstitutionsForAction:
func (t TTSSpeechManager) _phonemeSubstitutionsForAction(action objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("_phonemeSubstitutionsForAction:"), action)
	return objectivec.Object{ID: rv}
}

// PhonemeSubstitutionsForAction is an exported wrapper for the private method _phonemeSubstitutionsForAction.
func (t TTSSpeechManager) PhonemeSubstitutionsForAction(action objectivec.IObject) objectivec.IObject {
	return t._phonemeSubstitutionsForAction(action)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_processAudioBufferCallback:
func (t TTSSpeechManager) _processAudioBufferCallback(callback objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("_processAudioBufferCallback:"), callback)
}

// ProcessAudioBufferCallback is an exported wrapper for the private method _processAudioBufferCallback.
func (t TTSSpeechManager) ProcessAudioBufferCallback(callback objectivec.IObject) {
	t._processAudioBufferCallback(callback)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_processDidContinueCallback:
func (t TTSSpeechManager) _processDidContinueCallback(callback objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("_processDidContinueCallback:"), callback)
}

// ProcessDidContinueCallback is an exported wrapper for the private method _processDidContinueCallback.
func (t TTSSpeechManager) ProcessDidContinueCallback(callback objectivec.IObject) {
	t._processDidContinueCallback(callback)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_processDidEncounterMarker:
func (t TTSSpeechManager) _processDidEncounterMarker(marker objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("_processDidEncounterMarker:"), marker)
}

// ProcessDidEncounterMarker is an exported wrapper for the private method _processDidEncounterMarker.
func (t TTSSpeechManager) ProcessDidEncounterMarker(marker objectivec.IObject) {
	t._processDidEncounterMarker(marker)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_processDidPauseCallback:
func (t TTSSpeechManager) _processDidPauseCallback(callback objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("_processDidPauseCallback:"), callback)
}

// ProcessDidPauseCallback is an exported wrapper for the private method _processDidPauseCallback.
func (t TTSSpeechManager) ProcessDidPauseCallback(callback objectivec.IObject) {
	t._processDidPauseCallback(callback)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_processDidStartCallback:
func (t TTSSpeechManager) _processDidStartCallback(callback objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("_processDidStartCallback:"), callback)
}

// ProcessDidStartCallback is an exported wrapper for the private method _processDidStartCallback.
func (t TTSSpeechManager) ProcessDidStartCallback(callback objectivec.IObject) {
	t._processDidStartCallback(callback)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_processWillSpeechRange:
func (t TTSSpeechManager) _processWillSpeechRange(range_ objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("_processWillSpeechRange:"), range_)
}

// ProcessWillSpeechRange is an exported wrapper for the private method _processWillSpeechRange.
func (t TTSSpeechManager) ProcessWillSpeechRange(range_ objectivec.IObject) {
	t._processWillSpeechRange(range_)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_resetInterruptionTracking
func (t TTSSpeechManager) _resetInterruptionTracking() {
	objc.Send[objc.ID](t.ID, objc.Sel("_resetInterruptionTracking"))
}

// ResetInterruptionTracking is an exported wrapper for the private method _resetInterruptionTracking.
func (t TTSSpeechManager) ResetInterruptionTracking() {
	t._resetInterruptionTracking()
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_setVoiceForAction:snippet:
func (t TTSSpeechManager) _setVoiceForActionSnippet(action objectivec.IObject, snippet objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("_setVoiceForAction:snippet:"), action, snippet)
}

// SetVoiceForActionSnippet is an exported wrapper for the private method _setVoiceForActionSnippet.
func (t TTSSpeechManager) SetVoiceForActionSnippet(action objectivec.IObject, snippet objectivec.IObject) {
	t._setVoiceForActionSnippet(action, snippet)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_speechJobFinished:action:
func (t TTSSpeechManager) _speechJobFinishedAction(finished bool, action objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("_speechJobFinished:action:"), finished, action)
}

// SpeechJobFinishedAction is an exported wrapper for the private method _speechJobFinishedAction.
func (t TTSSpeechManager) SpeechJobFinishedAction(finished bool, action objectivec.IObject) {
	t._speechJobFinishedAction(finished, action)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_startNextSpeechJob
func (t TTSSpeechManager) _startNextSpeechJob() {
	objc.Send[objc.ID](t.ID, objc.Sel("_startNextSpeechJob"))
}

// StartNextSpeechJob is an exported wrapper for the private method _startNextSpeechJob.
func (t TTSSpeechManager) StartNextSpeechJob() {
	t._startNextSpeechJob()
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_stopSpeaking:
func (t TTSSpeechManager) _stopSpeaking(speaking objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("_stopSpeaking:"), speaking)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_tearDown
func (t TTSSpeechManager) _tearDown() {
	objc.Send[objc.ID](t.ID, objc.Sel("_tearDown"))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_updateAudioSessionProperties
func (t TTSSpeechManager) _updateAudioSessionProperties() {
	objc.Send[objc.ID](t.ID, objc.Sel("_updateAudioSessionProperties"))
}

// UpdateAudioSessionProperties is an exported wrapper for the private method _updateAudioSessionProperties.
func (t TTSSpeechManager) UpdateAudioSessionProperties() {
	t._updateAudioSessionProperties()
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_updateAuxiliarySession
func (t TTSSpeechManager) _updateAuxiliarySession() {
	objc.Send[objc.ID](t.ID, objc.Sel("_updateAuxiliarySession"))
}

// UpdateAuxiliarySession is an exported wrapper for the private method _updateAuxiliarySession.
func (t TTSSpeechManager) UpdateAuxiliarySession() {
	t._updateAuxiliarySession()
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_updateUserSubstitutions
func (t TTSSpeechManager) _updateUserSubstitutions() {
	objc.Send[objc.ID](t.ID, objc.Sel("_updateUserSubstitutions"))
}

// UpdateUserSubstitutions is an exported wrapper for the private method _updateUserSubstitutions.
func (t TTSSpeechManager) UpdateUserSubstitutions() {
	t._updateUserSubstitutions()
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/clearSpeechQueue
func (t TTSSpeechManager) ClearSpeechQueue() {
	objc.Send[objc.ID](t.ID, objc.Sel("clearSpeechQueue"))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/continueSpeaking
func (t TTSSpeechManager) ContinueSpeaking() {
	objc.Send[objc.ID](t.ID, objc.Sel("continueSpeaking"))
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/dispatchSpeechAction:
func (t TTSSpeechManager) DispatchSpeechAction(action objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("dispatchSpeechAction:"), action)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/externalVoiceIdentifierUsedForLanguage:
func (t TTSSpeechManager) ExternalVoiceIdentifierUsedForLanguage(language objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("externalVoiceIdentifierUsedForLanguage:"), language)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/handleAudioInterruption:
func (t TTSSpeechManager) HandleAudioInterruption(interruption objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("handleAudioInterruption:"), interruption)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/handleAudioSessionObservers:
func (t TTSSpeechManager) HandleAudioSessionObservers(observers bool) {
	objc.Send[objc.ID](t.ID, objc.Sel("handleAudioSessionObservers:"), observers)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/handleMediaServicesWereLost:
func (t TTSSpeechManager) HandleMediaServicesWereLost(lost objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("handleMediaServicesWereLost:"), lost)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/handleMediaServicesWereReset:
func (t TTSSpeechManager) HandleMediaServicesWereReset(reset objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("handleMediaServicesWereReset:"), reset)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/pauseSpeaking:
func (t TTSSpeechManager) PauseSpeaking(speaking int64) {
	objc.Send[objc.ID](t.ID, objc.Sel("pauseSpeaking:"), speaking)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/setRequestWillStart:
func (t TTSSpeechManager) SetRequestWillStart(start VoidHandler) {
_block0, _ := NewVoidBlock(start)
	objc.Send[objc.ID](t.ID, objc.Sel("setRequestWillStart:"), _block0)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/speechSynthesizer:didContinueSpeakingRequest:
func (t TTSSpeechManager) SpeechSynthesizerDidContinueSpeakingRequest(synthesizer objectivec.IObject, request objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("speechSynthesizer:didContinueSpeakingRequest:"), synthesizer, request)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/speechSynthesizer:didEncounterMarker:forRequest:
func (t TTSSpeechManager) SpeechSynthesizerDidEncounterMarkerForRequest(synthesizer objectivec.IObject, marker objectivec.IObject, request objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("speechSynthesizer:didEncounterMarker:forRequest:"), synthesizer, marker, request)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/speechSynthesizer:didFinishSpeakingRequest:successfully:withError:
func (t TTSSpeechManager) SpeechSynthesizerDidFinishSpeakingRequestSuccessfullyWithError(synthesizer objectivec.IObject, request objectivec.IObject, successfully bool, error_ objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("speechSynthesizer:didFinishSpeakingRequest:successfully:withError:"), synthesizer, request, successfully, error_)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/speechSynthesizer:didPauseSpeakingRequest:
func (t TTSSpeechManager) SpeechSynthesizerDidPauseSpeakingRequest(synthesizer objectivec.IObject, request objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("speechSynthesizer:didPauseSpeakingRequest:"), synthesizer, request)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/speechSynthesizer:didStartSpeakingRequest:
func (t TTSSpeechManager) SpeechSynthesizerDidStartSpeakingRequest(synthesizer objectivec.IObject, request objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("speechSynthesizer:didStartSpeakingRequest:"), synthesizer, request)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/stopSpeaking
func (t TTSSpeechManager) StopSpeaking() {
	objc.Send[objc.ID](t.ID, objc.Sel("stopSpeaking"))
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/stopSpeaking:
func (t TTSSpeechManager) StopSpeakingWithSpeaking(speaking int64) {
	objc.Send[objc.ID](t.ID, objc.Sel("stopSpeaking:"), speaking)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/tearDown
func (t TTSSpeechManager) TearDown() {
	objc.Send[objc.ID](t.ID, objc.Sel("tearDown"))
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/voiceIdentifierUsedForLanguage:
func (t TTSSpeechManager) VoiceIdentifierUsedForLanguage(language objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("voiceIdentifierUsedForLanguage:"), language)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_isCharacterNativelySpeakable:languageCode:
func (_TTSSpeechManagerClass TTSSpeechManagerClass) _isCharacterNativelySpeakableLanguageCode(speakable uint16, code objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("_isCharacterNativelySpeakable:languageCode:"), speakable, code)
	return rv
}

// IsCharacterNativelySpeakableLanguageCode is an exported wrapper for the private method _isCharacterNativelySpeakableLanguageCode.
func (_TTSSpeechManagerClass TTSSpeechManagerClass) IsCharacterNativelySpeakableLanguageCode(speakable uint16, code objectivec.IObject) bool {
	return _TTSSpeechManagerClass._isCharacterNativelySpeakableLanguageCode(speakable, code)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_resetAvailableVoices
func (_TTSSpeechManagerClass TTSSpeechManagerClass) _resetAvailableVoices() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("_resetAvailableVoices"))
	return objectivec.Object{ID: rv}
}

// ResetAvailableVoices is an exported wrapper for the private method _resetAvailableVoices.
func (_TTSSpeechManagerClass TTSSpeechManagerClass) ResetAvailableVoices() objectivec.IObject {
	return _TTSSpeechManagerClass._resetAvailableVoices()
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/_resetAvailableVoices:
func (_TTSSpeechManagerClass TTSSpeechManagerClass) _resetAvailableVoicesWithVoices(voices bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("_resetAvailableVoices:"), voices)
	return objectivec.Object{ID: rv}
}

// ResetAvailableVoicesWithVoices is an exported wrapper for the private method _resetAvailableVoicesWithVoices.
func (_TTSSpeechManagerClass TTSSpeechManagerClass) ResetAvailableVoicesWithVoices(voices bool) objectivec.IObject {
	return _TTSSpeechManagerClass._resetAvailableVoicesWithVoices(voices)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/audioFileSettingsForVoice:
func (_TTSSpeechManagerClass TTSSpeechManagerClass) AudioFileSettingsForVoice(voice objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("audioFileSettingsForVoice:"), voice)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/avSpeechVoicesForTTSAXResources:
func (_TTSSpeechManagerClass TTSSpeechManagerClass) AvSpeechVoicesForTTSAXResources(tTSAXResources objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("avSpeechVoicesForTTSAXResources:"), tTSAXResources)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/availableLanguageCodes
func (_TTSSpeechManagerClass TTSSpeechManagerClass) AvailableLanguageCodes() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("availableLanguageCodes"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/availableSuperCompactVoices
func (_TTSSpeechManagerClass TTSSpeechManagerClass) AvailableSuperCompactVoices() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("availableSuperCompactVoices"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/availableVoices
func (_TTSSpeechManagerClass TTSSpeechManagerClass) AvailableVoices() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("availableVoices"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/availableVoices:
func (_TTSSpeechManagerClass TTSSpeechManagerClass) AvailableVoicesWithVoices(voices bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("availableVoices:"), voices)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/createRegularExpressionFromString:
func (_TTSSpeechManagerClass TTSSpeechManagerClass) CreateRegularExpressionFromString(string_ objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("createRegularExpressionFromString:"), string_)
	return rv
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/currentLanguageCode
func (_TTSSpeechManagerClass TTSSpeechManagerClass) CurrentLanguageCode() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("currentLanguageCode"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/currentProcessAllowedToSaveVoiceInfo
func (_TTSSpeechManagerClass TTSSpeechManagerClass) CurrentProcessAllowedToSaveVoiceInfo() bool {
	rv := objc.Send[bool](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("currentProcessAllowedToSaveVoiceInfo"))
	return rv
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/languageCodeForVoiceIdentifier:
func (_TTSSpeechManagerClass TTSSpeechManagerClass) LanguageCodeForVoiceIdentifier(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("languageCodeForVoiceIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/literalStringMarkup:string:speakCap:
func (_TTSSpeechManagerClass TTSSpeechManagerClass) LiteralStringMarkupStringSpeakCap(markup objectivec.IObject, string_ objectivec.IObject, cap_ bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("literalStringMarkup:string:speakCap:"), markup, string_, cap_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/matchedRangesForString:withRegularExpression:
func (_TTSSpeechManagerClass TTSSpeechManagerClass) MatchedRangesForStringWithRegularExpression(string_ objectivec.IObject, expression unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("matchedRangesForString:withRegularExpression:"), string_, expression)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/pauseMarkupString:
func (_TTSSpeechManagerClass TTSSpeechManagerClass) PauseMarkupString(string_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("pauseMarkupString:"), string_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/remapLanguageCode:
func (_TTSSpeechManagerClass TTSSpeechManagerClass) RemapLanguageCode(code objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("remapLanguageCode:"), code)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/spellOutLetterCaseMarkupString:string:
func (_TTSSpeechManagerClass TTSSpeechManagerClass) SpellOutLetterCaseMarkupStringString(string_ objectivec.IObject, string_2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("spellOutLetterCaseMarkupString:string:"), string_, string_2)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/spellOutMarkupString:string:
func (_TTSSpeechManagerClass TTSSpeechManagerClass) SpellOutMarkupStringString(string_ objectivec.IObject, string_2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("spellOutMarkupString:string:"), string_, string_2)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/test_actionStartTap:
func (_TTSSpeechManagerClass TTSSpeechManagerClass) Test_actionStartTap(tap VoidHandler) {
_block0, _ := NewVoidBlock(tap)
	objc.Send[objc.ID](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("test_actionStartTap:"), _block0)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/test_setAvailableVoices:
func (_TTSSpeechManagerClass TTSSpeechManagerClass) Test_setAvailableVoices(voices objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("test_setAvailableVoices:"), voices)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/test_setUnitTestMode:
func (_TTSSpeechManagerClass TTSSpeechManagerClass) Test_setUnitTestMode(mode bool) {
	objc.Send[objc.ID](objc.ID(_TTSSpeechManagerClass.class), objc.Sel("test_setUnitTestMode:"), mode)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/audioDeactivatorTimer
func (t TTSSpeechManager) AudioDeactivatorTimer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t.ID, objc.Sel("audioDeactivatorTimer"))
	return rv
}
func (t TTSSpeechManager) SetAudioDeactivatorTimer(value unsafe.Pointer) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioDeactivatorTimer:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/audioDeviceId
func (t TTSSpeechManager) AudioDeviceId() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("audioDeviceId"))
	return rv
}
func (t TTSSpeechManager) SetAudioDeviceId(value uint32) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioDeviceId:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/audioInterruptionStartedTime
func (t TTSSpeechManager) AudioInterruptionStartedTime() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("audioInterruptionStartedTime"))
	return rv
}
func (t TTSSpeechManager) SetAudioInterruptionStartedTime(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioInterruptionStartedTime:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/audioOperationQueue
func (t TTSSpeechManager) AudioOperationQueue() objectivec.Object {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("audioOperationQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (t TTSSpeechManager) SetAudioOperationQueue(value objectivec.Object) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioOperationQueue:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/audioQueueFlags
func (t TTSSpeechManager) AudioQueueFlags() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("audioQueueFlags"))
	return rv
}
func (t TTSSpeechManager) SetAudioQueueFlags(value uint32) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioQueueFlags:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/audioSession
func (t TTSSpeechManager) AudioSession() objc.ID {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("audioSession"))
	return rv
}
func (t TTSSpeechManager) SetAudioSession(value objc.ID) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioSession:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/audioSessionCategory
func (t TTSSpeechManager) AudioSessionCategory() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("audioSessionCategory"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechManager) SetAudioSessionCategory(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioSessionCategory:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/audioSessionCategoryOptions
func (t TTSSpeechManager) AudioSessionCategoryOptions() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("audioSessionCategoryOptions"))
	return rv
}
func (t TTSSpeechManager) SetAudioSessionCategoryOptions(value uint64) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioSessionCategoryOptions:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/audioSessionInactiveTimeout
func (t TTSSpeechManager) AudioSessionInactiveTimeout() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("audioSessionInactiveTimeout"))
	return rv
}
func (t TTSSpeechManager) SetAudioSessionInactiveTimeout(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioSessionInactiveTimeout:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/debugDescription
func (t TTSSpeechManager) DebugDescription() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/description
func (t TTSSpeechManager) Description() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/didRequestPauseSpeakingDuringAudioInterruption
func (t TTSSpeechManager) DidRequestPauseSpeakingDuringAudioInterruption() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("didRequestPauseSpeakingDuringAudioInterruption"))
	return rv
}
func (t TTSSpeechManager) SetDidRequestPauseSpeakingDuringAudioInterruption(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setDidRequestPauseSpeakingDuringAudioInterruption:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/didRequestResumeSpeakingDuringAudioInterruption
func (t TTSSpeechManager) DidRequestResumeSpeakingDuringAudioInterruption() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("didRequestResumeSpeakingDuringAudioInterruption"))
	return rv
}
func (t TTSSpeechManager) SetDidRequestResumeSpeakingDuringAudioInterruption(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setDidRequestResumeSpeakingDuringAudioInterruption:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/didRequestStartSpeakingDuringAudioInterruption
func (t TTSSpeechManager) DidRequestStartSpeakingDuringAudioInterruption() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("didRequestStartSpeakingDuringAudioInterruption"))
	return rv
}
func (t TTSSpeechManager) SetDidRequestStartSpeakingDuringAudioInterruption(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setDidRequestStartSpeakingDuringAudioInterruption:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/hash
func (t TTSSpeechManager) Hash() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/isInAudioInterruption
func (t TTSSpeechManager) IsInAudioInterruption() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isInAudioInterruption"))
	return rv
}
func (t TTSSpeechManager) SetIsInAudioInterruption(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setIsInAudioInterruption:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/isPaused
func (t TTSSpeechManager) IsPaused() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isPaused"))
	return rv
}
func (t TTSSpeechManager) SetIsPaused(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setIsPaused:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/isSpeaking
func (t TTSSpeechManager) IsSpeaking() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isSpeaking"))
	return rv
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/originalSpeechRateForJobOverride
func (t TTSSpeechManager) OriginalSpeechRateForJobOverride() foundation.NSNumber {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("originalSpeechRateForJobOverride"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (t TTSSpeechManager) SetOriginalSpeechRateForJobOverride(value foundation.NSNumber) {
	objc.Send[struct{}](t.ID, objc.Sel("setOriginalSpeechRateForJobOverride:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/outputChannels
func (t TTSSpeechManager) OutputChannels() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("outputChannels"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t TTSSpeechManager) SetOutputChannels(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setOutputChannels:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/requestedActionDuringAudioInterruption
func (t TTSSpeechManager) RequestedActionDuringAudioInterruption() ITTSSpeechAction {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("requestedActionDuringAudioInterruption"))
	return TTSSpeechActionFromID(objc.ID(rv))
}
func (t TTSSpeechManager) SetRequestedActionDuringAudioInterruption(value ITTSSpeechAction) {
	objc.Send[struct{}](t.ID, objc.Sel("setRequestedActionDuringAudioInterruption:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/setActiveOptions
func (t TTSSpeechManager) SetActiveOptions() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("setActiveOptions"))
	return rv
}
func (t TTSSpeechManager) SetSetActiveOptions(value uint64) {
	objc.Send[struct{}](t.ID, objc.Sel("setSetActiveOptions:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/shouldHandleAudioInterruptions
func (t TTSSpeechManager) ShouldHandleAudioInterruptions() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("shouldHandleAudioInterruptions"))
	return rv
}
func (t TTSSpeechManager) SetShouldHandleAudioInterruptions(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setShouldHandleAudioInterruptions:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/showControlCenterControls
func (t TTSSpeechManager) ShowControlCenterControls() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("showControlCenterControls"))
	return rv
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/speechEnabled
func (t TTSSpeechManager) SpeechEnabled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("speechEnabled"))
	return rv
}
func (t TTSSpeechManager) SetSpeechEnabled(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setSpeechEnabled:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/speechSource
func (t TTSSpeechManager) SpeechSource() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("speechSource"))
	return foundation.NSStringFromID(rv).String()
}
func (t TTSSpeechManager) SetSpeechSource(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setSpeechSource:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/superclass
func (t TTSSpeechManager) Superclass() objc.Class {
	rv := objc.Send[objc.Class](t.ID, objc.Sel("superclass"))
	return rv
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/usesAuxiliarySession
func (t TTSSpeechManager) UsesAuxiliarySession() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("usesAuxiliarySession"))
	return rv
}
func (t TTSSpeechManager) SetUsesAuxiliarySession(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setUsesAuxiliarySession:"), value)
}
// See: https://developer.apple.com/documentation/TextToSpeech/TTSSpeechManager/wasSpeakingBeforeAudioInterruption
func (t TTSSpeechManager) WasSpeakingBeforeAudioInterruption() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("wasSpeakingBeforeAudioInterruption"))
	return rv
}
func (t TTSSpeechManager) SetWasSpeakingBeforeAudioInterruption(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setWasSpeakingBeforeAudioInterruption:"), value)
}

// SetRequestWillStartSync is a synchronous wrapper around [TTSSpeechManager.SetRequestWillStart].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSSpeechManager) SetRequestWillStartSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.SetRequestWillStart(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Test_actionStartTapSync is a synchronous wrapper around [TTSSpeechManager.Test_actionStartTap].
// It blocks until the completion handler fires or the context is cancelled.
func (tc TTSSpeechManagerClass) Test_actionStartTapSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	tc.Test_actionStartTap(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

