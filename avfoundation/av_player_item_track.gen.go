// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerItemTrack] class.
var (
	_AVPlayerItemTrackClass     AVPlayerItemTrackClass
	_AVPlayerItemTrackClassOnce sync.Once
)

func getAVPlayerItemTrackClass() AVPlayerItemTrackClass {
	_AVPlayerItemTrackClassOnce.Do(func() {
		_AVPlayerItemTrackClass = AVPlayerItemTrackClass{class: objc.GetClass("AVPlayerItemTrack")}
	})
	return _AVPlayerItemTrackClass
}

// GetAVPlayerItemTrackClass returns the class object for AVPlayerItemTrack.
func GetAVPlayerItemTrackClass() AVPlayerItemTrackClass {
	return getAVPlayerItemTrackClass()
}

type AVPlayerItemTrackClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlayerItemTrackClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerItemTrackClass) Alloc() AVPlayerItemTrack {
	rv := objc.Send[AVPlayerItemTrack](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the presentation state of an asset track during
// playback.
//
// # Setting the enabled state
//
//   - [AVPlayerItemTrack.Enabled]: A Boolean value that indicates whether the player item presents the track’s media during playback.
//   - [AVPlayerItemTrack.SetEnabled]
//
// # Configuring video properties
//
//   - [AVPlayerItemTrack.CurrentVideoFrameRate]: The current frame rate of the video track as it plays.
//   - [AVPlayerItemTrack.VideoFieldMode]: A mode that specifies the handling of video frames that contain multiple fields.
//   - [AVPlayerItemTrack.SetVideoFieldMode]
//   - [AVPlayerItemTrack.AVPlayerItemTrackVideoFieldModeDeinterlaceFields]: A video field mode that requests deinterlacing of video fields.
//
// # Accessing the asset track
//
//   - [AVPlayerItemTrack.AssetTrack]: An asset track that provides the media for the player item track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrack
type AVPlayerItemTrack struct {
	objectivec.Object
}

// AVPlayerItemTrackFromID constructs a [AVPlayerItemTrack] from an objc.ID.
//
// An object that represents the presentation state of an asset track during
// playback.
func AVPlayerItemTrackFromID(id objc.ID) AVPlayerItemTrack {
	return AVPlayerItemTrack{objectivec.Object{ID: id}}
}

// NOTE: AVPlayerItemTrack adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerItemTrack] class.
//
// # Setting the enabled state
//
//   - [IAVPlayerItemTrack.Enabled]: A Boolean value that indicates whether the player item presents the track’s media during playback.
//   - [IAVPlayerItemTrack.SetEnabled]
//
// # Configuring video properties
//
//   - [IAVPlayerItemTrack.CurrentVideoFrameRate]: The current frame rate of the video track as it plays.
//   - [IAVPlayerItemTrack.VideoFieldMode]: A mode that specifies the handling of video frames that contain multiple fields.
//   - [IAVPlayerItemTrack.SetVideoFieldMode]
//   - [IAVPlayerItemTrack.AVPlayerItemTrackVideoFieldModeDeinterlaceFields]: A video field mode that requests deinterlacing of video fields.
//
// # Accessing the asset track
//
//   - [IAVPlayerItemTrack.AssetTrack]: An asset track that provides the media for the player item track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrack
type IAVPlayerItemTrack interface {
	objectivec.IObject

	// Topic: Setting the enabled state

	// A Boolean value that indicates whether the player item presents the track’s media during playback.
	Enabled() bool
	SetEnabled(value bool)

	// Topic: Configuring video properties

	// The current frame rate of the video track as it plays.
	CurrentVideoFrameRate() float32
	// A mode that specifies the handling of video frames that contain multiple fields.
	VideoFieldMode() string
	SetVideoFieldMode(value string)
	// A video field mode that requests deinterlacing of video fields.
	AVPlayerItemTrackVideoFieldModeDeinterlaceFields() string

	// Topic: Accessing the asset track

	// An asset track that provides the media for the player item track.
	AssetTrack() IAVAssetTrack
}

// Init initializes the instance.
func (p AVPlayerItemTrack) Init() AVPlayerItemTrack {
	rv := objc.Send[AVPlayerItemTrack](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerItemTrack) Autorelease() AVPlayerItemTrack {
	rv := objc.Send[AVPlayerItemTrack](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemTrack creates a new AVPlayerItemTrack instance.
func NewAVPlayerItemTrack() AVPlayerItemTrack {
	class := getAVPlayerItemTrackClass()
	rv := objc.Send[AVPlayerItemTrack](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether the player item presents the
// track’s media during playback.
//
// # Discussion
//
// This property isn’t key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrack/isEnabled
func (p AVPlayerItemTrack) Enabled() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isEnabled"))
	return rv
}
func (p AVPlayerItemTrack) SetEnabled(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setEnabled:"), value)
}

// The current frame rate of the video track as it plays.
//
// # Discussion
//
// If the media type of the [AssetTrack] is [video], the property indicates
// the current frame rate of the track as it plays, in frames per second. If
// the item isn’t playing, or if the media type of the track isn’t video,
// the value of this property is `0.0`.
//
// This property isn’t key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrack/currentVideoFrameRate
//
// [video]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/video
func (p AVPlayerItemTrack) CurrentVideoFrameRate() float32 {
	rv := objc.Send[float32](p.ID, objc.Sel("currentVideoFrameRate"))
	return rv
}

// A mode that specifies the handling of video frames that contain multiple
// fields.
//
// # Discussion
//
// A value of `nil` indicates default processing of video frames. To
// deinterlace video fields, set this property value to
// [AVPlayerItemTrackVideoFieldModeDeinterlaceFields].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrack/videoFieldMode
func (p AVPlayerItemTrack) VideoFieldMode() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("videoFieldMode"))
	return foundation.NSStringFromID(rv).String()
}
func (p AVPlayerItemTrack) SetVideoFieldMode(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setVideoFieldMode:"), objc.String(value))
}

// A video field mode that requests deinterlacing of video fields.
//
// See: https://developer.apple.com/documentation/avfoundation/avplayeritemtrackvideofieldmodedeinterlacefields
func (p AVPlayerItemTrack) AVPlayerItemTrackVideoFieldModeDeinterlaceFields() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("AVPlayerItemTrackVideoFieldModeDeinterlaceFields"))
	return foundation.NSStringFromID(rv).String()
}

// An asset track that provides the media for the player item track.
//
// # Discussion
//
// This property isn’t key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrack/assetTrack
func (p AVPlayerItemTrack) AssetTrack() IAVAssetTrack {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("assetTrack"))
	return AVAssetTrackFromID(objc.ID(rv))
}
