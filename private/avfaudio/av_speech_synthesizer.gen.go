// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSpeechSynthesizer] class.
var (
	_AVSpeechSynthesizerClass     AVSpeechSynthesizerClass
	_AVSpeechSynthesizerClassOnce sync.Once
)

func getAVSpeechSynthesizerClass() AVSpeechSynthesizerClass {
	_AVSpeechSynthesizerClassOnce.Do(func() {
		_AVSpeechSynthesizerClass = AVSpeechSynthesizerClass{class: objc.GetClass("AVSpeechSynthesizer")}
	})
	return _AVSpeechSynthesizerClass
}

// GetAVSpeechSynthesizerClass returns the class object for AVSpeechSynthesizer.
func GetAVSpeechSynthesizerClass() AVSpeechSynthesizerClass {
	return getAVSpeechSynthesizerClass()
}

type AVSpeechSynthesizerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSpeechSynthesizerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSpeechSynthesizerClass) Alloc() AVSpeechSynthesizer {
	rv := objc.Send[AVSpeechSynthesizer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVSpeechSynthesizer._applyWebKitBehaviors]
//   - [AVSpeechSynthesizer._convertBoundary]
//   - [AVSpeechSynthesizer._enqueueNextJob]
//   - [AVSpeechSynthesizer._handleSpeechDoneSuccessful]
//   - [AVSpeechSynthesizer._speakUtterance]
//   - [AVSpeechSynthesizer.AudioDeviceId]
//   - [AVSpeechSynthesizer.SetAudioDeviceId]
//   - [AVSpeechSynthesizer.AudioQueueFlags]
//   - [AVSpeechSynthesizer.CoreSynth]
//   - [AVSpeechSynthesizer.DetectSSMLAndModifyUtterances]
//   - [AVSpeechSynthesizer.SetDetectSSMLAndModifyUtterances]
//   - [AVSpeechSynthesizer.InflightUtterance]
//   - [AVSpeechSynthesizer.SetInflightUtterance]
//   - [AVSpeechSynthesizer.InitializedWebKitUsage]
//   - [AVSpeechSynthesizer.SetInitializedWebKitUsage]
//   - [AVSpeechSynthesizer.IsInAudioInterruption]
//   - [AVSpeechSynthesizer.IsInternalSynth]
//   - [AVSpeechSynthesizer.SetIsInternalSynth]
//   - [AVSpeechSynthesizer.ProcessSpeechJobFinishedSuccessful]
//   - [AVSpeechSynthesizer.SetActiveOptions]
//   - [AVSpeechSynthesizer.SetAudioQueueFlags]
//   - [AVSpeechSynthesizer.SetAudioSessionInactiveTimeout]
//   - [AVSpeechSynthesizer.SetSetActiveOptions]
//   - [AVSpeechSynthesizer.SetSkipLuthorRules]
//   - [AVSpeechSynthesizer.SetSupportsAccurateWordCallbacks]
//   - [AVSpeechSynthesizer.SkipLuthorRules]
//   - [AVSpeechSynthesizer.SpeechManager]
//   - [AVSpeechSynthesizer.SpeechQueue]
//   - [AVSpeechSynthesizer.SpeechSource]
//   - [AVSpeechSynthesizer.SetSpeechSource]
//   - [AVSpeechSynthesizer.SupportsAccurateWordCallbacks]
//   - [AVSpeechSynthesizer.Paused]
//   - [AVSpeechSynthesizer.Speaking]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer
type AVSpeechSynthesizer struct {
	objectivec.Object
}

// AVSpeechSynthesizerFromID constructs a [AVSpeechSynthesizer] from an objc.ID.
func AVSpeechSynthesizerFromID(id objc.ID) AVSpeechSynthesizer {
	return AVSpeechSynthesizer{objectivec.Object{ID: id}}
}

// Ensure AVSpeechSynthesizer implements IAVSpeechSynthesizer.
var _ IAVSpeechSynthesizer = AVSpeechSynthesizer{}

// An interface definition for the [AVSpeechSynthesizer] class.
//
// # Methods
//
//   - [IAVSpeechSynthesizer._applyWebKitBehaviors]
//   - [IAVSpeechSynthesizer._convertBoundary]
//   - [IAVSpeechSynthesizer._enqueueNextJob]
//   - [IAVSpeechSynthesizer._handleSpeechDoneSuccessful]
//   - [IAVSpeechSynthesizer._speakUtterance]
//   - [IAVSpeechSynthesizer.AudioDeviceId]
//   - [IAVSpeechSynthesizer.SetAudioDeviceId]
//   - [IAVSpeechSynthesizer.AudioQueueFlags]
//   - [IAVSpeechSynthesizer.CoreSynth]
//   - [IAVSpeechSynthesizer.DetectSSMLAndModifyUtterances]
//   - [IAVSpeechSynthesizer.SetDetectSSMLAndModifyUtterances]
//   - [IAVSpeechSynthesizer.InflightUtterance]
//   - [IAVSpeechSynthesizer.SetInflightUtterance]
//   - [IAVSpeechSynthesizer.InitializedWebKitUsage]
//   - [IAVSpeechSynthesizer.SetInitializedWebKitUsage]
//   - [IAVSpeechSynthesizer.IsInAudioInterruption]
//   - [IAVSpeechSynthesizer.IsInternalSynth]
//   - [IAVSpeechSynthesizer.SetIsInternalSynth]
//   - [IAVSpeechSynthesizer.ProcessSpeechJobFinishedSuccessful]
//   - [IAVSpeechSynthesizer.SetActiveOptions]
//   - [IAVSpeechSynthesizer.SetAudioQueueFlags]
//   - [IAVSpeechSynthesizer.SetAudioSessionInactiveTimeout]
//   - [IAVSpeechSynthesizer.SetSetActiveOptions]
//   - [IAVSpeechSynthesizer.SetSkipLuthorRules]
//   - [IAVSpeechSynthesizer.SetSupportsAccurateWordCallbacks]
//   - [IAVSpeechSynthesizer.SkipLuthorRules]
//   - [IAVSpeechSynthesizer.SpeechManager]
//   - [IAVSpeechSynthesizer.SpeechQueue]
//   - [IAVSpeechSynthesizer.SpeechSource]
//   - [IAVSpeechSynthesizer.SetSpeechSource]
//   - [IAVSpeechSynthesizer.SupportsAccurateWordCallbacks]
//   - [IAVSpeechSynthesizer.Paused]
//   - [IAVSpeechSynthesizer.Speaking]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer
type IAVSpeechSynthesizer interface {
	objectivec.IObject

	// Topic: Methods

	_applyWebKitBehaviors()
	_convertBoundary(boundary int64) int64
	_enqueueNextJob()
	_handleSpeechDoneSuccessful(done objectivec.IObject, successful bool)
	_speakUtterance(utterance objectivec.IObject)
	AudioDeviceId() uint32
	SetAudioDeviceId(value uint32)
	AudioQueueFlags() uint32
	CoreSynth() objectivec.IObject
	DetectSSMLAndModifyUtterances() bool
	SetDetectSSMLAndModifyUtterances(value bool)
	InflightUtterance() IAVSpeechUtterance
	SetInflightUtterance(value IAVSpeechUtterance)
	InitializedWebKitUsage() bool
	SetInitializedWebKitUsage(value bool)
	IsInAudioInterruption() bool
	IsInternalSynth() bool
	SetIsInternalSynth(value bool)
	ProcessSpeechJobFinishedSuccessful(finished objectivec.IObject, successful bool)
	SetActiveOptions() uint64
	SetAudioQueueFlags(flags uint32)
	SetAudioSessionInactiveTimeout(timeout float64)
	SetSetActiveOptions(options uint64)
	SetSkipLuthorRules(rules objectivec.IObject)
	SetSupportsAccurateWordCallbacks(callbacks objectivec.IObject)
	SkipLuthorRules() objectivec.IObject
	SpeechManager() objectivec.IObject
	SpeechQueue() objectivec.IObject
	SpeechSource() string
	SetSpeechSource(value string)
	SupportsAccurateWordCallbacks() objectivec.IObject
	Paused() bool
	Speaking() bool
}

// Init initializes the instance.
func (a AVSpeechSynthesizer) Init() AVSpeechSynthesizer {
	rv := objc.Send[AVSpeechSynthesizer](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVSpeechSynthesizer) Autorelease() AVSpeechSynthesizer {
	rv := objc.Send[AVSpeechSynthesizer](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSpeechSynthesizer creates a new AVSpeechSynthesizer instance.
func NewAVSpeechSynthesizer() AVSpeechSynthesizer {
	class := getAVSpeechSynthesizerClass()
	rv := objc.Send[AVSpeechSynthesizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/_applyWebKitBehaviors
func (a AVSpeechSynthesizer) _applyWebKitBehaviors() {
	objc.Send[objc.ID](a.ID, objc.Sel("_applyWebKitBehaviors"))
}

// ApplyWebKitBehaviors is an exported wrapper for the private method _applyWebKitBehaviors.
func (a AVSpeechSynthesizer) ApplyWebKitBehaviors() error {
	if !objc.RespondsToSelector(a.ID, objc.Sel("_applyWebKitBehaviors")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_applyWebKitBehaviors"}
		return err
	}
	a._applyWebKitBehaviors()
	return nil
}

// CanApplyWebKitBehaviors reports whether the receiver responds to the private selector _applyWebKitBehaviors.
func (a AVSpeechSynthesizer) CanApplyWebKitBehaviors() bool {
	return objc.RespondsToSelector(a.ID, objc.Sel("_applyWebKitBehaviors"))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/_convertBoundary:
func (a AVSpeechSynthesizer) _convertBoundary(boundary int64) int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("_convertBoundary:"), boundary)
	return rv
}

// ConvertBoundary is an exported wrapper for the private method _convertBoundary.
func (a AVSpeechSynthesizer) ConvertBoundary(boundary int64) (int64, error) {
	if !objc.RespondsToSelector(a.ID, objc.Sel("_convertBoundary:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_convertBoundary:"}
		return 0, err
	}
	return a._convertBoundary(boundary), nil
}

// CanConvertBoundary reports whether the receiver responds to the private selector _convertBoundary:.
func (a AVSpeechSynthesizer) CanConvertBoundary() bool {
	return objc.RespondsToSelector(a.ID, objc.Sel("_convertBoundary:"))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/_enqueueNextJob
func (a AVSpeechSynthesizer) _enqueueNextJob() {
	objc.Send[objc.ID](a.ID, objc.Sel("_enqueueNextJob"))
}

// EnqueueNextJob is an exported wrapper for the private method _enqueueNextJob.
func (a AVSpeechSynthesizer) EnqueueNextJob() error {
	if !objc.RespondsToSelector(a.ID, objc.Sel("_enqueueNextJob")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_enqueueNextJob"}
		return err
	}
	a._enqueueNextJob()
	return nil
}

// CanEnqueueNextJob reports whether the receiver responds to the private selector _enqueueNextJob.
func (a AVSpeechSynthesizer) CanEnqueueNextJob() bool {
	return objc.RespondsToSelector(a.ID, objc.Sel("_enqueueNextJob"))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/_handleSpeechDone:successful:
func (a AVSpeechSynthesizer) _handleSpeechDoneSuccessful(done objectivec.IObject, successful bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("_handleSpeechDone:successful:"), done, successful)
}

// HandleSpeechDoneSuccessful is an exported wrapper for the private method _handleSpeechDoneSuccessful.
func (a AVSpeechSynthesizer) HandleSpeechDoneSuccessful(done objectivec.IObject, successful bool) error {
	if !objc.RespondsToSelector(a.ID, objc.Sel("_handleSpeechDone:successful:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_handleSpeechDone:successful:"}
		return err
	}
	a._handleSpeechDoneSuccessful(done, successful)
	return nil
}

// CanHandleSpeechDoneSuccessful reports whether the receiver responds to the private selector _handleSpeechDone:successful:.
func (a AVSpeechSynthesizer) CanHandleSpeechDoneSuccessful() bool {
	return objc.RespondsToSelector(a.ID, objc.Sel("_handleSpeechDone:successful:"))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/_speakUtterance:
func (a AVSpeechSynthesizer) _speakUtterance(utterance objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("_speakUtterance:"), utterance)
}

// SpeakUtterance is an exported wrapper for the private method _speakUtterance.
func (a AVSpeechSynthesizer) SpeakUtterance(utterance objectivec.IObject) error {
	if !objc.RespondsToSelector(a.ID, objc.Sel("_speakUtterance:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_speakUtterance:"}
		return err
	}
	a._speakUtterance(utterance)
	return nil
}

// CanSpeakUtterance reports whether the receiver responds to the private selector _speakUtterance:.
func (a AVSpeechSynthesizer) CanSpeakUtterance() bool {
	return objc.RespondsToSelector(a.ID, objc.Sel("_speakUtterance:"))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/audioQueueFlags
func (a AVSpeechSynthesizer) AudioQueueFlags() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("audioQueueFlags"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/coreSynth
func (a AVSpeechSynthesizer) CoreSynth() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("coreSynth"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/isInAudioInterruption
func (a AVSpeechSynthesizer) IsInAudioInterruption() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isInAudioInterruption"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/processSpeechJobFinished:successful:
func (a AVSpeechSynthesizer) ProcessSpeechJobFinishedSuccessful(finished objectivec.IObject, successful bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("processSpeechJobFinished:successful:"), finished, successful)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/setActiveOptions
func (a AVSpeechSynthesizer) SetActiveOptions() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("setActiveOptions"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/setAudioQueueFlags:
func (a AVSpeechSynthesizer) SetAudioQueueFlags(flags uint32) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAudioQueueFlags:"), flags)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/setAudioSessionInactiveTimeout:
func (a AVSpeechSynthesizer) SetAudioSessionInactiveTimeout(timeout float64) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAudioSessionInactiveTimeout:"), timeout)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/setSetActiveOptions:
func (a AVSpeechSynthesizer) SetSetActiveOptions(options uint64) {
	objc.Send[objc.ID](a.ID, objc.Sel("setSetActiveOptions:"), options)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/setSkipLuthorRules:
func (a AVSpeechSynthesizer) SetSkipLuthorRules(rules objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setSkipLuthorRules:"), rules)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/setSupportsAccurateWordCallbacks:
func (a AVSpeechSynthesizer) SetSupportsAccurateWordCallbacks(callbacks objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setSupportsAccurateWordCallbacks:"), callbacks)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/skipLuthorRules
func (a AVSpeechSynthesizer) SkipLuthorRules() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("skipLuthorRules"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/speechManager
func (a AVSpeechSynthesizer) SpeechManager() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("speechManager"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/speechQueue
func (a AVSpeechSynthesizer) SpeechQueue() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("speechQueue"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/supportsAccurateWordCallbacks
func (a AVSpeechSynthesizer) SupportsAccurateWordCallbacks() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("supportsAccurateWordCallbacks"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/_supportsSpeakingWithPersonalVoices
func (_AVSpeechSynthesizerClass AVSpeechSynthesizerClass) _supportsSpeakingWithPersonalVoices() bool {
	rv := objc.Send[bool](objc.ID(_AVSpeechSynthesizerClass.class), objc.Sel("_supportsSpeakingWithPersonalVoices"))
	return rv
}

// SupportsSpeakingWithPersonalVoices is an exported wrapper for the private method _supportsSpeakingWithPersonalVoices.
func (_AVSpeechSynthesizerClass AVSpeechSynthesizerClass) SupportsSpeakingWithPersonalVoices() (bool, error) {
	if !objc.RespondsToSelector(objc.ID(_AVSpeechSynthesizerClass.class), objc.Sel("_supportsSpeakingWithPersonalVoices")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_supportsSpeakingWithPersonalVoices"}
		return false, err
	}
	return _AVSpeechSynthesizerClass._supportsSpeakingWithPersonalVoices(), nil
}

// CanSupportsSpeakingWithPersonalVoices reports whether the receiver responds to the private selector _supportsSpeakingWithPersonalVoices.
func (_AVSpeechSynthesizerClass AVSpeechSynthesizerClass) CanSupportsSpeakingWithPersonalVoices() bool {
	return objc.RespondsToSelector(objc.ID(_AVSpeechSynthesizerClass.class), objc.Sel("_supportsSpeakingWithPersonalVoices"))
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/isSoftAppUsageProtectionDisabled
func (_AVSpeechSynthesizerClass AVSpeechSynthesizerClass) IsSoftAppUsageProtectionDisabled() bool {
	rv := objc.Send[bool](objc.ID(_AVSpeechSynthesizerClass.class), objc.Sel("isSoftAppUsageProtectionDisabled"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/audioDeviceId
func (a AVSpeechSynthesizer) AudioDeviceId() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("audioDeviceId"))
	return rv
}
func (a AVSpeechSynthesizer) SetAudioDeviceId(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setAudioDeviceId:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/detectSSMLAndModifyUtterances
func (a AVSpeechSynthesizer) DetectSSMLAndModifyUtterances() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("detectSSMLAndModifyUtterances"))
	return rv
}
func (a AVSpeechSynthesizer) SetDetectSSMLAndModifyUtterances(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setDetectSSMLAndModifyUtterances:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/inflightUtterance
func (a AVSpeechSynthesizer) InflightUtterance() IAVSpeechUtterance {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inflightUtterance"))
	return AVSpeechUtteranceFromID(objc.ID(rv))
}
func (a AVSpeechSynthesizer) SetInflightUtterance(value IAVSpeechUtterance) {
	objc.Send[struct{}](a.ID, objc.Sel("setInflightUtterance:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/initializedWebKitUsage
func (a AVSpeechSynthesizer) InitializedWebKitUsage() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("initializedWebKitUsage"))
	return rv
}
func (a AVSpeechSynthesizer) SetInitializedWebKitUsage(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setInitializedWebKitUsage:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/isInternalSynth
func (a AVSpeechSynthesizer) IsInternalSynth() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isInternalSynth"))
	return rv
}
func (a AVSpeechSynthesizer) SetIsInternalSynth(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setIsInternalSynth:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/paused
func (a AVSpeechSynthesizer) Paused() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("paused"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/speaking
func (a AVSpeechSynthesizer) Speaking() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("speaking"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/speechSource
func (a AVSpeechSynthesizer) SpeechSource() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("speechSource"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVSpeechSynthesizer) SetSpeechSource(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setSpeechSource:"), objc.String(value))
}
