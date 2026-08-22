// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A protocol with methods for managing multi-utterance speech recognition requests.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskDelegate
type SFSpeechRecognitionTaskDelegate interface {
	objectivec.IObject
}

// SFSpeechRecognitionTaskDelegateObject wraps an existing Objective-C object that conforms to the SFSpeechRecognitionTaskDelegate protocol.
type SFSpeechRecognitionTaskDelegateObject struct {
	objectivec.Object
}

func (o SFSpeechRecognitionTaskDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// SFSpeechRecognitionTaskDelegateObjectFromID constructs a [SFSpeechRecognitionTaskDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func SFSpeechRecognitionTaskDelegateObjectFromID(id objc.ID) SFSpeechRecognitionTaskDelegateObject {
	return SFSpeechRecognitionTaskDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate when the task first detects speech in the source audio.
//
// task: The speech recognition task (an [SFSpeechRecognitionTask] object) that
// represents the request.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskDelegate/speechRecognitionDidDetectSpeech(_:)
func (o SFSpeechRecognitionTaskDelegateObject) SpeechRecognitionDidDetectSpeech(task ISFSpeechRecognitionTask) {
	objc.Send[struct{}](o.ID, objc.Sel("speechRecognitionDidDetectSpeech:"), task)
}

// Tells the delegate when the task is no longer accepting new audio input,
// even if final processing is in progress.
//
// task: The speech recognition task (an [SFSpeechRecognitionTask] object) that
// represents the request.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskDelegate/speechRecognitionTaskFinishedReadingAudio(_:)
func (o SFSpeechRecognitionTaskDelegateObject) SpeechRecognitionTaskFinishedReadingAudio(task ISFSpeechRecognitionTask) {
	objc.Send[struct{}](o.ID, objc.Sel("speechRecognitionTaskFinishedReadingAudio:"), task)
}

// Tells the delegate that a hypothesized transcription is available.
//
// task: The speech recognition task (an [SFSpeechRecognitionTask] object) that
// represents the request.
//
// transcription: The hypothesized transcription in an [SFTranscription] object.
//
// # Discussion
//
// This method is called for all recognitions, including partial recognitions.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskDelegate/speechRecognitionTask(_:didHypothesizeTranscription:)
func (o SFSpeechRecognitionTaskDelegateObject) SpeechRecognitionTaskDidHypothesizeTranscription(task ISFSpeechRecognitionTask, transcription ISFTranscription) {
	objc.Send[struct{}](o.ID, objc.Sel("speechRecognitionTask:didHypothesizeTranscription:"), task, transcription)
}

// Tells the delegate when the final utterance is recognized.
//
// task: The speech recognition task (an [SFSpeechRecognitionTask] object) that
// represents the request.
//
// recognitionResult: A recognized utterance that contains one or more transcription hypotheses
// in an [SFSpeechRecognitionResult] object.
//
// # Discussion
//
// When this method is called, the delegate should expect no further
// information about the utterance to be reported.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskDelegate/speechRecognitionTask(_:didFinishRecognition:)
func (o SFSpeechRecognitionTaskDelegateObject) SpeechRecognitionTaskDidFinishRecognition(task ISFSpeechRecognitionTask, recognitionResult ISFSpeechRecognitionResult) {
	objc.Send[struct{}](o.ID, objc.Sel("speechRecognitionTask:didFinishRecognition:"), task, recognitionResult)
}

// Tells the delegate when the recognition of all requested utterances is
// finished.
//
// task: The speech recognition task (an [SFSpeechRecognitionTask] object) that
// represents the request.
//
// successfully: A Boolean value that indicates whether the task was successful. When this
// parameter is `false`, use the [SFSpeechRecognitionTask.Error] property of
// the task to get information about why the task was unsuccessful.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskDelegate/speechRecognitionTask(_:didFinishSuccessfully:)
func (o SFSpeechRecognitionTaskDelegateObject) SpeechRecognitionTaskDidFinishSuccessfully(task ISFSpeechRecognitionTask, successfully bool) {
	objc.Send[struct{}](o.ID, objc.Sel("speechRecognitionTask:didFinishSuccessfully:"), task, successfully)
}

// Tells the delegate how much audio has been processed by the task.
//
// task: The speech recognition task (an [SFSpeechRecognitionTask] object) that
// represents the request.
//
// duration: The seconds of audio input that the recognizer has processed.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskDelegate/speechRecognitionTask(_:didProcessAudioDuration:)
func (o SFSpeechRecognitionTaskDelegateObject) SpeechRecognitionTaskDidProcessAudioDuration(task ISFSpeechRecognitionTask, duration foundation.NSTimeInterval) {
	objc.Send[struct{}](o.ID, objc.Sel("speechRecognitionTask:didProcessAudioDuration:"), task, duration)
}

// Tells the delegate that the task has been canceled.
//
// task: The speech recognition task (an [SFSpeechRecognitionTask] object) that
// represents the request.
//
// # Discussion
//
// A speech recognition task can be canceled by the user, by your app, or by
// the system.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskDelegate/speechRecognitionTaskWasCancelled(_:)
func (o SFSpeechRecognitionTaskDelegateObject) SpeechRecognitionTaskWasCancelled(task ISFSpeechRecognitionTask) {
	objc.Send[struct{}](o.ID, objc.Sel("speechRecognitionTaskWasCancelled:"), task)
}

// SFSpeechRecognitionTaskDelegateConfig holds optional typed callbacks for [SFSpeechRecognitionTaskDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/speech/sfspeechrecognitiontaskdelegate
type SFSpeechRecognitionTaskDelegateConfig struct {

	// Tracking task progress
	// SpeechRecognitionDidDetectSpeech — Tells the delegate when the task first detects speech in the source audio.
	SpeechRecognitionDidDetectSpeech func(task SFSpeechRecognitionTask)
	// SpeechRecognitionTaskFinishedReadingAudio — Tells the delegate when the task is no longer accepting new audio input, even if final processing is in progress.
	SpeechRecognitionTaskFinishedReadingAudio func(task SFSpeechRecognitionTask)

	// Getting transcriptions
	// SpeechRecognitionTaskDidHypothesizeTranscription — Tells the delegate that a hypothesized transcription is available.
	SpeechRecognitionTaskDidHypothesizeTranscription func(task SFSpeechRecognitionTask, transcription SFTranscription)

	// Finishing a speech recognition task
	// SpeechRecognitionTaskDidFinishRecognition — Tells the delegate when the final utterance is recognized.
	SpeechRecognitionTaskDidFinishRecognition func(task SFSpeechRecognitionTask, recognitionResult SFSpeechRecognitionResult)
	// SpeechRecognitionTaskDidFinishSuccessfully — Tells the delegate when the recognition of all requested utterances is finished.
	SpeechRecognitionTaskDidFinishSuccessfully func(task SFSpeechRecognitionTask, successfully bool)
	// SpeechRecognitionTaskDidProcessAudioDuration — Tells the delegate how much audio has been processed by the task.
	SpeechRecognitionTaskDidProcessAudioDuration func(task SFSpeechRecognitionTask, duration foundation.NSTimeInterval)
	// SpeechRecognitionTaskWasCancelled — Tells the delegate that the task has been canceled.
	SpeechRecognitionTaskWasCancelled func(task SFSpeechRecognitionTask)
}

// NewSFSpeechRecognitionTaskDelegate creates an Objective-C object implementing the [SFSpeechRecognitionTaskDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [SFSpeechRecognitionTaskDelegateObject] satisfies the [SFSpeechRecognitionTaskDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/speech/sfspeechrecognitiontaskdelegate
func NewSFSpeechRecognitionTaskDelegate(config SFSpeechRecognitionTaskDelegateConfig) SFSpeechRecognitionTaskDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoSFSpeechRecognitionTaskDelegate_%d", n)

	var methods []objc.MethodDef

	if config.SpeechRecognitionDidDetectSpeech != nil {
		fn := config.SpeechRecognitionDidDetectSpeech
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("speechRecognitionDidDetectSpeech:"),
			Fn: func(self objc.ID, _cmd objc.SEL, taskID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("SFSpeechRecognitionTaskDelegate", "speechRecognitionDidDetectSpeech:")
					}
				}()
				task := SFSpeechRecognitionTaskFromID(taskID)
				fn(task)
				_delegateDone = true
			},
		})
	}

	if config.SpeechRecognitionTaskFinishedReadingAudio != nil {
		fn := config.SpeechRecognitionTaskFinishedReadingAudio
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("speechRecognitionTaskFinishedReadingAudio:"),
			Fn: func(self objc.ID, _cmd objc.SEL, taskID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("SFSpeechRecognitionTaskDelegate", "speechRecognitionTaskFinishedReadingAudio:")
					}
				}()
				task := SFSpeechRecognitionTaskFromID(taskID)
				fn(task)
				_delegateDone = true
			},
		})
	}

	if config.SpeechRecognitionTaskDidHypothesizeTranscription != nil {
		fn := config.SpeechRecognitionTaskDidHypothesizeTranscription
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("speechRecognitionTask:didHypothesizeTranscription:"),
			Fn: func(self objc.ID, _cmd objc.SEL, taskID objc.ID, transcriptionID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("SFSpeechRecognitionTaskDelegate", "speechRecognitionTask:didHypothesizeTranscription:")
					}
				}()
				task := SFSpeechRecognitionTaskFromID(taskID)
				transcription := SFTranscriptionFromID(transcriptionID)
				fn(task, transcription)
				_delegateDone = true
			},
		})
	}

	if config.SpeechRecognitionTaskDidFinishRecognition != nil {
		fn := config.SpeechRecognitionTaskDidFinishRecognition
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("speechRecognitionTask:didFinishRecognition:"),
			Fn: func(self objc.ID, _cmd objc.SEL, taskID objc.ID, recognitionResultID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("SFSpeechRecognitionTaskDelegate", "speechRecognitionTask:didFinishRecognition:")
					}
				}()
				task := SFSpeechRecognitionTaskFromID(taskID)
				recognitionResult := SFSpeechRecognitionResultFromID(recognitionResultID)
				fn(task, recognitionResult)
				_delegateDone = true
			},
		})
	}

	if config.SpeechRecognitionTaskDidFinishSuccessfully != nil {
		fn := config.SpeechRecognitionTaskDidFinishSuccessfully
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("speechRecognitionTask:didFinishSuccessfully:"),
			Fn: func(self objc.ID, _cmd objc.SEL, taskID objc.ID, successfully bool) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("SFSpeechRecognitionTaskDelegate", "speechRecognitionTask:didFinishSuccessfully:")
					}
				}()
				task := SFSpeechRecognitionTaskFromID(taskID)
				fn(task, successfully)
				_delegateDone = true
			},
		})
	}

	if config.SpeechRecognitionTaskDidProcessAudioDuration != nil {
		fn := config.SpeechRecognitionTaskDidProcessAudioDuration
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("speechRecognitionTask:didProcessAudioDuration:"),
			Fn: func(self objc.ID, _cmd objc.SEL, taskID objc.ID, duration foundation.NSTimeInterval) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("SFSpeechRecognitionTaskDelegate", "speechRecognitionTask:didProcessAudioDuration:")
					}
				}()
				task := SFSpeechRecognitionTaskFromID(taskID)
				fn(task, duration)
				_delegateDone = true
			},
		})
	}

	if config.SpeechRecognitionTaskWasCancelled != nil {
		fn := config.SpeechRecognitionTaskWasCancelled
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("speechRecognitionTaskWasCancelled:"),
			Fn: func(self objc.ID, _cmd objc.SEL, taskID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("SFSpeechRecognitionTaskDelegate", "speechRecognitionTaskWasCancelled:")
					}
				}()
				task := SFSpeechRecognitionTaskFromID(taskID)
				fn(task)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("SFSpeechRecognitionTaskDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewSFSpeechRecognitionTaskDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return SFSpeechRecognitionTaskDelegateObjectFromID(instance)
}
