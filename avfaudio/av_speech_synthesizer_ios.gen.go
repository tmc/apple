// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.
//go:build ios
// +build ios

package avfaudio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A Boolean value that specifies whether the app manages the audio session.
//
// # Discussion
//
// If you set this value to false, the system creates a separate audio session
// to automatically manage speech, interruptions, and mixing and ducking the
// speech with other audio sources.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/usesApplicationAudioSession
func (s AVSpeechSynthesizer) UsesApplicationAudioSession() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("usesApplicationAudioSession"))
	return rv
}
func (s AVSpeechSynthesizer) SetUsesApplicationAudioSession(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setUsesApplicationAudioSession:"), value)
}

// A Boolean value that specifies whether to send synthesized speech to an
// active call.
//
// # Discussion
//
// This property has no effect when there isn’t an active call.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/mixToTelephonyUplink
func (s AVSpeechSynthesizer) MixToTelephonyUplink() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("mixToTelephonyUplink"))
	return rv
}
func (s AVSpeechSynthesizer) SetMixToTelephonyUplink(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setMixToTelephonyUplink:"), value)
}

// An array of audio session channels to route generated speech.
//
// # Discussion
//
// The system replicates speech audio to each audio session channel.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSpeechSynthesizer/outputChannels
func (s AVSpeechSynthesizer) OutputChannels() []objc.ID {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("outputChannels"))
	return rv
}
func (s AVSpeechSynthesizer) SetOutputChannels(value []objc.ID) {
	objc.Send[struct{}](s.ID, objc.Sel("setOutputChannels:"), objectivec.IObjectSliceToNSArray(value))
}
