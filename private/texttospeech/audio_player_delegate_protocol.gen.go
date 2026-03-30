// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// AVAudioPlayerDelegate protocol.
//
// See: https://developer.apple.com/documentation/TextToSpeech/AVAudioPlayerDelegate
type AVAudioPlayerDelegate interface {
	objectivec.IObject
}

// AVAudioPlayerDelegateObject wraps an existing Objective-C object that conforms to the AVAudioPlayerDelegate protocol.
type AVAudioPlayerDelegateObject struct {
	objectivec.Object
}

func (o AVAudioPlayerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVAudioPlayerDelegateObjectFromID constructs a [AVAudioPlayerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVAudioPlayerDelegateObjectFromID(id objc.ID) AVAudioPlayerDelegateObject {
	return AVAudioPlayerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/TextToSpeech/AVAudioPlayerDelegate/audioPlayerDecodeErrorDidOccur:error:
func (o AVAudioPlayerDelegateObject) AudioPlayerDecodeErrorDidOccurError(occur objectivec.IObject, error_ objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("audioPlayerDecodeErrorDidOccur:error:"), occur, error_)
}

// See: https://developer.apple.com/documentation/TextToSpeech/AVAudioPlayerDelegate/audioPlayerDidFinishPlaying:successfully:
func (o AVAudioPlayerDelegateObject) AudioPlayerDidFinishPlayingSuccessfully(playing objectivec.IObject, successfully bool) {
	objc.Send[struct{}](o.ID, objc.Sel("audioPlayerDidFinishPlaying:successfully:"), playing, successfully)
}
