// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/audiotoolbox"
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
	rv := objc.SendIfResponds[TTSAUSSEWrapper](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TTSAUSSEWrapper.AudioUnit]
//   - [TTSAUSSEWrapper.SetAudioUnit]
//   - [TTSAUSSEWrapper.CancelSpeechRequest]
//   - [TTSAUSSEWrapper.SynthesizeSpeechRequest]
//   - [TTSAUSSEWrapper.InitWithAudioUnit]
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
type ITTSAUSSEWrapper interface {
	objectivec.IObject

	// Topic: Methods

	AudioUnit() audiotoolbox.AUAudioUnit
	SetAudioUnit(value audiotoolbox.AUAudioUnit)
	CancelSpeechRequest()
	SynthesizeSpeechRequest(request objectivec.IObject)
	InitWithAudioUnit(unit objectivec.IObject) TTSAUSSEWrapper
}

// Init initializes the instance.
func (t TTSAUSSEWrapper) Init() TTSAUSSEWrapper {
	rv := objc.SendIfResponds[TTSAUSSEWrapper](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSAUSSEWrapper) Autorelease() TTSAUSSEWrapper {
	rv := objc.SendIfResponds[TTSAUSSEWrapper](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSAUSSEWrapper creates a new TTSAUSSEWrapper instance.
func NewTTSAUSSEWrapper() TTSAUSSEWrapper {
	class := getTTSAUSSEWrapperClass()
	rv := objc.SendIfResponds[TTSAUSSEWrapper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewTTSAUSSEWrapperWithAudioUnit(unit objectivec.IObject) TTSAUSSEWrapper {
	instance := getTTSAUSSEWrapperClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithAudioUnit:"), unit)
	return TTSAUSSEWrapperFromID(rv)
}

func (t TTSAUSSEWrapper) CancelSpeechRequest() {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("cancelSpeechRequest"))
}
func (t TTSAUSSEWrapper) SynthesizeSpeechRequest(request objectivec.IObject) {
	objc.SendIfResponds[objc.ID](t.ID, objc.Sel("synthesizeSpeechRequest:"), request)
}
func (t TTSAUSSEWrapper) InitWithAudioUnit(unit objectivec.IObject) TTSAUSSEWrapper {
	rv := objc.SendIfResponds[TTSAUSSEWrapper](t.ID, objc.Sel("initWithAudioUnit:"), unit)
	return rv
}

func (_TTSAUSSEWrapperClass TTSAUSSEWrapperClass) MakeAU(au objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_TTSAUSSEWrapperClass.class), objc.Sel("makeAU:"), au)
	return objectivec.Object{ID: rv}
}

func (t TTSAUSSEWrapper) AudioUnit() audiotoolbox.AUAudioUnit {
	rv := objc.SendIfResponds[objc.ID](t.ID, objc.Sel("audioUnit"))
	return audiotoolbox.AUAudioUnitFromID(objc.ID(rv))
}
func (t TTSAUSSEWrapper) SetAudioUnit(value audiotoolbox.AUAudioUnit) {
	objc.SendIfResponds[struct{}](t.ID, objc.Sel("setAudioUnit:"), value)
}
