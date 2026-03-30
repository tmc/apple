// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.
//go:build ios
// +build ios

package avfaudio

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The intended spatial experience for this output node.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioOutputNode/intendedSpatialExperience-3uznq
func (a AVAudioOutputNode) IntendedSpatialExperience() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("intendedSpatialExperience"))
	return objectivec.Object{ID: rv}
}
func (a AVAudioOutputNode) SetIntendedSpatialExperience(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setIntendedSpatialExperience:"), value)
}
