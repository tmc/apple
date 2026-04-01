// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioMixInputParameters] class.
var (
	_AVAudioMixInputParametersClass     AVAudioMixInputParametersClass
	_AVAudioMixInputParametersClassOnce sync.Once
)

func getAVAudioMixInputParametersClass() AVAudioMixInputParametersClass {
	_AVAudioMixInputParametersClassOnce.Do(func() {
		_AVAudioMixInputParametersClass = AVAudioMixInputParametersClass{class: objc.GetClass("AVAudioMixInputParameters")}
	})
	return _AVAudioMixInputParametersClass
}

// GetAVAudioMixInputParametersClass returns the class object for AVAudioMixInputParameters.
func GetAVAudioMixInputParametersClass() AVAudioMixInputParametersClass {
	return getAVAudioMixInputParametersClass()
}

type AVAudioMixInputParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioMixInputParametersClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioMixInputParametersClass) Alloc() AVAudioMixInputParameters {
	rv := objc.Send[AVAudioMixInputParameters](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the parameters that you apply when adding an
// audio track to a mix.
//
// # Overview
//
// You use an instance [AVAudioMixInputParameters] to apply audio volume ramps
// for an input to an audio mix. Mix parameters are associated with audio
// tracks via the [AVAudioMixInputParameters.TrackID] property.
//
// Audio volume is currently supported as a time-varying parameter.
// [AVAudioMixInputParameters] has a mutable subclass,
// [AVMutableAudioMixInputParameters].
//
// Before the first time at which a volume is set, a volume of 1.0 used; after
// the last time for which a volume has been set, the last volume is used.
// Within the time range of a volume ramp, the volume is interpolated between
// the start volume and end volume of the ramp. For example, setting the
// volume to 1.0 at time 0 and also setting a volume ramp from a volume of 0.5
// to 0.2 with a timeRange of [4.0, 5.0] results in an audio volume parameters
// that hold the volume constant at 1.0 from 0.0 sec to 4.0 sec, then cause it
// to jump to 0.5 and descend to 0.2 from 4.0 sec to 9.0 sec, holding constant
// at 0.2 thereafter.
//
// Given that this is an immutable variant of the object, you should not
// allocate and initialize a version of this class yourself. Other classes may
// return instances of this class.
//
// # Getting the track ID
//
//   - [AVAudioMixInputParameters.TrackID]: The identifier of the audio track to which the parameters should be applied.
//
// # Getting volume ramps
//
//   - [AVAudioMixInputParameters.GetVolumeRampForTimeStartVolumeEndVolumeTimeRange]: Retrieves the volume ramp that includes the specified time.
//
// # Getting the time pitch algorithm setting
//
//   - [AVAudioMixInputParameters.AudioTimePitchAlgorithm]: The processing algorithm used to manage audio pitch for scaled audio edits.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAudioMixInputParameters
type AVAudioMixInputParameters struct {
	objectivec.Object
}

// AVAudioMixInputParametersFromID constructs a [AVAudioMixInputParameters] from an objc.ID.
//
// An object that represents the parameters that you apply when adding an
// audio track to a mix.
func AVAudioMixInputParametersFromID(id objc.ID) AVAudioMixInputParameters {
	return AVAudioMixInputParameters{objectivec.Object{ID: id}}
}

// NOTE: AVAudioMixInputParameters adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioMixInputParameters] class.
//
// # Getting the track ID
//
//   - [IAVAudioMixInputParameters.TrackID]: The identifier of the audio track to which the parameters should be applied.
//
// # Getting volume ramps
//
//   - [IAVAudioMixInputParameters.GetVolumeRampForTimeStartVolumeEndVolumeTimeRange]: Retrieves the volume ramp that includes the specified time.
//
// # Getting the time pitch algorithm setting
//
//   - [IAVAudioMixInputParameters.AudioTimePitchAlgorithm]: The processing algorithm used to manage audio pitch for scaled audio edits.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAudioMixInputParameters
type IAVAudioMixInputParameters interface {
	objectivec.IObject

	// Topic: Getting the track ID

	// The identifier of the audio track to which the parameters should be applied.
	TrackID() int32

	// Topic: Getting volume ramps

	// Retrieves the volume ramp that includes the specified time.
	GetVolumeRampForTimeStartVolumeEndVolumeTimeRange(time coremedia.CMTime, timeRange *coremedia.CMTimeRange) (float32, float32, bool)

	// Topic: Getting the time pitch algorithm setting

	// The processing algorithm used to manage audio pitch for scaled audio edits.
	AudioTimePitchAlgorithm() AVAudioTimePitchAlgorithm
}

// Init initializes the instance.
func (a AVAudioMixInputParameters) Init() AVAudioMixInputParameters {
	rv := objc.Send[AVAudioMixInputParameters](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioMixInputParameters) Autorelease() AVAudioMixInputParameters {
	rv := objc.Send[AVAudioMixInputParameters](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioMixInputParameters creates a new AVAudioMixInputParameters instance.
func NewAVAudioMixInputParameters() AVAudioMixInputParameters {
	class := getAVAudioMixInputParametersClass()
	rv := objc.Send[AVAudioMixInputParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Retrieves the volume ramp that includes the specified time.
//
// time: If a ramp with a time range that contains the specified time has been set,
// information about the effective ramp for that time is supplied. Otherwise,
// information about the first ramp that starts after the specified time is
// supplied.
//
// startVolume: A pointer to a float to receive the starting volume value for the volume
// ramp.
//
// This value may be [NULL].
//
// endVolume: A pointer to a float to receive the ending volume value for the volume
// ramp.
//
// This value may be [NULL].
//
// timeRange: A pointer to a [CMTimeRange] to receive the time range of the volume ramp.
//
// This value may be [NULL].
//
// # Return Value
//
// true if the values were retrieved successfully, otherwise false. Returns
// false if `time` is beyond the duration of the last volume ramp that has
// been set.
//
// # Discussion
//
// The process of setting up volume ramps requires the configuration of an
// instance of [AVMutableAudioMixInputParameters].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAudioMixInputParameters/getVolumeRamp(for:startVolume:endVolume:timeRange:)
//
// [CMTimeRange]: https://developer.apple.com/documentation/CoreMedia/CMTimeRange
func (a AVAudioMixInputParameters) GetVolumeRampForTimeStartVolumeEndVolumeTimeRange(time coremedia.CMTime, timeRange *coremedia.CMTimeRange) (float32, float32, bool) {
	var startVolume float32
	var endVolume float32
	rv := objc.Send[bool](a.ID, objc.Sel("getVolumeRampForTime:startVolume:endVolume:timeRange:"), time, unsafe.Pointer(&startVolume), unsafe.Pointer(&endVolume), timeRange)
	return startVolume, endVolume, rv
}

// The identifier of the audio track to which the parameters should be
// applied.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAudioMixInputParameters/trackID
func (a AVAudioMixInputParameters) TrackID() int32 {
	rv := objc.Send[int32](a.ID, objc.Sel("trackID"))
	return rv
}

// The processing algorithm used to manage audio pitch for scaled audio edits.
//
// # Discussion
//
// The supported constants are defined in Time Pitch Algorithm Settings. An
// [InvalidArgumentException] will be raised if this property is set to a
// value other than the defined constants.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAudioMixInputParameters/audioTimePitchAlgorithm
func (a AVAudioMixInputParameters) AudioTimePitchAlgorithm() AVAudioTimePitchAlgorithm {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("audioTimePitchAlgorithm"))
	return AVAudioTimePitchAlgorithm(foundation.NSStringFromID(rv).String())
}
