// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A protocol that defines the methods to respond to audio playback events and decoding errors.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerDelegate
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

// Tells the delegate when the audio finishes playing.
//
// player: The audio player that finishes playing.
//
// flag: A Boolean value that indicates whether the audio finishes playing
// successfully.
//
// # Discussion
//
// The system doesn’t call this method on an audio interruption.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerDelegate/audioPlayerDidFinishPlaying(_:successfully:)
func (o AVAudioPlayerDelegateObject) AudioPlayerDidFinishPlayingSuccessfully(player IAVAudioPlayer, flag bool) {
	objc.Send[struct{}](o.ID, objc.Sel("audioPlayerDidFinishPlaying:successfully:"), player, flag)
}

// Tells the delegate when an audio player encounters a decoding error during
// playback.
//
// player: The audio player that encounters the decoding error.
//
// error: The decoding error.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayerDelegate/audioPlayerDecodeErrorDidOccur(_:error:)
func (o AVAudioPlayerDelegateObject) AudioPlayerDecodeErrorDidOccurError(player IAVAudioPlayer, error_ foundation.INSError) {
	objc.Send[struct{}](o.ID, objc.Sel("audioPlayerDecodeErrorDidOccur:error:"), player, error_)
}

// AVAudioPlayerDelegateConfig holds optional typed callbacks for [AVAudioPlayerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfaudio/avaudioplayerdelegate
type AVAudioPlayerDelegateConfig struct {

	// Responding to Playback Completion
	// AudioPlayerDidFinishPlayingSuccessfully — Tells the delegate when the audio finishes playing.
	AudioPlayerDidFinishPlayingSuccessfully func(player AVAudioPlayer, flag bool)

	// Responding to Audio Decoding Errors
	// AudioPlayerDecodeErrorDidOccurError — Tells the delegate when an audio player encounters a decoding error during playback.
	AudioPlayerDecodeErrorDidOccurError func(player AVAudioPlayer, error_ foundation.NSError)
}

// NewAVAudioPlayerDelegate creates an Objective-C object implementing the [AVAudioPlayerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AVAudioPlayerDelegateObject] satisfies the [AVAudioPlayerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfaudio/avaudioplayerdelegate
func NewAVAudioPlayerDelegate(config AVAudioPlayerDelegateConfig) AVAudioPlayerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAVAudioPlayerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.AudioPlayerDidFinishPlayingSuccessfully != nil {
		fn := config.AudioPlayerDidFinishPlayingSuccessfully
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("audioPlayerDidFinishPlaying:successfully:"),
			Fn: func(self objc.ID, _cmd objc.SEL, playerID objc.ID, flag bool) {
				player := AVAudioPlayerFromID(playerID)
				fn(player, flag)
			},
		})
	}

	if config.AudioPlayerDecodeErrorDidOccurError != nil {
		fn := config.AudioPlayerDecodeErrorDidOccurError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("audioPlayerDecodeErrorDidOccur:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, playerID objc.ID, error_ID objc.ID) {
				player := AVAudioPlayerFromID(playerID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(player, error_)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("AVAudioPlayerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAVAudioPlayerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AVAudioPlayerDelegateObjectFromID(instance)
}
