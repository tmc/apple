// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SFSpeechRecognitionTask] class.
var (
	_SFSpeechRecognitionTaskClass     SFSpeechRecognitionTaskClass
	_SFSpeechRecognitionTaskClassOnce sync.Once
)

func getSFSpeechRecognitionTaskClass() SFSpeechRecognitionTaskClass {
	_SFSpeechRecognitionTaskClassOnce.Do(func() {
		_SFSpeechRecognitionTaskClass = SFSpeechRecognitionTaskClass{class: objc.GetClass("SFSpeechRecognitionTask")}
	})
	return _SFSpeechRecognitionTaskClass
}

// GetSFSpeechRecognitionTaskClass returns the class object for SFSpeechRecognitionTask.
func GetSFSpeechRecognitionTaskClass() SFSpeechRecognitionTaskClass {
	return getSFSpeechRecognitionTaskClass()
}

type SFSpeechRecognitionTaskClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SFSpeechRecognitionTaskClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SFSpeechRecognitionTaskClass) Alloc() SFSpeechRecognitionTask {
	rv := objc.Send[SFSpeechRecognitionTask](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A task object for monitoring the speech recognition progress.
//
// # Overview
//
// Use an [SFSpeechRecognitionTask] object to determine the state of a speech
// recognition task, to cancel an ongoing task, or to signal the end of the
// task.
//
// You don’t create speech recognition task objects directly. Instead, you
// receive one of these objects after calling
// [SFSpeechRecognizer.RecognitionTaskWithRequestResultHandler] or
// [SFSpeechRecognizer.RecognitionTaskWithRequestDelegate] on your
// [SFSpeechRecognizer] object.
//
// # Canceling a speech recognition task
//
//   - [SFSpeechRecognitionTask.Cancel]: Cancels the current speech recognition task.
//   - [SFSpeechRecognitionTask.IsCancelled]: A Boolean value that indicates whether the speech recognition task was canceled.
//
// # Finishing a speech recognition task
//
//   - [SFSpeechRecognitionTask.Finish]: Stops accepting new audio and finishes processing on the audio input that has already been accepted.
//   - [SFSpeechRecognitionTask.IsFinishing]: A Boolean value that indicates whether audio input has stopped.
//
// # Monitoring recognition progress
//
//   - [SFSpeechRecognitionTask.State]: The current state of the speech recognition task.
//   - [SFSpeechRecognitionTask.Error]: An error object that specifies the error that occurred during a speech recognition task.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTask
type SFSpeechRecognitionTask struct {
	objectivec.Object
}

// SFSpeechRecognitionTaskFromID constructs a [SFSpeechRecognitionTask] from an objc.ID.
//
// A task object for monitoring the speech recognition progress.
func SFSpeechRecognitionTaskFromID(id objc.ID) SFSpeechRecognitionTask {
	return SFSpeechRecognitionTask{objectivec.Object{ID: id}}
}

// NOTE: SFSpeechRecognitionTask adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SFSpeechRecognitionTask] class.
//
// # Canceling a speech recognition task
//
//   - [ISFSpeechRecognitionTask.Cancel]: Cancels the current speech recognition task.
//   - [ISFSpeechRecognitionTask.IsCancelled]: A Boolean value that indicates whether the speech recognition task was canceled.
//
// # Finishing a speech recognition task
//
//   - [ISFSpeechRecognitionTask.Finish]: Stops accepting new audio and finishes processing on the audio input that has already been accepted.
//   - [ISFSpeechRecognitionTask.IsFinishing]: A Boolean value that indicates whether audio input has stopped.
//
// # Monitoring recognition progress
//
//   - [ISFSpeechRecognitionTask.State]: The current state of the speech recognition task.
//   - [ISFSpeechRecognitionTask.Error]: An error object that specifies the error that occurred during a speech recognition task.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTask
type ISFSpeechRecognitionTask interface {
	objectivec.IObject

	// Topic: Canceling a speech recognition task

	// Cancels the current speech recognition task.
	Cancel()
	// A Boolean value that indicates whether the speech recognition task was canceled.
	IsCancelled() bool

	// Topic: Finishing a speech recognition task

	// Stops accepting new audio and finishes processing on the audio input that has already been accepted.
	Finish()
	// A Boolean value that indicates whether audio input has stopped.
	IsFinishing() bool

	// Topic: Monitoring recognition progress

	// The current state of the speech recognition task.
	State() SFSpeechRecognitionTaskState
	// An error object that specifies the error that occurred during a speech recognition task.
	Error() foundation.NSError
}

// Init initializes the instance.
func (s SFSpeechRecognitionTask) Init() SFSpeechRecognitionTask {
	rv := objc.Send[SFSpeechRecognitionTask](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SFSpeechRecognitionTask) Autorelease() SFSpeechRecognitionTask {
	rv := objc.Send[SFSpeechRecognitionTask](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSpeechRecognitionTask creates a new SFSpeechRecognitionTask instance.
func NewSFSpeechRecognitionTask() SFSpeechRecognitionTask {
	class := getSFSpeechRecognitionTaskClass()
	rv := objc.Send[SFSpeechRecognitionTask](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Cancels the current speech recognition task.
//
// # Discussion
//
// You can cancel recognition tasks for both prerecorded and live audio input.
// For example, you might cancel a task in response to a user action or
// because the recording was interrupted.
//
// When canceling a task, be sure to release any resources associated with the
// task, such as the audio input resources you are using to capture audio
// samples.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTask/cancel()
func (s SFSpeechRecognitionTask) Cancel() {
	objc.Send[objc.ID](s.ID, objc.Sel("cancel"))
}

// Stops accepting new audio and finishes processing on the audio input that
// has already been accepted.
//
// # Discussion
//
// For audio buffer–based recognition, recognition does not finish until
// this method is called, so be sure to call it when the audio source is
// exhausted.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTask/finish()
func (s SFSpeechRecognitionTask) Finish() {
	objc.Send[objc.ID](s.ID, objc.Sel("finish"))
}

// A Boolean value that indicates whether the speech recognition task was
// canceled.
//
// # Discussion
//
// By default, the value of this property is `false`.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTask/isCancelled
func (s SFSpeechRecognitionTask) IsCancelled() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isCancelled"))
	return rv
}

// A Boolean value that indicates whether audio input has stopped.
//
// # Discussion
//
// By default, the value of this property is `false`.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTask/isFinishing
func (s SFSpeechRecognitionTask) IsFinishing() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isFinishing"))
	return rv
}

// The current state of the speech recognition task.
//
// # Discussion
//
// Check the value of this property to get the state of the in-progress speech
// recognition session. For valid values, see [SFSpeechRecognitionTaskState].
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTask/state
//
// [SFSpeechRecognitionTaskState]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTaskState
func (s SFSpeechRecognitionTask) State() SFSpeechRecognitionTaskState {
	rv := objc.Send[SFSpeechRecognitionTaskState](s.ID, objc.Sel("state"))
	return SFSpeechRecognitionTaskState(rv)
}

// An error object that specifies the error that occurred during a speech
// recognition task.
//
// # Discussion
//
// The system may return one of the errors listed in the table below.
//
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTask/error
func (s SFSpeechRecognitionTask) Error() foundation.NSError {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
