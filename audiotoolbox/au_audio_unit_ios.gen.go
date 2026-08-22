// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.
//go:build ios
// +build ios

package audiotoolbox

import (
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The AUAudioUnit’s intended spatial experience.
//
// # Discussion
//
// Only useful for output AUAudioUnits - setting on a non-output AU is a
// no-op. The default value of CAAutomaticSpatialAudio means the output
// AUAudioUnit uses its AVAudioSession’s spatial experience. See
// CASpatialAudioExperience for more details.
//
// See: https://developer.apple.com/documentation/AudioToolbox/AUAudioUnit/intendedSpatialExperience-1dvhd
func (a AUAudioUnit) IntendedSpatialExperience() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a.ID, objc.Sel("intendedSpatialExperience"))
	return rv
}
func (a AUAudioUnit) SetIntendedSpatialExperience(value unsafe.Pointer) {
	objc.Send[struct{}](a.ID, objc.Sel("setIntendedSpatialExperience:"), value)
}
