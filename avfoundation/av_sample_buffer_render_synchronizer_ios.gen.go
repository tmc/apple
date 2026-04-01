// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
//go:build ios
// +build ios

package avfoundation

import (
	"github.com/tmc/apple/objc"
)

// The intended spatial audio experience applied to all
// AVSampleBufferAudioRenderers within this synchronizer.
//
// # Discussion
//
// The default value of CAAutomaticSpatialAudio means the renderers use their
// AVAudioSession’s intended spatial experience. If the anchoring strategy
// is impossible (e.g. it uses a destroyed UIScene’s identifier), the
// renderers follow a “front” anchoring strategy instead.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/intendedSpatialAudioExperience-2wthu
func (s AVSampleBufferRenderSynchronizer) IntendedSpatialAudioExperience() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s.ID, objc.Sel("intendedSpatialAudioExperience"))
	return rv
}
func (s AVSampleBufferRenderSynchronizer) SetIntendedSpatialAudioExperience(value unsafe.Pointer) {
	objc.Send[struct{}](s.ID, objc.Sel("setIntendedSpatialAudioExperience:"), value)
}
