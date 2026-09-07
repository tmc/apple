// Code generated from Apple documentation for avfaudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

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
	rv := objc.SendIfResponds[AVSpeechSynthesizer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVSpeechSynthesizer._configureCoreSynthesizer]
//   - [AVSpeechSynthesizer.AudioDeviceId]
//   - [AVSpeechSynthesizer.SetAudioDeviceId]
//   - [AVSpeechSynthesizer.AudioQueueFlags]
//   - [AVSpeechSynthesizer.CoreSynthesizer]
//   - [AVSpeechSynthesizer.ExistingCoreSynthesizer]
//   - [AVSpeechSynthesizer.IsInternalSynth]
//   - [AVSpeechSynthesizer.SetIsInternalSynth]
//   - [AVSpeechSynthesizer.SetActiveOptions]
//   - [AVSpeechSynthesizer.SetAudioQueueFlags]
//   - [AVSpeechSynthesizer.SetSetActiveOptions]
//   - [AVSpeechSynthesizer.Paused]
//   - [AVSpeechSynthesizer.Speaking]
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
//   - [IAVSpeechSynthesizer._configureCoreSynthesizer]
//   - [IAVSpeechSynthesizer.AudioDeviceId]
//   - [IAVSpeechSynthesizer.SetAudioDeviceId]
//   - [IAVSpeechSynthesizer.AudioQueueFlags]
//   - [IAVSpeechSynthesizer.CoreSynthesizer]
//   - [IAVSpeechSynthesizer.ExistingCoreSynthesizer]
//   - [IAVSpeechSynthesizer.IsInternalSynth]
//   - [IAVSpeechSynthesizer.SetIsInternalSynth]
//   - [IAVSpeechSynthesizer.SetActiveOptions]
//   - [IAVSpeechSynthesizer.SetAudioQueueFlags]
//   - [IAVSpeechSynthesizer.SetSetActiveOptions]
//   - [IAVSpeechSynthesizer.Paused]
//   - [IAVSpeechSynthesizer.Speaking]
type IAVSpeechSynthesizer interface {
	objectivec.IObject

	// Topic: Methods

	_configureCoreSynthesizer(synthesizer objectivec.IObject)
	AudioDeviceId() uint32
	SetAudioDeviceId(value uint32)
	AudioQueueFlags() uint32
	CoreSynthesizer() unsafe.Pointer
	ExistingCoreSynthesizer() unsafe.Pointer
	IsInternalSynth() bool
	SetIsInternalSynth(value bool)
	SetActiveOptions() uint64
	SetAudioQueueFlags(flags uint32)
	SetSetActiveOptions(options uint64)
	Paused() bool
	Speaking() bool
}

// Init initializes the instance.
func (a AVSpeechSynthesizer) Init() AVSpeechSynthesizer {
	rv := objc.SendIfResponds[AVSpeechSynthesizer](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVSpeechSynthesizer) Autorelease() AVSpeechSynthesizer {
	rv := objc.SendIfResponds[AVSpeechSynthesizer](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSpeechSynthesizer creates a new AVSpeechSynthesizer instance.
func NewAVSpeechSynthesizer() AVSpeechSynthesizer {
	class := getAVSpeechSynthesizerClass()
	rv := objc.SendIfResponds[AVSpeechSynthesizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (a AVSpeechSynthesizer) _configureCoreSynthesizer(synthesizer objectivec.IObject) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("_configureCoreSynthesizer:"), synthesizer)
}

// ConfigureCoreSynthesizer is an exported wrapper for the private method _configureCoreSynthesizer.
func (a AVSpeechSynthesizer) ConfigureCoreSynthesizer(synthesizer objectivec.IObject) error {
	if !objc.RespondsToSelector(a.ID, objc.Sel("_configureCoreSynthesizer:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_configureCoreSynthesizer:"}
		return err
	}
	a._configureCoreSynthesizer(synthesizer)
	return nil
}

// CanConfigureCoreSynthesizer reports whether the receiver responds to the private selector _configureCoreSynthesizer:.
func (a AVSpeechSynthesizer) CanConfigureCoreSynthesizer() bool {
	return objc.RespondsToSelector(a.ID, objc.Sel("_configureCoreSynthesizer:"))
}
func (a AVSpeechSynthesizer) AudioQueueFlags() uint32 {
	rv := objc.SendIfResponds[uint32](a.ID, objc.Sel("audioQueueFlags"))
	return rv
}
func (a AVSpeechSynthesizer) SetActiveOptions() uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("setActiveOptions"))
	return rv
}
func (a AVSpeechSynthesizer) SetAudioQueueFlags(flags uint32) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("setAudioQueueFlags:"), flags)
}
func (a AVSpeechSynthesizer) SetSetActiveOptions(options uint64) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("setSetActiveOptions:"), options)
}

func (_AVSpeechSynthesizerClass AVSpeechSynthesizerClass) _supportsSpeakingWithPersonalVoices() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_AVSpeechSynthesizerClass.class), objc.Sel("_supportsSpeakingWithPersonalVoices"))
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
func (_AVSpeechSynthesizerClass AVSpeechSynthesizerClass) IsSoftAppUsageProtectionDisabled() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_AVSpeechSynthesizerClass.class), objc.Sel("isSoftAppUsageProtectionDisabled"))
	return rv
}

func (a AVSpeechSynthesizer) AudioDeviceId() uint32 {
	rv := objc.SendIfResponds[uint32](a.ID, objc.Sel("audioDeviceId"))
	return rv
}
func (a AVSpeechSynthesizer) SetAudioDeviceId(value uint32) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setAudioDeviceId:"), value)
}
func (a AVSpeechSynthesizer) CoreSynthesizer() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](a.ID, objc.Sel("coreSynthesizer"))
	return rv
}
func (a AVSpeechSynthesizer) ExistingCoreSynthesizer() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](a.ID, objc.Sel("existingCoreSynthesizer"))
	return rv
}
func (a AVSpeechSynthesizer) IsInternalSynth() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("isInternalSynth"))
	return rv
}
func (a AVSpeechSynthesizer) SetIsInternalSynth(value bool) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setIsInternalSynth:"), value)
}
func (a AVSpeechSynthesizer) Paused() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("paused"))
	return rv
}
func (a AVSpeechSynthesizer) Speaking() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("speaking"))
	return rv
}
