// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// AVAudioPlayerDelegate protocol.
type AVAudioPlayerDelegate interface {
	objectivec.IObject

	// AudioPlayerDecodeErrorDidOccurError protocol.
	AudioPlayerDecodeErrorDidOccurError(occur objectivec.IObject, error_ objectivec.IObject)

	// AudioPlayerDidFinishPlayingSuccessfully protocol.
	AudioPlayerDidFinishPlayingSuccessfully(playing objectivec.IObject, successfully bool)
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

func (o AVAudioPlayerDelegateObject) AudioPlayerDecodeErrorDidOccurError(occur objectivec.IObject, error_ objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("audioPlayerDecodeErrorDidOccur:error:"), occur, error_)
}
func (o AVAudioPlayerDelegateObject) AudioPlayerDidFinishPlayingSuccessfully(playing objectivec.IObject, successfully bool) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("audioPlayerDidFinishPlaying:successfully:"), playing, successfully)
}
