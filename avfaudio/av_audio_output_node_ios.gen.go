// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.
//go:build ios
// +build ios

package avfaudio

import (
	"github.com/tmc/apple/objc"
)

// The intended spatial experience for this output node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioOutputNode/intendedSpatialExperience-3uznq
func (a AVAudioOutputNode) IntendedSpatialExperience() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a.ID, objc.Sel("intendedSpatialExperience"))
	return rv
}
func (a AVAudioOutputNode) SetIntendedSpatialExperience(value unsafe.Pointer) {
	objc.Send[struct{}](a.ID, objc.Sel("setIntendedSpatialExperience:"), value)
}
