// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVCDuckFadeDuration] class.
var (
	_AVVCDuckFadeDurationClass     AVVCDuckFadeDurationClass
	_AVVCDuckFadeDurationClassOnce sync.Once
)

func getAVVCDuckFadeDurationClass() AVVCDuckFadeDurationClass {
	_AVVCDuckFadeDurationClassOnce.Do(func() {
		_AVVCDuckFadeDurationClass = AVVCDuckFadeDurationClass{class: objc.GetClass("AVVCDuckFadeDuration")}
	})
	return _AVVCDuckFadeDurationClass
}

// GetAVVCDuckFadeDurationClass returns the class object for AVVCDuckFadeDuration.
func GetAVVCDuckFadeDurationClass() AVVCDuckFadeDurationClass {
	return getAVVCDuckFadeDurationClass()
}

type AVVCDuckFadeDurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVCDuckFadeDurationClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVCDuckFadeDurationClass) Alloc() AVVCDuckFadeDuration {
	rv := objc.Send[AVVCDuckFadeDuration](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVVCDuckFadeDuration.FadeIn]
//   - [AVVCDuckFadeDuration.SetFadeIn]
//   - [AVVCDuckFadeDuration.FadeOut]
//   - [AVVCDuckFadeDuration.SetFadeOut]
//   - [AVVCDuckFadeDuration.InitWithFadeInFadeOut]
// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckFadeDuration
type AVVCDuckFadeDuration struct {
	objectivec.Object
}

// AVVCDuckFadeDurationFromID constructs a [AVVCDuckFadeDuration] from an objc.ID.
func AVVCDuckFadeDurationFromID(id objc.ID) AVVCDuckFadeDuration {
	return AVVCDuckFadeDuration{objectivec.Object{ID: id}}
}
// Ensure AVVCDuckFadeDuration implements IAVVCDuckFadeDuration.
var _ IAVVCDuckFadeDuration = AVVCDuckFadeDuration{}

// An interface definition for the [AVVCDuckFadeDuration] class.
//
// # Methods
//
//   - [IAVVCDuckFadeDuration.FadeIn]
//   - [IAVVCDuckFadeDuration.SetFadeIn]
//   - [IAVVCDuckFadeDuration.FadeOut]
//   - [IAVVCDuckFadeDuration.SetFadeOut]
//   - [IAVVCDuckFadeDuration.InitWithFadeInFadeOut]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckFadeDuration
type IAVVCDuckFadeDuration interface {
	objectivec.IObject

	// Topic: Methods

	FadeIn() foundation.NSNumber
	SetFadeIn(value foundation.NSNumber)
	FadeOut() foundation.NSNumber
	SetFadeOut(value foundation.NSNumber)
	InitWithFadeInFadeOut(in objectivec.IObject, out objectivec.IObject) AVVCDuckFadeDuration
}

// Init initializes the instance.
func (v AVVCDuckFadeDuration) Init() AVVCDuckFadeDuration {
	rv := objc.Send[AVVCDuckFadeDuration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVCDuckFadeDuration) Autorelease() AVVCDuckFadeDuration {
	rv := objc.Send[AVVCDuckFadeDuration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCDuckFadeDuration creates a new AVVCDuckFadeDuration instance.
func NewAVVCDuckFadeDuration() AVVCDuckFadeDuration {
	class := getAVVCDuckFadeDurationClass()
	rv := objc.Send[AVVCDuckFadeDuration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckFadeDuration/initWithFadeIn:fadeOut:
func NewVCDuckFadeDurationWithFadeInFadeOut(in objectivec.IObject, out objectivec.IObject) AVVCDuckFadeDuration {
	instance := getAVVCDuckFadeDurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFadeIn:fadeOut:"), in, out)
	return AVVCDuckFadeDurationFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckFadeDuration/initWithFadeIn:fadeOut:
func (v AVVCDuckFadeDuration) InitWithFadeInFadeOut(in objectivec.IObject, out objectivec.IObject) AVVCDuckFadeDuration {
	rv := objc.Send[AVVCDuckFadeDuration](v.ID, objc.Sel("initWithFadeIn:fadeOut:"), in, out)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckFadeDuration/fadeIn
func (v AVVCDuckFadeDuration) FadeIn() foundation.NSNumber {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("fadeIn"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (v AVVCDuckFadeDuration) SetFadeIn(value foundation.NSNumber) {
	objc.Send[struct{}](v.ID, objc.Sel("setFadeIn:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckFadeDuration/fadeOut
func (v AVVCDuckFadeDuration) FadeOut() foundation.NSNumber {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("fadeOut"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (v AVVCDuckFadeDuration) SetFadeOut(value foundation.NSNumber) {
	objc.Send[struct{}](v.ID, objc.Sel("setFadeOut:"), value)
}

