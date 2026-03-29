// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.
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
// [Settings] property for the [AVNumberOfChannelsKey] value. Use this
// property to help record specific audio channels.
//
// [AVNumberOfChannelsKey]: https://developer.apple.com/documentation/AVFAudio/AVNumberOfChannelsKey
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/channelAssignments
func (a AVAudioRecorder) ChannelAssignments() []objc.ID {
rv := objc.Send[[]objc.ID](a.ID, objc.Sel("channelAssignments"))
return rv
}
func (a AVAudioRecorder) SetChannelAssignments(value []objc.ID) {
objc.Send[struct{}](a.ID, objc.Sel("setChannelAssignments:"), objectivec.IObjectSliceToNSArray(value))
}

