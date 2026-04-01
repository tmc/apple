// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.
//go:build ios
// +build ios

package avfaudio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An array of channel descriptions for the audio player.
//
// # Discussion
//
// The default value for this property is `nil`. When the value is non-`nil`,
// this array must have the same number of channels the [NumberOfChannels]
// property returns. You can use this property to assign output to play to
// different channels.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/channelAssignments
func (a AVAudioPlayer) ChannelAssignments() []objc.ID {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("channelAssignments"))
	return rv
}
func (a AVAudioPlayer) SetChannelAssignments(value []objc.ID) {
	objc.Send[struct{}](a.ID, objc.Sel("setChannelAssignments:"), objectivec.IObjectSliceToNSArray(value))
}

// The intended spatial experience for this player.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/intendedSpatialExperience-6py9z
func (a AVAudioPlayer) IntendedSpatialExperience() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a.ID, objc.Sel("intendedSpatialExperience"))
	return rv
}
func (a AVAudioPlayer) SetIntendedSpatialExperience(value unsafe.Pointer) {
	objc.Send[struct{}](a.ID, objc.Sel("setIntendedSpatialExperience:"), value)
}
