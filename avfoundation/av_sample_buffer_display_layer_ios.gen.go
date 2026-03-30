// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
//go:build ios
// +build ios

package avfoundation

import (
	"github.com/tmc/apple/objc"
)

// A Boolean value that indicates whether video playback prevents the system
// from automatically backgrounding an app.
//
// # Discussion
//
// The default value is true, which indicates the system doesn’t
// automatically background an app while playing video. The value of this
// property doesn’t prevent the user from backgrounding an app.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/preventsAutomaticBackgroundingDuringVideoPlayback
func (s AVSampleBufferDisplayLayer) PreventsAutomaticBackgroundingDuringVideoPlayback() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("preventsAutomaticBackgroundingDuringVideoPlayback"))
	return rv
}
func (s AVSampleBufferDisplayLayer) SetPreventsAutomaticBackgroundingDuringVideoPlayback(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setPreventsAutomaticBackgroundingDuringVideoPlayback:"), value)
}
