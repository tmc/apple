// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMutableAudioMixInputParameters] class.
var (
	_AVMutableAudioMixInputParametersClass     AVMutableAudioMixInputParametersClass
	_AVMutableAudioMixInputParametersClassOnce sync.Once
)

func getAVMutableAudioMixInputParametersClass() AVMutableAudioMixInputParametersClass {
	_AVMutableAudioMixInputParametersClassOnce.Do(func() {
		_AVMutableAudioMixInputParametersClass = AVMutableAudioMixInputParametersClass{class: objc.GetClass("AVMutableAudioMixInputParameters")}
	})
	return _AVMutableAudioMixInputParametersClass
}

// GetAVMutableAudioMixInputParametersClass returns the class object for AVMutableAudioMixInputParameters.
func GetAVMutableAudioMixInputParametersClass() AVMutableAudioMixInputParametersClass {
	return getAVMutableAudioMixInputParametersClass()
}

type AVMutableAudioMixInputParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMutableAudioMixInputParametersClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMutableAudioMixInputParametersClass) Alloc() AVMutableAudioMixInputParameters {
	rv := objc.Send[AVMutableAudioMixInputParameters](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// The parameters you use when adding an audio track to a mix.
//
// # Setting the volume
//
//   - [AVMutableAudioMixInputParameters.SetVolumeAtTime]: Sets the value of the audio volume starting at the specified time.
//   - [AVMutableAudioMixInputParameters.SetVolumeRampFromStartVolumeToEndVolumeTimeRange]: Sets a volume ramp to apply during a specified time range.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters
type AVMutableAudioMixInputParameters struct {
	AVAudioMixInputParameters
}

// AVMutableAudioMixInputParametersFromID constructs a [AVMutableAudioMixInputParameters] from an objc.ID.
//
// The parameters you use when adding an audio track to a mix.
func AVMutableAudioMixInputParametersFromID(id objc.ID) AVMutableAudioMixInputParameters {
	return AVMutableAudioMixInputParameters{AVAudioMixInputParameters: AVAudioMixInputParametersFromID(id)}
}
// NOTE: AVMutableAudioMixInputParameters adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMutableAudioMixInputParameters] class.
//
// # Setting the volume
//
//   - [IAVMutableAudioMixInputParameters.SetVolumeAtTime]: Sets the value of the audio volume starting at the specified time.
//   - [IAVMutableAudioMixInputParameters.SetVolumeRampFromStartVolumeToEndVolumeTimeRange]: Sets a volume ramp to apply during a specified time range.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters
type IAVMutableAudioMixInputParameters interface {
	IAVAudioMixInputParameters

	// Topic: Setting the volume

	// Sets the value of the audio volume starting at the specified time.
	SetVolumeAtTime(volume float32, time uintptr)
	// Sets a volume ramp to apply during a specified time range.
	SetVolumeRampFromStartVolumeToEndVolumeTimeRange(startVolume float32, endVolume float32, timeRange uintptr)
}

// Init initializes the instance.
func (m AVMutableAudioMixInputParameters) Init() AVMutableAudioMixInputParameters {
	rv := objc.Send[AVMutableAudioMixInputParameters](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMutableAudioMixInputParameters) Autorelease() AVMutableAudioMixInputParameters {
	rv := objc.Send[AVMutableAudioMixInputParameters](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMutableAudioMixInputParameters creates a new AVMutableAudioMixInputParameters instance.
func NewAVMutableAudioMixInputParameters() AVMutableAudioMixInputParameters {
	class := getAVMutableAudioMixInputParametersClass()
	rv := objc.Send[AVMutableAudioMixInputParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a mutable input parameters object for a given track.
//
// track: The track to associate with the input parameters object.
//
// # Return Value
// 
// A mutable input parameters object with no volume ramps and [TrackID] set to
// `track`’s identifier.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/init(track:)
func NewMutableAudioMixInputParametersWithTrack(track IAVAssetTrack) AVMutableAudioMixInputParameters {
	rv := objc.Send[objc.ID](objc.ID(getAVMutableAudioMixInputParametersClass().class), objc.Sel("audioMixInputParametersWithTrack:"), track)
	return AVMutableAudioMixInputParametersFromID(rv)
}

// Sets the value of the audio volume starting at the specified time.
//
// volume: The volume. The value must be between `0.0` and `1.0`.
//
// time: The start time at which to set the volume.
//
// # Discussion
// 
// This method adds a volume ramp starting at `time`. This volume setting
// remains in effect until the end of the track unless you set a different
// volume level to start at a later time.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/setVolume(_:at:)
func (m AVMutableAudioMixInputParameters) SetVolumeAtTime(volume float32, time uintptr) {
	objc.Send[objc.ID](m.ID, objc.Sel("setVolume:atTime:"), volume, time)
}
// Sets a volume ramp to apply during a specified time range.
//
// startVolume: The starting volume. The value must be between `0.0` and 1.0.
//
// endVolume: The end volume. The value must be between `0.0` and `1.0`.
//
// timeRange: The time range over which to apply the ramp.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/setVolumeRamp(fromStartVolume:toEndVolume:timeRange:)
func (m AVMutableAudioMixInputParameters) SetVolumeRampFromStartVolumeToEndVolumeTimeRange(startVolume float32, endVolume float32, timeRange uintptr) {
	objc.Send[objc.ID](m.ID, objc.Sel("setVolumeRampFromStartVolume:toEndVolume:timeRange:"), startVolume, endVolume, timeRange)
}

// Creates a mutable input parameters object.
//
// # Return Value
// 
// A mutable input parameters object with no volume ramps and [TrackID]
// initialized to [kCMPersistentTrackID_Invalid].
//
// [kCMPersistentTrackID_Invalid]: https://developer.apple.com/documentation/CoreMedia/kCMPersistentTrackID_Invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableAudioMixInputParameters/audioMixInputParameters
func (_AVMutableAudioMixInputParametersClass AVMutableAudioMixInputParametersClass) AudioMixInputParameters() AVMutableAudioMixInputParameters {
	rv := objc.Send[objc.ID](objc.ID(_AVMutableAudioMixInputParametersClass.class), objc.Sel("audioMixInputParameters"))
	return AVMutableAudioMixInputParametersFromID(rv)
}

