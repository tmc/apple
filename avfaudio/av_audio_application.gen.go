// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioApplication] class.
var (
	_AVAudioApplicationClass     AVAudioApplicationClass
	_AVAudioApplicationClassOnce sync.Once
)

func getAVAudioApplicationClass() AVAudioApplicationClass {
	_AVAudioApplicationClassOnce.Do(func() {
		_AVAudioApplicationClass = AVAudioApplicationClass{class: objc.GetClass("AVAudioApplication")}
	})
	return _AVAudioApplicationClass
}

// GetAVAudioApplicationClass returns the class object for AVAudioApplication.
func GetAVAudioApplicationClass() AVAudioApplicationClass {
	return getAVAudioApplicationClass()
}

type AVAudioApplicationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioApplicationClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioApplicationClass) Alloc() AVAudioApplication {
	rv := objc.Send[AVAudioApplication](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that manages one or more audio sessions that belong to an app.
//
// # Overview
// 
// Access the shared audio application instance to control app-level audio
// operations, such as requesting microphone permission and controlling audio
// input muting.
//
// # Requesting audio recording permission
//
//   - [AVAudioApplication.RecordPermission]: The app’s permission to record audio.
//
// # Managing audio input mute state
//
//   - [AVAudioApplication.InputMuted]: A Boolean value that indicates whether the app’s audio input is in a muted state.
//   - [AVAudioApplication.SetInputMutedError]: Sets a Boolean value that indicates whether the app’s audio input is in a muted state.
//   - [AVAudioApplication.SetInputMuteStateChangeHandlerError]: Sets a callback to handle changes to application-level audio muting states.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication
type AVAudioApplication struct {
	objectivec.Object
}

// AVAudioApplicationFromID constructs a [AVAudioApplication] from an objc.ID.
//
// An object that manages one or more audio sessions that belong to an app.
func AVAudioApplicationFromID(id objc.ID) AVAudioApplication {
	return AVAudioApplication{objectivec.Object{ID: id}}
}
// NOTE: AVAudioApplication adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioApplication] class.
//
// # Requesting audio recording permission
//
//   - [IAVAudioApplication.RecordPermission]: The app’s permission to record audio.
//
// # Managing audio input mute state
//
//   - [IAVAudioApplication.InputMuted]: A Boolean value that indicates whether the app’s audio input is in a muted state.
//   - [IAVAudioApplication.SetInputMutedError]: Sets a Boolean value that indicates whether the app’s audio input is in a muted state.
//   - [IAVAudioApplication.SetInputMuteStateChangeHandlerError]: Sets a callback to handle changes to application-level audio muting states.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication
type IAVAudioApplication interface {
	objectivec.IObject

	// Topic: Requesting audio recording permission

	// The app’s permission to record audio.
	RecordPermission() AVAudioApplicationRecordPermission

	// Topic: Managing audio input mute state

	// A Boolean value that indicates whether the app’s audio input is in a muted state.
	InputMuted() bool
	// Sets a Boolean value that indicates whether the app’s audio input is in a muted state.
	SetInputMutedError(muted bool) (bool, error)
	// Sets a callback to handle changes to application-level audio muting states.
	SetInputMuteStateChangeHandlerError(inputMuteHandler func(bool) bool) (bool, error)
}

// Init initializes the instance.
func (a AVAudioApplication) Init() AVAudioApplication {
	rv := objc.Send[AVAudioApplication](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioApplication) Autorelease() AVAudioApplication {
	rv := objc.Send[AVAudioApplication](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioApplication creates a new AVAudioApplication instance.
func NewAVAudioApplication() AVAudioApplication {
	class := getAVAudioApplicationClass()
	rv := objc.Send[AVAudioApplication](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Sets a Boolean value that indicates whether the app’s audio input is in a
// muted state.
//
// muted: A Boolean value that indicates the new mute state.
//
// # Discussion
// 
// In platforms that use [AVAudioSession], setting the value to [true] mutes
// all sources of audio input in the app. In macOS, the system instead invokes
// the callback that you register by calling
// [SetInputMuteStateChangeHandlerError] to handle input muting.
//
// [AVAudioSession]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/setInputMuted(_:)
func (a AVAudioApplication) SetInputMutedError(muted bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("setInputMuted:error:"), muted, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setInputMuted:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Sets a callback to handle changes to application-level audio muting states.
//
// inputMuteHandler: A callback that the system invokes when the input mute state changes. If
// the callback receives a [true] value, mute all input audio samples until
// the next time the system calls the handler. Return a value of [true] if you
// muted input successfully, or in exceptional cases, return [false] to
// indicate the mute action fails.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method to set a closure to handle your macOS app’s input muting
// logic. The system calls thie closure when the input mute state changes,
// either due to setting the [InputMuted] state, or due to a Bluetooth audio
// accessory gesture (certain AirPods / Beats headphones) changing the mute
// state.
// 
// Since the input mute handling logic should happen a single place,
// subsequent calls to this method overwrite any previously registered block
// with the one you provide. You can specify a `nil` to cancel the callback.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/setInputMuteStateChangeHandler(_:)
func (a AVAudioApplication) SetInputMuteStateChangeHandlerError(inputMuteHandler func(bool) bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("setInputMuteStateChangeHandler:error:"), inputMuteHandler, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setInputMuteStateChangeHandler:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Determines whether the app has permission to record audio.
//
// response: A Boolean value that indicates whether the user grants the app permission
// to record audio.
//
// # Discussion
// 
// Recording audio requires explicit permission from the user. The first time
// your app attempts to record audio input, the system automatically prompts
// the user for permission. You can also explicitly ask for permission by
// calling this method. This method returns immediately, but the system waits
// for user input if the user hasn’t previously granted or denied recording
// permission.
// 
// Unless a user grants your app permission to record audio, it captures only
// silence (zeroed out audio samples).
// 
// After a user responds to a recording permission prompt from your app, the
// system remembers their choice and won’t prompt them again. If a user
// denies the app recording permission, they can grant it access in the
// Privacy & Security section of the Settings app.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/requestRecordPermission(completionHandler:)
func (_AVAudioApplicationClass AVAudioApplicationClass) RequestRecordPermissionWithCompletionHandler(response BoolHandler) {
_block0, _ := NewBoolBlock(response)
	objc.Send[objc.ID](objc.ID(_AVAudioApplicationClass.class), objc.Sel("requestRecordPermissionWithCompletionHandler:"), _block0)
}

// The app’s permission to record audio.
//
// # Discussion
// 
// See [RequestRecordPermissionWithCompletionHandler] for more information.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/recordPermission-swift.property
func (a AVAudioApplication) RecordPermission() AVAudioApplicationRecordPermission {
	rv := objc.Send[AVAudioApplicationRecordPermission](a.ID, objc.Sel("recordPermission"))
	return AVAudioApplicationRecordPermission(rv)
}
// A Boolean value that indicates whether the app’s audio input is in a
// muted state.
//
// # Discussion
// 
// Set a new value for this property by calling the [SetInputMutedError]
// method.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/isInputMuted
func (a AVAudioApplication) InputMuted() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isInputMuted"))
	return rv
}

// Accesses the shared audio application instance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/shared
func (_AVAudioApplicationClass AVAudioApplicationClass) SharedInstance() AVAudioApplication {
	rv := objc.Send[objc.ID](objc.ID(_AVAudioApplicationClass.class), objc.Sel("sharedInstance"))
	return AVAudioApplicationFromID(objc.ID(rv))
}

// RequestRecordPermission is a synchronous wrapper around [AVAudioApplication.RequestRecordPermissionWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (ac AVAudioApplicationClass) RequestRecordPermission(ctx context.Context) (bool, error) {
	done := make(chan bool, 1)
	ac.RequestRecordPermissionWithCompletionHandler(func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

