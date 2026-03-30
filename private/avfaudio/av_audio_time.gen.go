// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioTime] class.
var (
	_AVAudioTimeClass     AVAudioTimeClass
	_AVAudioTimeClassOnce sync.Once
)

func getAVAudioTimeClass() AVAudioTimeClass {
	_AVAudioTimeClassOnce.Do(func() {
		_AVAudioTimeClass = AVAudioTimeClass{class: objc.GetClass("AVAudioTime")}
	})
	return _AVAudioTimeClass
}

// GetAVAudioTimeClass returns the class object for AVAudioTime.
func GetAVAudioTimeClass() AVAudioTimeClass {
	return getAVAudioTimeClass()
}

type AVAudioTimeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioTimeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioTimeClass) Alloc() AVAudioTime {
	rv := objc.Send[AVAudioTime](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioTime.HostTimeValid]
//   - [AVAudioTime.SampleTimeValid]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime
type AVAudioTime struct {
	objectivec.Object
}

// AVAudioTimeFromID constructs a [AVAudioTime] from an objc.ID.
func AVAudioTimeFromID(id objc.ID) AVAudioTime {
	return AVAudioTime{objectivec.Object{ID: id}}
}

// Ensure AVAudioTime implements IAVAudioTime.
var _ IAVAudioTime = AVAudioTime{}

// An interface definition for the [AVAudioTime] class.
//
// # Methods
//
//   - [IAVAudioTime.HostTimeValid]
//   - [IAVAudioTime.SampleTimeValid]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime
type IAVAudioTime interface {
	objectivec.IObject

	// Topic: Methods

	HostTimeValid() bool
	SampleTimeValid() bool
}

// Init initializes the instance.
func (a AVAudioTime) Init() AVAudioTime {
	rv := objc.Send[AVAudioTime](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioTime) Autorelease() AVAudioTime {
	rv := objc.Send[AVAudioTime](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioTime creates a new AVAudioTime instance.
func NewAVAudioTime() AVAudioTime {
	class := getAVAudioTimeClass()
	rv := objc.Send[AVAudioTime](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/hostTimeValid
func (a AVAudioTime) HostTimeValid() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("hostTimeValid"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/sampleTimeValid
func (a AVAudioTime) SampleTimeValid() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("sampleTimeValid"))
	return rv
}
