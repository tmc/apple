// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.
//go:build ios
// +build ios

package avfaudio

import (
	"github.com/tmc/apple/objc"
)

// Requests the app’s permission to add audio to calls.
//
// # Discussion
//
// The system immediately returns a response if a person has already granted
// or denied the app permission, or if the service is in a disabled state.
// Otherwise, it presents a dialog to request permission and returns a result
// when a person dismisses the UI.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/requestMicrophoneInjectionPermission(completionHandler:)
func (_AVAudioApplicationClass AVAudioApplicationClass) RequestMicrophoneInjectionPermissionWithCompletionHandler(response AVAudioApplicationMicrophoneInjectionPermissionHandler) {
	_block0, _cleanup0 := NewAVAudioApplicationMicrophoneInjectionPermissionBlock(response)
	defer _cleanup0()
	objc.Send[objc.ID](objc.ID(_AVAudioApplicationClass.class), objc.Sel("requestMicrophoneInjectionPermissionWithCompletionHandler:"), _block0)
}

// A value that indicates an app’s permission to add audio to calls.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioApplication/microphoneInjectionPermission-swift.property
func (a AVAudioApplication) MicrophoneInjectionPermission() AVAudioApplicationMicrophoneInjectionPermission {
	rv := objc.Send[AVAudioApplicationMicrophoneInjectionPermission](a.ID, objc.Sel("microphoneInjectionPermission"))
	return AVAudioApplicationMicrophoneInjectionPermission(rv)
}
