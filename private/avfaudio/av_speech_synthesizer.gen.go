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
func (s AVSpeechSynthesizer) Init() AVSpeechSynthesizer {
	rv := objc.Send[AVSpeechSynthesizer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSpeechSynthesizer) Autorelease() AVSpeechSynthesizer {
	rv := objc.Send[AVSpeechSynthesizer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSpeechSynthesizer creates a new AVSpeechSynthesizer instance.
func NewAVSpeechSynthesizer() AVSpeechSynthesizer {
	class := getAVSpeechSynthesizerClass()
	rv := objc.Send[AVSpeechSynthesizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/_applyWebKitBehaviors
func (s AVSpeechSynthesizer) _applyWebKitBehaviors() {
	objc.Send[objc.ID](s.ID, objc.Sel("_applyWebKitBehaviors"))
}

// ApplyWebKitBehaviors is an exported wrapper for the private method _applyWebKitBehaviors.
func (s AVSpeechSynthesizer) ApplyWebKitBehaviors() {
	s._applyWebKitBehaviors()
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/_convertBoundary:
func (s AVSpeechSynthesizer) _convertBoundary(boundary int64) int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("_convertBoundary:"), boundary)
	return rv
}

// ConvertBoundary is an exported wrapper for the private method _convertBoundary.
func (s AVSpeechSynthesizer) ConvertBoundary(boundary int64) int64 {
	return s._convertBoundary(boundary)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/_enqueueNextJob
func (s AVSpeechSynthesizer) _enqueueNextJob() {
	objc.Send[objc.ID](s.ID, objc.Sel("_enqueueNextJob"))
}

// EnqueueNextJob is an exported wrapper for the private method _enqueueNextJob.
func (s AVSpeechSynthesizer) EnqueueNextJob() {
	s._enqueueNextJob()
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/_handleSpeechDone:successful:
func (s AVSpeechSynthesizer) _handleSpeechDoneSuccessful(done objectivec.IObject, successful bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("_handleSpeechDone:successful:"), done, successful)
}

// HandleSpeechDoneSuccessful is an exported wrapper for the private method _handleSpeechDoneSuccessful.
func (s AVSpeechSynthesizer) HandleSpeechDoneSuccessful(done objectivec.IObject, successful bool) {
	s._handleSpeechDoneSuccessful(done, successful)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/_speakUtterance:
func (s AVSpeechSynthesizer) _speakUtterance(utterance objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("_speakUtterance:"), utterance)
}

// SpeakUtterance is an exported wrapper for the private method _speakUtterance.
func (s AVSpeechSynthesizer) SpeakUtterance(utterance objectivec.IObject) {
	s._speakUtterance(utterance)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/audioQueueFlags
func (s AVSpeechSynthesizer) AudioQueueFlags() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("audioQueueFlags"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/coreSynth
func (s AVSpeechSynthesizer) CoreSynth() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("coreSynth"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/isInAudioInterruption
func (s AVSpeechSynthesizer) IsInAudioInterruption() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isInAudioInterruption"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/processSpeechJobFinished:successful:
func (s AVSpeechSynthesizer) ProcessSpeechJobFinishedSuccessful(finished objectivec.IObject, successful bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("processSpeechJobFinished:successful:"), finished, successful)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/setActiveOptions
func (s AVSpeechSynthesizer) SetActiveOptions() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("setActiveOptions"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/setAudioQueueFlags:
func (s AVSpeechSynthesizer) SetAudioQueueFlags(flags uint32) {
	objc.Send[objc.ID](s.ID, objc.Sel("setAudioQueueFlags:"), flags)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/setAudioSessionInactiveTimeout:
func (s AVSpeechSynthesizer) SetAudioSessionInactiveTimeout(timeout float64) {
	objc.Send[objc.ID](s.ID, objc.Sel("setAudioSessionInactiveTimeout:"), timeout)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/setSetActiveOptions:
func (s AVSpeechSynthesizer) SetSetActiveOptions(options uint64) {
	objc.Send[objc.ID](s.ID, objc.Sel("setSetActiveOptions:"), options)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/setSkipLuthorRules:
func (s AVSpeechSynthesizer) SetSkipLuthorRules(rules objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("setSkipLuthorRules:"), rules)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/setSupportsAccurateWordCallbacks:
func (s AVSpeechSynthesizer) SetSupportsAccurateWordCallbacks(callbacks objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("setSupportsAccurateWordCallbacks:"), callbacks)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/skipLuthorRules
func (s AVSpeechSynthesizer) SkipLuthorRules() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("skipLuthorRules"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/speechManager
func (s AVSpeechSynthesizer) SpeechManager() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("speechManager"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/speechQueue
func (s AVSpeechSynthesizer) SpeechQueue() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("speechQueue"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/supportsAccurateWordCallbacks
func (s AVSpeechSynthesizer) SupportsAccurateWordCallbacks() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("supportsAccurateWordCallbacks"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/_supportsSpeakingWithPersonalVoices
func (_AVSpeechSynthesizerClass AVSpeechSynthesizerClass) _supportsSpeakingWithPersonalVoices() bool {
	rv := objc.Send[bool](objc.ID(_AVSpeechSynthesizerClass.class), objc.Sel("_supportsSpeakingWithPersonalVoices"))
	return rv
}

// SupportsSpeakingWithPersonalVoices is an exported wrapper for the private method _supportsSpeakingWithPersonalVoices.
func (_AVSpeechSynthesizerClass AVSpeechSynthesizerClass) SupportsSpeakingWithPersonalVoices() bool {
	return _AVSpeechSynthesizerClass._supportsSpeakingWithPersonalVoices()
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/isSoftAppUsageProtectionDisabled
func (_AVSpeechSynthesizerClass AVSpeechSynthesizerClass) IsSoftAppUsageProtectionDisabled() bool {
	rv := objc.Send[bool](objc.ID(_AVSpeechSynthesizerClass.class), objc.Sel("isSoftAppUsageProtectionDisabled"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/audioDeviceId
func (s AVSpeechSynthesizer) AudioDeviceId() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("audioDeviceId"))
	return rv
}
func (s AVSpeechSynthesizer) SetAudioDeviceId(value uint32) {
	objc.Send[struct{}](s.ID, objc.Sel("setAudioDeviceId:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/detectSSMLAndModifyUtterances
func (s AVSpeechSynthesizer) DetectSSMLAndModifyUtterances() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("detectSSMLAndModifyUtterances"))
	return rv
}
func (s AVSpeechSynthesizer) SetDetectSSMLAndModifyUtterances(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setDetectSSMLAndModifyUtterances:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/inflightUtterance
func (s AVSpeechSynthesizer) InflightUtterance() IAVSpeechUtterance {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("inflightUtterance"))
	return AVSpeechUtteranceFromID(objc.ID(rv))
}
func (s AVSpeechSynthesizer) SetInflightUtterance(value IAVSpeechUtterance) {
	objc.Send[struct{}](s.ID, objc.Sel("setInflightUtterance:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/initializedWebKitUsage
func (s AVSpeechSynthesizer) InitializedWebKitUsage() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("initializedWebKitUsage"))
	return rv
}
func (s AVSpeechSynthesizer) SetInitializedWebKitUsage(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setInitializedWebKitUsage:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/isInternalSynth
func (s AVSpeechSynthesizer) IsInternalSynth() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isInternalSynth"))
	return rv
}
func (s AVSpeechSynthesizer) SetIsInternalSynth(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setIsInternalSynth:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/paused
func (s AVSpeechSynthesizer) Paused() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("paused"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/speaking
func (s AVSpeechSynthesizer) Speaking() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("speaking"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/speechSource
func (s AVSpeechSynthesizer) SpeechSource() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("speechSource"))
	return foundation.NSStringFromID(rv).String()
}
func (s AVSpeechSynthesizer) SetSpeechSource(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setSpeechSource:"), objc.String(value))
}
