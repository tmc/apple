// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.
//go:build ios
// +build ios

package avfaudio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An array of channel descriptions associated with the audio recorder.
//
// # Discussion
//
// The default value of this property is `nil`. When the value is non-`nil`,
// this value must have the same number of channels as defined in the
// [AVAudioRecorder.Settings] property for the [AVNumberOfChannelsKey] value.
// Use this property to help record specific audio channels.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/channelAssignments
//
// [AVNumberOfChannelsKey]: https://developer.apple.com/documentation/AVFAudio/AVNumberOfChannelsKey
func (a AVAudioRecorder) ChannelAssignments() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("channelAssignments"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (a AVAudioRecorder) SetChannelAssignments(value []objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setChannelAssignments:"), objectivec.IObjectSliceToNSArray(value))
}
