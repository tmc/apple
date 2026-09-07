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

// The class instance for the [BFSpeechChannel] class.
var (
	_BFSpeechChannelClass     BFSpeechChannelClass
	_BFSpeechChannelClassOnce sync.Once
)

func getBFSpeechChannelClass() BFSpeechChannelClass {
	_BFSpeechChannelClassOnce.Do(func() {
		_BFSpeechChannelClass = BFSpeechChannelClass{class: objc.GetClass("BFSpeechChannel")}
	})
	return _BFSpeechChannelClass
}

// GetBFSpeechChannelClass returns the class object for BFSpeechChannel.
func GetBFSpeechChannelClass() BFSpeechChannelClass {
	return getBFSpeechChannelClass()
}

type BFSpeechChannelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (bc BFSpeechChannelClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BFSpeechChannelClass) Alloc() BFSpeechChannel {
	rv := objc.SendIfResponds[BFSpeechChannel](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [BFSpeechChannel._callSpeechDone]
//   - [BFSpeechChannel._closeAudioFile]
//   - [BFSpeechChannel._createExtFileForUrl]
//   - [BFSpeechChannel._forceResetChannel]
//   - [BFSpeechChannel._resetSynthesizer]
//   - [BFSpeechChannel._safelyCallPendingStopBlock]
//   - [BFSpeechChannel._subscribeToDefaultDeviceChanges]
//   - [BFSpeechChannel.AudioDeviceId]
//   - [BFSpeechChannel.SetAudioDeviceId]
//   - [BFSpeechChannel.AudioFile]
//   - [BFSpeechChannel.SetAudioFile]
//   - [BFSpeechChannel.BadStartingCharacterSet]
//   - [BFSpeechChannel.CallbackQueue]
//   - [BFSpeechChannel.SetCallbackQueue]
//   - [BFSpeechChannel.ChannelOperationQueue]
//   - [BFSpeechChannel.SetChannelOperationQueue]
//   - [BFSpeechChannel.ChannelState]
//   - [BFSpeechChannel.SetChannelState]
//   - [BFSpeechChannel.CharacterLiteralMode]
//   - [BFSpeechChannel.SetCharacterLiteralMode]
//   - [BFSpeechChannel.ContinueSpeech]
//   - [BFSpeechChannel.CurrentRequest]
//   - [BFSpeechChannel.SetCurrentRequest]
//   - [BFSpeechChannel.CurrentSpeechString]
//   - [BFSpeechChannel.SetCurrentSpeechString]
//   - [BFSpeechChannel.LegacyIdentifier]
//   - [BFSpeechChannel.SetLegacyIdentifier]
//   - [BFSpeechChannel.NumberLiteralMode]
//   - [BFSpeechChannel.SetNumberLiteralMode]
//   - [BFSpeechChannel.OutputUrl]
//   - [BFSpeechChannel.SetOutputUrl]
//   - [BFSpeechChannel.Pause]
//   - [BFSpeechChannel.Pitch]
//   - [BFSpeechChannel.SetPitch]
//   - [BFSpeechChannel.PitchMod]
//   - [BFSpeechChannel.SetPitchMod]
//   - [BFSpeechChannel.Rate]
//   - [BFSpeechChannel.SetRate]
//   - [BFSpeechChannel.RefCon]
//   - [BFSpeechChannel.SetRefCon]
//   - [BFSpeechChannel.Reset]
//   - [BFSpeechChannel.SetCfWordCallback]
//   - [BFSpeechChannel.SetErrorCallback]
//   - [BFSpeechChannel.SetPendingStopBlock]
//   - [BFSpeechChannel.SetPhonemeCallback]
//   - [BFSpeechChannel.SetSpeechDoneCallback]
//   - [BFSpeechChannel.SetSyncCallback]
//   - [BFSpeechChannel.SetWordCallback]
//   - [BFSpeechChannel.Speak]
//   - [BFSpeechChannel.SpeechActive]
//   - [BFSpeechChannel.SetSpeechActive]
//   - [BFSpeechChannel.SpeechFinished]
//   - [BFSpeechChannel.SetSpeechFinished]
//   - [BFSpeechChannel.SpeechPaused]
//   - [BFSpeechChannel.SetSpeechPaused]
//   - [BFSpeechChannel.SpeechSource]
//   - [BFSpeechChannel.SetSpeechSource]
//   - [BFSpeechChannel.SpeechSynthesizerDidContinueSpeakingRequest]
//   - [BFSpeechChannel.SpeechSynthesizerDidEncounterMarkerForRequest]
//   - [BFSpeechChannel.SpeechSynthesizerDidFinishSpeakingRequestSuccessfullyPhonemesSpokenWithError]
//   - [BFSpeechChannel.SpeechSynthesizerDidFinishSpeakingRequestSuccessfullyWithError]
//   - [BFSpeechChannel.SpeechSynthesizerDidPauseSpeakingRequest]
//   - [BFSpeechChannel.Stop]
//   - [BFSpeechChannel.Synthesizer]
//   - [BFSpeechChannel.SetSynthesizer]
//   - [BFSpeechChannel.SynthesizerShouldReset]
//   - [BFSpeechChannel.SetSynthesizerShouldReset]
//   - [BFSpeechChannel.Terminate]
//   - [BFSpeechChannel.Voice]
//   - [BFSpeechChannel.SetVoice]
//   - [BFSpeechChannel.Volume]
//   - [BFSpeechChannel.SetVolume]
//   - [BFSpeechChannel.InitWithVoice]
//   - [BFSpeechChannel.DebugDescription]
//   - [BFSpeechChannel.Description]
//   - [BFSpeechChannel.Hash]
//   - [BFSpeechChannel.Superclass]
type BFSpeechChannel struct {
	objectivec.Object
}

// BFSpeechChannelFromID constructs a [BFSpeechChannel] from an objc.ID.
func BFSpeechChannelFromID(id objc.ID) BFSpeechChannel {
	return BFSpeechChannel{objectivec.Object{ID: id}}
}

// Ensure BFSpeechChannel implements IBFSpeechChannel.
var _ IBFSpeechChannel = BFSpeechChannel{}

// An interface definition for the [BFSpeechChannel] class.
//
// # Methods
//
//   - [IBFSpeechChannel._callSpeechDone]
//   - [IBFSpeechChannel._closeAudioFile]
//   - [IBFSpeechChannel._createExtFileForUrl]
//   - [IBFSpeechChannel._forceResetChannel]
//   - [IBFSpeechChannel._resetSynthesizer]
//   - [IBFSpeechChannel._safelyCallPendingStopBlock]
//   - [IBFSpeechChannel._subscribeToDefaultDeviceChanges]
//   - [IBFSpeechChannel.AudioDeviceId]
//   - [IBFSpeechChannel.SetAudioDeviceId]
//   - [IBFSpeechChannel.AudioFile]
//   - [IBFSpeechChannel.SetAudioFile]
//   - [IBFSpeechChannel.BadStartingCharacterSet]
//   - [IBFSpeechChannel.CallbackQueue]
//   - [IBFSpeechChannel.SetCallbackQueue]
//   - [IBFSpeechChannel.ChannelOperationQueue]
//   - [IBFSpeechChannel.SetChannelOperationQueue]
//   - [IBFSpeechChannel.ChannelState]
//   - [IBFSpeechChannel.SetChannelState]
//   - [IBFSpeechChannel.CharacterLiteralMode]
//   - [IBFSpeechChannel.SetCharacterLiteralMode]
//   - [IBFSpeechChannel.ContinueSpeech]
//   - [IBFSpeechChannel.CurrentRequest]
//   - [IBFSpeechChannel.SetCurrentRequest]
//   - [IBFSpeechChannel.CurrentSpeechString]
//   - [IBFSpeechChannel.SetCurrentSpeechString]
//   - [IBFSpeechChannel.LegacyIdentifier]
//   - [IBFSpeechChannel.SetLegacyIdentifier]
//   - [IBFSpeechChannel.NumberLiteralMode]
//   - [IBFSpeechChannel.SetNumberLiteralMode]
//   - [IBFSpeechChannel.OutputUrl]
//   - [IBFSpeechChannel.SetOutputUrl]
//   - [IBFSpeechChannel.Pause]
//   - [IBFSpeechChannel.Pitch]
//   - [IBFSpeechChannel.SetPitch]
//   - [IBFSpeechChannel.PitchMod]
//   - [IBFSpeechChannel.SetPitchMod]
//   - [IBFSpeechChannel.Rate]
//   - [IBFSpeechChannel.SetRate]
//   - [IBFSpeechChannel.RefCon]
//   - [IBFSpeechChannel.SetRefCon]
//   - [IBFSpeechChannel.Reset]
//   - [IBFSpeechChannel.SetCfWordCallback]
//   - [IBFSpeechChannel.SetErrorCallback]
//   - [IBFSpeechChannel.SetPendingStopBlock]
//   - [IBFSpeechChannel.SetPhonemeCallback]
//   - [IBFSpeechChannel.SetSpeechDoneCallback]
//   - [IBFSpeechChannel.SetSyncCallback]
//   - [IBFSpeechChannel.SetWordCallback]
//   - [IBFSpeechChannel.Speak]
//   - [IBFSpeechChannel.SpeechActive]
//   - [IBFSpeechChannel.SetSpeechActive]
//   - [IBFSpeechChannel.SpeechFinished]
//   - [IBFSpeechChannel.SetSpeechFinished]
//   - [IBFSpeechChannel.SpeechPaused]
//   - [IBFSpeechChannel.SetSpeechPaused]
//   - [IBFSpeechChannel.SpeechSource]
//   - [IBFSpeechChannel.SetSpeechSource]
//   - [IBFSpeechChannel.SpeechSynthesizerDidContinueSpeakingRequest]
//   - [IBFSpeechChannel.SpeechSynthesizerDidEncounterMarkerForRequest]
//   - [IBFSpeechChannel.SpeechSynthesizerDidFinishSpeakingRequestSuccessfullyPhonemesSpokenWithError]
//   - [IBFSpeechChannel.SpeechSynthesizerDidFinishSpeakingRequestSuccessfullyWithError]
//   - [IBFSpeechChannel.SpeechSynthesizerDidPauseSpeakingRequest]
//   - [IBFSpeechChannel.Stop]
//   - [IBFSpeechChannel.Synthesizer]
//   - [IBFSpeechChannel.SetSynthesizer]
//   - [IBFSpeechChannel.SynthesizerShouldReset]
//   - [IBFSpeechChannel.SetSynthesizerShouldReset]
//   - [IBFSpeechChannel.Terminate]
//   - [IBFSpeechChannel.Voice]
//   - [IBFSpeechChannel.SetVoice]
//   - [IBFSpeechChannel.Volume]
//   - [IBFSpeechChannel.SetVolume]
//   - [IBFSpeechChannel.InitWithVoice]
//   - [IBFSpeechChannel.DebugDescription]
//   - [IBFSpeechChannel.Description]
//   - [IBFSpeechChannel.Hash]
//   - [IBFSpeechChannel.Superclass]
type IBFSpeechChannel interface {
	objectivec.IObject

	// Topic: Methods

	_callSpeechDone()
	_closeAudioFile()
	_createExtFileForUrl(url foundation.NSURL) OpaqueExtAudioFileRef
	_forceResetChannel()
	_resetSynthesizer()
	_safelyCallPendingStopBlock()
	_subscribeToDefaultDeviceChanges()
	AudioDeviceId() uint32
	SetAudioDeviceId(value uint32)
	AudioFile() OpaqueExtAudioFileRef
	SetAudioFile(value OpaqueExtAudioFileRef)
	BadStartingCharacterSet() objectivec.IObject
	CallbackQueue() objectivec.Object
	SetCallbackQueue(value objectivec.Object)
	ChannelOperationQueue() objectivec.Object
	SetChannelOperationQueue(value objectivec.Object)
	ChannelState() uint64
	SetChannelState(value uint64)
	CharacterLiteralMode() bool
	SetCharacterLiteralMode(value bool)
	ContinueSpeech() int16
	CurrentRequest() ITTSSpeechRequest
	SetCurrentRequest(value ITTSSpeechRequest)
	CurrentSpeechString() ITTSSpeechString
	SetCurrentSpeechString(value ITTSSpeechString)
	LegacyIdentifier() uint64
	SetLegacyIdentifier(value uint64)
	NumberLiteralMode() bool
	SetNumberLiteralMode(value bool)
	OutputUrl() foundation.NSURL
	SetOutputUrl(value foundation.NSURL)
	Pause(pause int32) int16
	Pitch() foundation.NSNumber
	SetPitch(value foundation.NSNumber)
	PitchMod() foundation.NSNumber
	SetPitchMod(value foundation.NSNumber)
	Rate() foundation.NSNumber
	SetRate(value foundation.NSNumber)
	RefCon() unsafe.Pointer
	SetRefCon(value unsafe.Pointer)
	Reset()
	SetCfWordCallback(callback VoidHandler)
	SetErrorCallback(callback VoidHandler)
	SetPendingStopBlock(block VoidHandler)
	SetPhonemeCallback(callback VoidHandler)
	SetSpeechDoneCallback(callback VoidHandler)
	SetSyncCallback(callback VoidHandler)
	SetWordCallback(callback VoidHandler)
	Speak(speak objectivec.IObject) int16
	SpeechActive() bool
	SetSpeechActive(value bool)
	SpeechFinished() foundation.NSCondition
	SetSpeechFinished(value foundation.NSCondition)
	SpeechPaused() bool
	SetSpeechPaused(value bool)
	SpeechSource() string
	SetSpeechSource(value string)
	SpeechSynthesizerDidContinueSpeakingRequest(synthesizer objectivec.IObject, request objectivec.IObject)
	SpeechSynthesizerDidEncounterMarkerForRequest(synthesizer objectivec.IObject, marker objectivec.IObject, request objectivec.IObject)
	SpeechSynthesizerDidFinishSpeakingRequestSuccessfullyPhonemesSpokenWithError(synthesizer objectivec.IObject, request objectivec.IObject, successfully bool, spoken objectivec.IObject, error_ objectivec.IObject)
	SpeechSynthesizerDidFinishSpeakingRequestSuccessfullyWithError(synthesizer objectivec.IObject, request objectivec.IObject, successfully bool, error_ objectivec.IObject)
	SpeechSynthesizerDidPauseSpeakingRequest(synthesizer objectivec.IObject, request objectivec.IObject)
	Stop(stop int32) int16
	Synthesizer() ITTSSpeechSynthesizer
	SetSynthesizer(value ITTSSpeechSynthesizer)
	SynthesizerShouldReset() bool
	SetSynthesizerShouldReset(value bool)
	Terminate()
	Voice() ITTSSpeechVoice
	SetVoice(value ITTSSpeechVoice)
	Volume() foundation.NSNumber
	SetVolume(value foundation.NSNumber)
	InitWithVoice(voice objectivec.IObject) BFSpeechChannel
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (b BFSpeechChannel) Init() BFSpeechChannel {
	rv := objc.SendIfResponds[BFSpeechChannel](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BFSpeechChannel) Autorelease() BFSpeechChannel {
	rv := objc.SendIfResponds[BFSpeechChannel](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBFSpeechChannel creates a new BFSpeechChannel instance.
func NewBFSpeechChannel() BFSpeechChannel {
	class := getBFSpeechChannelClass()
	rv := objc.SendIfResponds[BFSpeechChannel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewBFSpeechChannelWithVoice(voice objectivec.IObject) BFSpeechChannel {
	instance := getBFSpeechChannelClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithVoice:"), voice)
	return BFSpeechChannelFromID(rv)
}

func (b BFSpeechChannel) _callSpeechDone() {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("_callSpeechDone"))
}

// CallSpeechDone is an exported wrapper for the private method _callSpeechDone.
func (b BFSpeechChannel) CallSpeechDone() error {
	if !objc.RespondsToSelector(b.ID, objc.Sel("_callSpeechDone")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_callSpeechDone"}
		return err
	}
	b._callSpeechDone()
	return nil
}

// CanCallSpeechDone reports whether the receiver responds to the private selector _callSpeechDone.
func (b BFSpeechChannel) CanCallSpeechDone() bool {
	return objc.RespondsToSelector(b.ID, objc.Sel("_callSpeechDone"))
}
func (b BFSpeechChannel) _closeAudioFile() {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("_closeAudioFile"))
}

// CloseAudioFile is an exported wrapper for the private method _closeAudioFile.
func (b BFSpeechChannel) CloseAudioFile() error {
	if !objc.RespondsToSelector(b.ID, objc.Sel("_closeAudioFile")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_closeAudioFile"}
		return err
	}
	b._closeAudioFile()
	return nil
}

// CanCloseAudioFile reports whether the receiver responds to the private selector _closeAudioFile.
func (b BFSpeechChannel) CanCloseAudioFile() bool {
	return objc.RespondsToSelector(b.ID, objc.Sel("_closeAudioFile"))
}
func (b BFSpeechChannel) _createExtFileForUrl(url foundation.NSURL) OpaqueExtAudioFileRef {
	rv := objc.SendIfResponds[OpaqueExtAudioFileRef](b.ID, objc.Sel("_createExtFileForUrl:"), url)
	return OpaqueExtAudioFileRef(rv)
}

// CreateExtFileForUrl is an exported wrapper for the private method _createExtFileForUrl.
func (b BFSpeechChannel) CreateExtFileForUrl(url foundation.NSURL) (OpaqueExtAudioFileRef, error) {
	if !objc.RespondsToSelector(b.ID, objc.Sel("_createExtFileForUrl:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_createExtFileForUrl:"}
		return *new(OpaqueExtAudioFileRef), err
	}
	return b._createExtFileForUrl(url), nil
}

// CanCreateExtFileForUrl reports whether the receiver responds to the private selector _createExtFileForUrl:.
func (b BFSpeechChannel) CanCreateExtFileForUrl() bool {
	return objc.RespondsToSelector(b.ID, objc.Sel("_createExtFileForUrl:"))
}
func (b BFSpeechChannel) _forceResetChannel() {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("_forceResetChannel"))
}

// ForceResetChannel is an exported wrapper for the private method _forceResetChannel.
func (b BFSpeechChannel) ForceResetChannel() error {
	if !objc.RespondsToSelector(b.ID, objc.Sel("_forceResetChannel")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_forceResetChannel"}
		return err
	}
	b._forceResetChannel()
	return nil
}

// CanForceResetChannel reports whether the receiver responds to the private selector _forceResetChannel.
func (b BFSpeechChannel) CanForceResetChannel() bool {
	return objc.RespondsToSelector(b.ID, objc.Sel("_forceResetChannel"))
}
func (b BFSpeechChannel) _resetSynthesizer() {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("_resetSynthesizer"))
}

// ResetSynthesizer is an exported wrapper for the private method _resetSynthesizer.
func (b BFSpeechChannel) ResetSynthesizer() error {
	if !objc.RespondsToSelector(b.ID, objc.Sel("_resetSynthesizer")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_resetSynthesizer"}
		return err
	}
	b._resetSynthesizer()
	return nil
}

// CanResetSynthesizer reports whether the receiver responds to the private selector _resetSynthesizer.
func (b BFSpeechChannel) CanResetSynthesizer() bool {
	return objc.RespondsToSelector(b.ID, objc.Sel("_resetSynthesizer"))
}
func (b BFSpeechChannel) _safelyCallPendingStopBlock() {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("_safelyCallPendingStopBlock"))
}

// SafelyCallPendingStopBlock is an exported wrapper for the private method _safelyCallPendingStopBlock.
func (b BFSpeechChannel) SafelyCallPendingStopBlock() error {
	if !objc.RespondsToSelector(b.ID, objc.Sel("_safelyCallPendingStopBlock")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_safelyCallPendingStopBlock"}
		return err
	}
	b._safelyCallPendingStopBlock()
	return nil
}

// CanSafelyCallPendingStopBlock reports whether the receiver responds to the private selector _safelyCallPendingStopBlock.
func (b BFSpeechChannel) CanSafelyCallPendingStopBlock() bool {
	return objc.RespondsToSelector(b.ID, objc.Sel("_safelyCallPendingStopBlock"))
}
func (b BFSpeechChannel) _subscribeToDefaultDeviceChanges() {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("_subscribeToDefaultDeviceChanges"))
}

// SubscribeToDefaultDeviceChanges is an exported wrapper for the private method _subscribeToDefaultDeviceChanges.
func (b BFSpeechChannel) SubscribeToDefaultDeviceChanges() error {
	if !objc.RespondsToSelector(b.ID, objc.Sel("_subscribeToDefaultDeviceChanges")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_subscribeToDefaultDeviceChanges"}
		return err
	}
	b._subscribeToDefaultDeviceChanges()
	return nil
}

// CanSubscribeToDefaultDeviceChanges reports whether the receiver responds to the private selector _subscribeToDefaultDeviceChanges.
func (b BFSpeechChannel) CanSubscribeToDefaultDeviceChanges() bool {
	return objc.RespondsToSelector(b.ID, objc.Sel("_subscribeToDefaultDeviceChanges"))
}
func (b BFSpeechChannel) BadStartingCharacterSet() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("badStartingCharacterSet"))
	return objectivec.Object{ID: rv}
}
func (b BFSpeechChannel) ContinueSpeech() int16 {
	rv := objc.SendIfResponds[int16](b.ID, objc.Sel("continueSpeech"))
	return rv
}
func (b BFSpeechChannel) Pause(pause int32) int16 {
	rv := objc.SendIfResponds[int16](b.ID, objc.Sel("pause:"), pause)
	return rv
}
func (b BFSpeechChannel) Reset() {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("reset"))
}

var _bfspeechchannel_setcfwordcallback_p0_key byte

func (b BFSpeechChannel) SetCfWordCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("setCfWordCallback:"), _block0)
}

var _bfspeechchannel_seterrorcallback_p0_key byte

func (b BFSpeechChannel) SetErrorCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("setErrorCallback:"), _block0)
}

var _bfspeechchannel_setpendingstopblock_p0_key byte

func (b BFSpeechChannel) SetPendingStopBlock(block VoidHandler) {
	_block0, _ := NewVoidBlock(block)
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("setPendingStopBlock:"), _block0)
}

var _bfspeechchannel_setphonemecallback_p0_key byte

func (b BFSpeechChannel) SetPhonemeCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("setPhonemeCallback:"), _block0)
}

var _bfspeechchannel_setspeechdonecallback_p0_key byte

func (b BFSpeechChannel) SetSpeechDoneCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("setSpeechDoneCallback:"), _block0)
}

var _bfspeechchannel_setsynccallback_p0_key byte

func (b BFSpeechChannel) SetSyncCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("setSyncCallback:"), _block0)
}

var _bfspeechchannel_setwordcallback_p0_key byte

func (b BFSpeechChannel) SetWordCallback(callback VoidHandler) {
	_block0, _ := NewVoidBlock(callback)
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("setWordCallback:"), _block0)
}
func (b BFSpeechChannel) Speak(speak objectivec.IObject) int16 {
	rv := objc.SendIfResponds[int16](b.ID, objc.Sel("speak:"), speak)
	return rv
}
func (b BFSpeechChannel) SpeechSynthesizerDidContinueSpeakingRequest(synthesizer objectivec.IObject, request objectivec.IObject) {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("speechSynthesizer:didContinueSpeakingRequest:"), synthesizer, request)
}
func (b BFSpeechChannel) SpeechSynthesizerDidEncounterMarkerForRequest(synthesizer objectivec.IObject, marker objectivec.IObject, request objectivec.IObject) {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("speechSynthesizer:didEncounterMarker:forRequest:"), synthesizer, marker, request)
}
func (b BFSpeechChannel) SpeechSynthesizerDidFinishSpeakingRequestSuccessfullyPhonemesSpokenWithError(synthesizer objectivec.IObject, request objectivec.IObject, successfully bool, spoken objectivec.IObject, error_ objectivec.IObject) {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("speechSynthesizer:didFinishSpeakingRequest:successfully:phonemesSpoken:withError:"), synthesizer, request, successfully, spoken, error_)
}
func (b BFSpeechChannel) SpeechSynthesizerDidFinishSpeakingRequestSuccessfullyWithError(synthesizer objectivec.IObject, request objectivec.IObject, successfully bool, error_ objectivec.IObject) {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("speechSynthesizer:didFinishSpeakingRequest:successfully:withError:"), synthesizer, request, successfully, error_)
}
func (b BFSpeechChannel) SpeechSynthesizerDidPauseSpeakingRequest(synthesizer objectivec.IObject, request objectivec.IObject) {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("speechSynthesizer:didPauseSpeakingRequest:"), synthesizer, request)
}
func (b BFSpeechChannel) Stop(stop int32) int16 {
	rv := objc.SendIfResponds[int16](b.ID, objc.Sel("stop:"), stop)
	return rv
}
func (b BFSpeechChannel) Terminate() {
	objc.SendIfResponds[objc.ID](b.ID, objc.Sel("terminate"))
}
func (b BFSpeechChannel) InitWithVoice(voice objectivec.IObject) BFSpeechChannel {
	rv := objc.SendIfResponds[BFSpeechChannel](b.ID, objc.Sel("initWithVoice:"), voice)
	return rv
}

func (b BFSpeechChannel) AudioDeviceId() uint32 {
	rv := objc.SendIfResponds[uint32](b.ID, objc.Sel("audioDeviceId"))
	return rv
}
func (b BFSpeechChannel) SetAudioDeviceId(value uint32) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setAudioDeviceId:"), value)
}
func (b BFSpeechChannel) AudioFile() OpaqueExtAudioFileRef {
	rv := objc.SendIfResponds[OpaqueExtAudioFileRef](b.ID, objc.Sel("audioFile"))
	return OpaqueExtAudioFileRef(rv)
}
func (b BFSpeechChannel) SetAudioFile(value OpaqueExtAudioFileRef) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setAudioFile:"), value)
}
func (b BFSpeechChannel) CallbackQueue() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("callbackQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (b BFSpeechChannel) SetCallbackQueue(value objectivec.Object) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setCallbackQueue:"), value)
}
func (b BFSpeechChannel) ChannelOperationQueue() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("channelOperationQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (b BFSpeechChannel) SetChannelOperationQueue(value objectivec.Object) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setChannelOperationQueue:"), value)
}
func (b BFSpeechChannel) ChannelState() uint64 {
	rv := objc.SendIfResponds[uint64](b.ID, objc.Sel("channelState"))
	return rv
}
func (b BFSpeechChannel) SetChannelState(value uint64) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setChannelState:"), value)
}
func (b BFSpeechChannel) CharacterLiteralMode() bool {
	rv := objc.SendIfResponds[bool](b.ID, objc.Sel("characterLiteralMode"))
	return rv
}
func (b BFSpeechChannel) SetCharacterLiteralMode(value bool) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setCharacterLiteralMode:"), value)
}
func (b BFSpeechChannel) CurrentRequest() ITTSSpeechRequest {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("currentRequest"))
	return TTSSpeechRequestFromID(objc.ID(rv))
}
func (b BFSpeechChannel) SetCurrentRequest(value ITTSSpeechRequest) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setCurrentRequest:"), value)
}
func (b BFSpeechChannel) CurrentSpeechString() ITTSSpeechString {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("currentSpeechString"))
	return TTSSpeechStringFromID(objc.ID(rv))
}
func (b BFSpeechChannel) SetCurrentSpeechString(value ITTSSpeechString) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setCurrentSpeechString:"), value)
}
func (b BFSpeechChannel) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (b BFSpeechChannel) Description() string {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (b BFSpeechChannel) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](b.ID, objc.Sel("hash"))
	return rv
}
func (b BFSpeechChannel) LegacyIdentifier() uint64 {
	rv := objc.SendIfResponds[uint64](b.ID, objc.Sel("legacyIdentifier"))
	return rv
}
func (b BFSpeechChannel) SetLegacyIdentifier(value uint64) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setLegacyIdentifier:"), value)
}
func (b BFSpeechChannel) NumberLiteralMode() bool {
	rv := objc.SendIfResponds[bool](b.ID, objc.Sel("numberLiteralMode"))
	return rv
}
func (b BFSpeechChannel) SetNumberLiteralMode(value bool) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setNumberLiteralMode:"), value)
}
func (b BFSpeechChannel) OutputUrl() foundation.NSURL {
	rv := objc.SendIfResponds[foundation.NSURL](b.ID, objc.Sel("outputUrl"))
	return foundation.NSURL(rv)
}
func (b BFSpeechChannel) SetOutputUrl(value foundation.NSURL) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setOutputUrl:"), value)
}
func (b BFSpeechChannel) Pitch() foundation.NSNumber {
	rv := objc.SendIfResponds[foundation.NSNumber](b.ID, objc.Sel("pitch"))
	return foundation.NSNumber(rv)
}
func (b BFSpeechChannel) SetPitch(value foundation.NSNumber) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setPitch:"), value)
}
func (b BFSpeechChannel) PitchMod() foundation.NSNumber {
	rv := objc.SendIfResponds[foundation.NSNumber](b.ID, objc.Sel("pitchMod"))
	return foundation.NSNumber(rv)
}
func (b BFSpeechChannel) SetPitchMod(value foundation.NSNumber) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setPitchMod:"), value)
}
func (b BFSpeechChannel) Rate() foundation.NSNumber {
	rv := objc.SendIfResponds[foundation.NSNumber](b.ID, objc.Sel("rate"))
	return foundation.NSNumber(rv)
}
func (b BFSpeechChannel) SetRate(value foundation.NSNumber) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setRate:"), value)
}
func (b BFSpeechChannel) RefCon() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](b.ID, objc.Sel("refCon"))
	return rv
}
func (b BFSpeechChannel) SetRefCon(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setRefCon:"), value)
}
func (b BFSpeechChannel) SpeechActive() bool {
	rv := objc.SendIfResponds[bool](b.ID, objc.Sel("speechActive"))
	return rv
}
func (b BFSpeechChannel) SetSpeechActive(value bool) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setSpeechActive:"), value)
}
func (b BFSpeechChannel) SpeechFinished() foundation.NSCondition {
	rv := objc.SendIfResponds[foundation.NSCondition](b.ID, objc.Sel("speechFinished"))
	return foundation.NSCondition(rv)
}
func (b BFSpeechChannel) SetSpeechFinished(value foundation.NSCondition) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setSpeechFinished:"), value)
}
func (b BFSpeechChannel) SpeechPaused() bool {
	rv := objc.SendIfResponds[bool](b.ID, objc.Sel("speechPaused"))
	return rv
}
func (b BFSpeechChannel) SetSpeechPaused(value bool) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setSpeechPaused:"), value)
}
func (b BFSpeechChannel) SpeechSource() string {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("speechSource"))
	return foundation.NSStringFromID(rv).String()
}
func (b BFSpeechChannel) SetSpeechSource(value string) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setSpeechSource:"), objc.String(value))
}
func (b BFSpeechChannel) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](b.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (b BFSpeechChannel) Synthesizer() ITTSSpeechSynthesizer {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("synthesizer"))
	return TTSSpeechSynthesizerFromID(objc.ID(rv))
}
func (b BFSpeechChannel) SetSynthesizer(value ITTSSpeechSynthesizer) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setSynthesizer:"), value)
}
func (b BFSpeechChannel) SynthesizerShouldReset() bool {
	rv := objc.SendIfResponds[bool](b.ID, objc.Sel("synthesizerShouldReset"))
	return rv
}
func (b BFSpeechChannel) SetSynthesizerShouldReset(value bool) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setSynthesizerShouldReset:"), value)
}
func (b BFSpeechChannel) Voice() ITTSSpeechVoice {
	rv := objc.SendIfResponds[objc.ID](b.ID, objc.Sel("voice"))
	return TTSSpeechVoiceFromID(objc.ID(rv))
}
func (b BFSpeechChannel) SetVoice(value ITTSSpeechVoice) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setVoice:"), value)
}
func (b BFSpeechChannel) Volume() foundation.NSNumber {
	rv := objc.SendIfResponds[foundation.NSNumber](b.ID, objc.Sel("volume"))
	return foundation.NSNumber(rv)
}
func (b BFSpeechChannel) SetVolume(value foundation.NSNumber) {
	objc.SendIfResponds[struct{}](b.ID, objc.Sel("setVolume:"), value)
}

// SetCfWordCallbackSync is a synchronous wrapper around [BFSpeechChannel.SetCfWordCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (b BFSpeechChannel) SetCfWordCallbackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	b.SetCfWordCallback(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetErrorCallbackSync is a synchronous wrapper around [BFSpeechChannel.SetErrorCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (b BFSpeechChannel) SetErrorCallbackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	b.SetErrorCallback(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetPendingStopBlockSync is a synchronous wrapper around [BFSpeechChannel.SetPendingStopBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (b BFSpeechChannel) SetPendingStopBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	b.SetPendingStopBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetPhonemeCallbackSync is a synchronous wrapper around [BFSpeechChannel.SetPhonemeCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (b BFSpeechChannel) SetPhonemeCallbackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	b.SetPhonemeCallback(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetSpeechDoneCallbackSync is a synchronous wrapper around [BFSpeechChannel.SetSpeechDoneCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (b BFSpeechChannel) SetSpeechDoneCallbackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	b.SetSpeechDoneCallback(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetSyncCallbackSync is a synchronous wrapper around [BFSpeechChannel.SetSyncCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (b BFSpeechChannel) SetSyncCallbackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	b.SetSyncCallback(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SetWordCallbackSync is a synchronous wrapper around [BFSpeechChannel.SetWordCallback].
// It blocks until the completion handler fires or the context is cancelled.
func (b BFSpeechChannel) SetWordCallbackSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	b.SetWordCallback(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
