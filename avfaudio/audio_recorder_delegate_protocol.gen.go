// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A protocol that defines the methods to respond to audio recording events and encoding errors.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorderDelegate
type AVAudioRecorderDelegate interface {
	objectivec.IObject
}

// AVAudioRecorderDelegateObject wraps an existing Objective-C object that conforms to the AVAudioRecorderDelegate protocol.
type AVAudioRecorderDelegateObject struct {
	objectivec.Object
}
func (o AVAudioRecorderDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVAudioRecorderDelegateObjectFromID constructs a [AVAudioRecorderDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVAudioRecorderDelegateObjectFromID(id objc.ID) AVAudioRecorderDelegateObject {
	return AVAudioRecorderDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate when recording stops or finishes due to reaching its
// time limit.
//
// recorder: The audio recorder that finished recording.
//
// flag: A Boolean value that indicates whether the recording stopped successfully.
//
// # Discussion
// 
// The system doesn’t call this method if the recorder stops due to an
// interruption.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorderDelegate/audioRecorderDidFinishRecording(_:successfully:)
func (o AVAudioRecorderDelegateObject) AudioRecorderDidFinishRecordingSuccessfully(recorder IAVAudioRecorder, flag bool) {
	objc.Send[struct{}](o.ID, objc.Sel("audioRecorderDidFinishRecording:successfully:"), recorder, flag)
	}
// Tells the delegate that the audio recorder encountered an encoding error
// during recording.
//
// recorder: The audio recorder that encountered the encoding error.
//
// error: An object that provides the details of the encoding error.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorderDelegate/audioRecorderEncodeErrorDidOccur(_:error:)
func (o AVAudioRecorderDelegateObject) AudioRecorderEncodeErrorDidOccurError(recorder IAVAudioRecorder, error_ foundation.INSError) {
	objc.Send[struct{}](o.ID, objc.Sel("audioRecorderEncodeErrorDidOccur:error:"), recorder, error_)
	}

// AVAudioRecorderDelegateConfig holds optional typed callbacks for [AVAudioRecorderDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfaudio/avaudiorecorderdelegate
type AVAudioRecorderDelegateConfig struct {

	// Responding to Recording Completion
	// AudioRecorderDidFinishRecordingSuccessfully — Tells the delegate when recording stops or finishes due to reaching its time limit.
	AudioRecorderDidFinishRecordingSuccessfully func(recorder AVAudioRecorder, flag bool)

	// Responding to Audio Encoding Errors
	// AudioRecorderEncodeErrorDidOccurError — Tells the delegate that the audio recorder encountered an encoding error during recording.
	AudioRecorderEncodeErrorDidOccurError func(recorder AVAudioRecorder, error_ foundation.NSError)
}

// NewAVAudioRecorderDelegate creates an Objective-C object implementing the [AVAudioRecorderDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AVAudioRecorderDelegateObject] satisfies the [AVAudioRecorderDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfaudio/avaudiorecorderdelegate
func NewAVAudioRecorderDelegate(config AVAudioRecorderDelegateConfig) AVAudioRecorderDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAVAudioRecorderDelegate_%d", n)

	var methods []objc.MethodDef

	if config.AudioRecorderDidFinishRecordingSuccessfully != nil {
		fn := config.AudioRecorderDidFinishRecordingSuccessfully
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("audioRecorderDidFinishRecording:successfully:"),
			Fn: func(self objc.ID, _cmd objc.SEL, recorderID objc.ID, flag bool) {
				recorder := AVAudioRecorderFromID(recorderID)
				fn(recorder, flag)
			},
		})
	}

	if config.AudioRecorderEncodeErrorDidOccurError != nil {
		fn := config.AudioRecorderEncodeErrorDidOccurError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("audioRecorderEncodeErrorDidOccur:error:"),
			Fn: func(self objc.ID, _cmd objc.SEL, recorderID objc.ID, error_ID objc.ID) {
				recorder := AVAudioRecorderFromID(recorderID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(recorder, error_)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("AVAudioRecorderDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAVAudioRecorderDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AVAudioRecorderDelegateObjectFromID(instance)
}

