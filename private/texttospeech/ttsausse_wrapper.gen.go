// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSAUSSEWrapper] class.
var (
	_TTSAUSSEWrapperClass     TTSAUSSEWrapperClass
	_TTSAUSSEWrapperClassOnce sync.Once
)

func getTTSAUSSEWrapperClass() TTSAUSSEWrapperClass {
	_TTSAUSSEWrapperClassOnce.Do(func() {
		_TTSAUSSEWrapperClass = TTSAUSSEWrapperClass{class: objc.GetClass("TTSAUSSEWrapper")}
	})
	return _TTSAUSSEWrapperClass
}

// GetTTSAUSSEWrapperClass returns the class object for TTSAUSSEWrapper.
func GetTTSAUSSEWrapperClass() TTSAUSSEWrapperClass {
	return getTTSAUSSEWrapperClass()
}

type TTSAUSSEWrapperClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSAUSSEWrapperClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSAUSSEWrapperClass) Alloc() TTSAUSSEWrapper {
	rv := objc.Send[TTSAUSSEWrapper](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [TTSAUSSEWrapper.AudioUnit]
//   - [TTSAUSSEWrapper.SetAudioUnit]
//   - [TTSAUSSEWrapper.CancelSpeechRequest]
//   - [TTSAUSSEWrapper.SynthesizeSpeechRequest]
//   - [TTSAUSSEWrapper.InitWithAudioUnit]
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUSSEWrapper
type TTSAUSSEWrapper struct {
	objectivec.Object
}

// TTSAUSSEWrapperFromID constructs a [TTSAUSSEWrapper] from an objc.ID.
func TTSAUSSEWrapperFromID(id objc.ID) TTSAUSSEWrapper {
	return TTSAUSSEWrapper{objectivec.Object{ID: id}}
}
// Ensure TTSAUSSEWrapper implements ITTSAUSSEWrapper.
var _ ITTSAUSSEWrapper = TTSAUSSEWrapper{}

// An interface definition for the [TTSAUSSEWrapper] class.
//
// # Methods
//
//   - [ITTSAUSSEWrapper.AudioUnit]
//   - [ITTSAUSSEWrapper.SetAudioUnit]
//   - [ITTSAUSSEWrapper.CancelSpeechRequest]
//   - [ITTSAUSSEWrapper.SynthesizeSpeechRequest]
//   - [ITTSAUSSEWrapper.InitWithAudioUnit]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUSSEWrapper
type ITTSAUSSEWrapper interface {
	objectivec.IObject

	// Topic: Methods

	AudioUnit() unsafe.Pointer
	SetAudioUnit(value unsafe.Pointer)
	CancelSpeechRequest()
	SynthesizeSpeechRequest(request objectivec.IObject)
	InitWithAudioUnit(unit objectivec.IObject) TTSAUSSEWrapper
}

// Init initializes the instance.
func (t TTSAUSSEWrapper) Init() TTSAUSSEWrapper {
	rv := objc.Send[TTSAUSSEWrapper](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSAUSSEWrapper) Autorelease() TTSAUSSEWrapper {
	rv := objc.Send[TTSAUSSEWrapper](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSAUSSEWrapper creates a new TTSAUSSEWrapper instance.
func NewTTSAUSSEWrapper() TTSAUSSEWrapper {
	class := getTTSAUSSEWrapperClass()
	rv := objc.Send[TTSAUSSEWrapper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUSSEWrapper/initWithAudioUnit:
func NewTTSAUSSEWrapperWithAudioUnit(unit objectivec.IObject) TTSAUSSEWrapper {
	instance := getTTSAUSSEWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAudioUnit:"), unit)
	return TTSAUSSEWrapperFromID(rv)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUSSEWrapper/cancelSpeechRequest
func (t TTSAUSSEWrapper) CancelSpeechRequest() {
	objc.Send[objc.ID](t.ID, objc.Sel("cancelSpeechRequest"))
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUSSEWrapper/synthesizeSpeechRequest:
func (t TTSAUSSEWrapper) SynthesizeSpeechRequest(request objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("synthesizeSpeechRequest:"), request)
}
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUSSEWrapper/initWithAudioUnit:
func (t TTSAUSSEWrapper) InitWithAudioUnit(unit objectivec.IObject) TTSAUSSEWrapper {
	rv := objc.Send[TTSAUSSEWrapper](t.ID, objc.Sel("initWithAudioUnit:"), unit)
	return rv
}

//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUSSEWrapper/makeAU:
func (_TTSAUSSEWrapperClass TTSAUSSEWrapperClass) MakeAU(au objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_TTSAUSSEWrapperClass.class), objc.Sel("makeAU:"), au)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAUSSEWrapper/audioUnit
func (t TTSAUSSEWrapper) AudioUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t.ID, objc.Sel("audioUnit"))
	return rv
}
func (t TTSAUSSEWrapper) SetAudioUnit(value unsafe.Pointer) {
	objc.Send[struct{}](t.ID, objc.Sel("setAudioUnit:"), value)
}

