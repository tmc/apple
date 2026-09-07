// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AudioPlayerDelegateHelper] class.
var (
	_AudioPlayerDelegateHelperClass     AudioPlayerDelegateHelperClass
	_AudioPlayerDelegateHelperClassOnce sync.Once
)

func getAudioPlayerDelegateHelperClass() AudioPlayerDelegateHelperClass {
	_AudioPlayerDelegateHelperClassOnce.Do(func() {
		_AudioPlayerDelegateHelperClass = AudioPlayerDelegateHelperClass{class: objc.GetClass("_TtC12TextToSpeechP33_FD049640B58505424D931116721C328425AudioPlayerDelegateHelper")}
	})
	return _AudioPlayerDelegateHelperClass
}

// GetAudioPlayerDelegateHelperClass returns the class object for _TtC12TextToSpeechP33_FD049640B58505424D931116721C328425AudioPlayerDelegateHelper.
func GetAudioPlayerDelegateHelperClass() AudioPlayerDelegateHelperClass {
	return getAudioPlayerDelegateHelperClass()
}

type AudioPlayerDelegateHelperClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AudioPlayerDelegateHelperClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AudioPlayerDelegateHelperClass) Alloc() AudioPlayerDelegateHelper {
	rv := objc.SendIfResponds[AudioPlayerDelegateHelper](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AudioPlayerDelegateHelper.AudioPlayerDidFinishPlayingSuccessfully]
type AudioPlayerDelegateHelper struct {
	objectivec.Object
}

// AudioPlayerDelegateHelperFromID constructs a [AudioPlayerDelegateHelper] from an objc.ID.
func AudioPlayerDelegateHelperFromID(id objc.ID) AudioPlayerDelegateHelper {
	return AudioPlayerDelegateHelper{objectivec.Object{ID: id}}
}

// Ensure AudioPlayerDelegateHelper implements IAudioPlayerDelegateHelper.
var _ IAudioPlayerDelegateHelper = AudioPlayerDelegateHelper{}

// An interface definition for the [AudioPlayerDelegateHelper] class.
//
// # Methods
//
//   - [IAudioPlayerDelegateHelper.AudioPlayerDidFinishPlayingSuccessfully]
type IAudioPlayerDelegateHelper interface {
	objectivec.IObject

	// Topic: Methods

	AudioPlayerDidFinishPlayingSuccessfully(playing objectivec.IObject, successfully bool)
}

// Init initializes the instance.
func (a AudioPlayerDelegateHelper) Init() AudioPlayerDelegateHelper {
	rv := objc.SendIfResponds[AudioPlayerDelegateHelper](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AudioPlayerDelegateHelper) Autorelease() AudioPlayerDelegateHelper {
	rv := objc.SendIfResponds[AudioPlayerDelegateHelper](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioPlayerDelegateHelper creates a new AudioPlayerDelegateHelper instance.
func NewAudioPlayerDelegateHelper() AudioPlayerDelegateHelper {
	class := getAudioPlayerDelegateHelperClass()
	rv := objc.SendIfResponds[AudioPlayerDelegateHelper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (a AudioPlayerDelegateHelper) AudioPlayerDidFinishPlayingSuccessfully(playing objectivec.IObject, successfully bool) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("audioPlayerDidFinishPlaying:successfully:"), playing, successfully)
}
