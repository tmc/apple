// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AudioPlayerImpl] class.
var (
	_AudioPlayerImplClass     AudioPlayerImplClass
	_AudioPlayerImplClassOnce sync.Once
)

func getAudioPlayerImplClass() AudioPlayerImplClass {
	_AudioPlayerImplClassOnce.Do(func() {
		_AudioPlayerImplClass = AudioPlayerImplClass{class: objc.GetClass("AudioPlayerImpl")}
	})
	return _AudioPlayerImplClass
}

// GetAudioPlayerImplClass returns the class object for AudioPlayerImpl.
func GetAudioPlayerImplClass() AudioPlayerImplClass {
	return getAudioPlayerImplClass()
}

type AudioPlayerImplClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AudioPlayerImplClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AudioPlayerImplClass) Alloc() AudioPlayerImpl {
	rv := objc.Send[AudioPlayerImpl](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AudioPlayerImpl
type AudioPlayerImpl struct {
	objectivec.Object
}

// AudioPlayerImplFromID constructs a [AudioPlayerImpl] from an objc.ID.
func AudioPlayerImplFromID(id objc.ID) AudioPlayerImpl {
	return AudioPlayerImpl{objectivec.Object{ID: id}}
}
// Ensure AudioPlayerImpl implements IAudioPlayerImpl.
var _ IAudioPlayerImpl = AudioPlayerImpl{}

// An interface definition for the [AudioPlayerImpl] class.
//
// See: https://developer.apple.com/documentation/AVFAudio/AudioPlayerImpl
type IAudioPlayerImpl interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a AudioPlayerImpl) Init() AudioPlayerImpl {
	rv := objc.Send[AudioPlayerImpl](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AudioPlayerImpl) Autorelease() AudioPlayerImpl {
	rv := objc.Send[AudioPlayerImpl](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioPlayerImpl creates a new AudioPlayerImpl instance.
func NewAudioPlayerImpl() AudioPlayerImpl {
	class := getAudioPlayerImplClass()
	rv := objc.Send[AudioPlayerImpl](objc.ID(class.class), objc.Sel("new"))
	return rv
}

