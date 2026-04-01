// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
//go:build ios
// +build ios

package avfoundation

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// Whether the player’s audio output is suppressed due to being on a
// non-mixable audio route.
//
// # Discussion
//
// If YES, the player’s audio output is suppressed. The player is muted
// while on a non-mixable audio route and cannot play audio. The player’s
// mute property does not reflect the true mute status. If NO, the player’s
// audio output is not suppressed. The player may be muted or unmuted while on
// a non-mixable audio route and can play audio. The player’s mute property
// reflects the true mute status. In a non-mixable audio route, only one
// player can play audio. To play audio in non-mixable states, the player must
// be specified as the priority participant in
// AVRoutingPlaybackArbiter.preferredParticipantForNonMixableAudioRoutes. If
// this player becomes the preferred player, it will gain audio priority and
// suppress the audio of all other players. If another participant becomes the
// preferred participant, this player will lose audio priority and have their
// audio suppressed. This property is key-value observed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/audioOutputSuppressedDueToNonMixableAudioRoute
func (p AVPlayer) AudioOutputSuppressedDueToNonMixableAudioRoute() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("audioOutputSuppressedDueToNonMixableAudioRoute"))
	return rv
}

// A Boolean value that indicates whether the player should automatically
// switch to external playback mode while the external screen mode is active.
//
// # Discussion
//
// The player automatically switches back to the external screen mode once
// video playback concludes. A brief transition may be visible on the external
// display when automatically switching between the two modes. The default
// value of this property is false. The value of this property has no effect
// if [AllowsExternalPlayback] is false.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/usesExternalPlaybackWhileExternalScreenIsActive
func (p AVPlayer) UsesExternalPlaybackWhileExternalScreenIsActive() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("usesExternalPlaybackWhileExternalScreenIsActive"))
	return rv
}
func (p AVPlayer) SetUsesExternalPlaybackWhileExternalScreenIsActive(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setUsesExternalPlaybackWhileExternalScreenIsActive:"), value)
}

// The video gravity of the player for external playback mode only.
//
// # Discussion
//
// Valid values are [resize], [resizeAspectFill], or [resizeAspect].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/externalPlaybackVideoGravity
//
// [resizeAspectFill]: https://developer.apple.com/documentation/AVFoundation/AVLayerVideoGravity/resizeAspectFill
// [resizeAspect]: https://developer.apple.com/documentation/AVFoundation/AVLayerVideoGravity/resizeAspect
// [resize]: https://developer.apple.com/documentation/AVFoundation/AVLayerVideoGravity/resize
func (p AVPlayer) ExternalPlaybackVideoGravity() AVLayerVideoGravity {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("externalPlaybackVideoGravity"))
	return AVLayerVideoGravity(foundation.NSStringFromID(rv).String())
}
func (p AVPlayer) SetExternalPlaybackVideoGravity(value AVLayerVideoGravity) {
	objc.Send[struct{}](p.ID, objc.Sel("setExternalPlaybackVideoGravity:"), value)
}

// A Boolean value that indicates whether video playback prevents the system
// from automatically backgrounding the app.
//
// # Discussion
//
// The default value is true, which indicates the system doesn’t
// automatically background an app while it’s actively playing video. A user
// may still choose to background an app.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/preventsAutomaticBackgroundingDuringVideoPlayback
func (p AVPlayer) PreventsAutomaticBackgroundingDuringVideoPlayback() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("preventsAutomaticBackgroundingDuringVideoPlayback"))
	return rv
}
func (p AVPlayer) SetPreventsAutomaticBackgroundingDuringVideoPlayback(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setPreventsAutomaticBackgroundingDuringVideoPlayback:"), value)
}

// A Boolean value that indicates whether the player allows AirPlay video
// playback.
//
// # Discussion
//
// The default value is true.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/allowsAirPlayVideo
func (p AVPlayer) AllowsAirPlayVideo() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("allowsAirPlayVideo"))
	return rv
}
func (p AVPlayer) SetAllowsAirPlayVideo(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAllowsAirPlayVideo:"), value)
}

// A Boolean value that indicates whether the player is playing video through
// AirPlay.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/isAirPlayVideoActive
func (p AVPlayer) AirPlayVideoActive() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAirPlayVideoActive"))
	return rv
}

// A Boolean value that indicates whether the player automatically switches to
// AirPlay Video while AirPlay Screen is active.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/usesAirPlayVideoWhileAirPlayScreenIsActive
func (p AVPlayer) UsesAirPlayVideoWhileAirPlayScreenIsActive() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("usesAirPlayVideoWhileAirPlayScreenIsActive"))
	return rv
}
func (p AVPlayer) SetUsesAirPlayVideoWhileAirPlayScreenIsActive(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setUsesAirPlayVideoWhileAirPlayScreenIsActive:"), value)
}

// The AVPlayer’s intended spatial audio experience.
//
// # Discussion
//
// The default value of CAAutomaticSpatialAudio means the player uses its
// AVAudioSession’s intended spatial experience. If the anchoring strategy
// is impossible (e.g. it uses a destroyed UIScene’s identifier), the player
// follows a “front” anchoring strategy instead.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/intendedSpatialAudioExperience-3uy8g
func (p AVPlayer) IntendedSpatialAudioExperience() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p.ID, objc.Sel("intendedSpatialAudioExperience"))
	return rv
}
func (p AVPlayer) SetIntendedSpatialAudioExperience(value unsafe.Pointer) {
	objc.Send[struct{}](p.ID, objc.Sel("setIntendedSpatialAudioExperience:"), value)
}

// The HDR modes that are available for playback.
//
// # Discussion
//
// This property indicates all of the HDR modes that the device can play. A
// value of `0` indicates the device doesn’t support HDR. Each value
// indicates that an appropriate HDR display is available for the specified
// HDR mode. Additionally, the device must be capable of playing the specified
// HDR type.
//
// This property doesn’t indicate whether a video contains HDR content, HDR
// video is currently playing, or if video is playing on an HDR display.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/availableHDRModes
func (_AVPlayerClass AVPlayerClass) AvailableHDRModes() AVPlayerHDRMode {
	rv := objc.Send[AVPlayerHDRMode](objc.ID(_AVPlayerClass.class), objc.Sel("availableHDRModes"))
	return AVPlayerHDRMode(rv)
}
