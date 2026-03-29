// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioTime] class.
var (
	_AVAudioTimeClass     AVAudioTimeClass
	_AVAudioTimeClassOnce sync.Once
)

func getAVAudioTimeClass() AVAudioTimeClass {
	_AVAudioTimeClassOnce.Do(func() {
		_AVAudioTimeClass = AVAudioTimeClass{class: objc.GetClass("AVAudioTime")}
	})
	return _AVAudioTimeClass
}

// GetAVAudioTimeClass returns the class object for AVAudioTime.
func GetAVAudioTimeClass() AVAudioTimeClass {
	return getAVAudioTimeClass()
}

type AVAudioTimeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioTimeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioTimeClass) Alloc() AVAudioTime {
	rv := objc.Send[AVAudioTime](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object you use to represent a moment in time.
//
// # Overview
// 
// The [AVAudioTime] object represents a single moment in time in two ways:
// 
// - As host time, using the system’s basic clock with
// `mach_absolute_time()` - As audio samples at a particular sample rate
// 
// A single [AVAudioTime] instance contains either or both representations,
// meaning it might represent only a sample time, a host time, or both.
// 
// Instances of this class are immutable.
//
// # Creating an Audio Time Instance
//
//   - [AVAudioTime.InitWithAudioTimeStampSampleRate]: Creates an audio time object with the specified timestamp and sample rate.
//   - [AVAudioTime.InitWithHostTime]: Creates an audio time object with the specified host time.
//   - [AVAudioTime.InitWithHostTimeSampleTimeAtRate]: Creates an audio time object with the specified host time, sample time, and sample rate.
//   - [AVAudioTime.InitWithSampleTimeAtRate]: Creates an audio time object with the specified timestamp and sample rate.
//   - [AVAudioTime.ExtrapolateTimeFromAnchor]: Creates an audio time object by converting between host time and sample time.
//
// # Manipulating Host Time
//
//   - [AVAudioTime.HostTime]: The host time.
//   - [AVAudioTime.HostTimeValid]: A Boolean value that indicates whether the host time value is valid.
//
// # Getting Sample Rate Information
//
//   - [AVAudioTime.SampleRate]: The sampling rate that the sample time property expresses.
//   - [AVAudioTime.SampleTime]: The time as a number of audio samples that the current audio device tracks.
//   - [AVAudioTime.SampleTimeValid]: A Boolean value that indicates whether the sample time and sample rate properties are in a valid state.
//
// # Getting the Core Audio Time Stamp
//
//   - [AVAudioTime.AudioTimeStamp]: The time as an audio timestamp.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime
type AVAudioTime struct {
	objectivec.Object
}

// AVAudioTimeFromID constructs a [AVAudioTime] from an objc.ID.
//
// An object you use to represent a moment in time.
func AVAudioTimeFromID(id objc.ID) AVAudioTime {
	return AVAudioTime{objectivec.Object{ID: id}}
}
// NOTE: AVAudioTime adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioTime] class.
//
// # Creating an Audio Time Instance
//
//   - [IAVAudioTime.InitWithAudioTimeStampSampleRate]: Creates an audio time object with the specified timestamp and sample rate.
//   - [IAVAudioTime.InitWithHostTime]: Creates an audio time object with the specified host time.
//   - [IAVAudioTime.InitWithHostTimeSampleTimeAtRate]: Creates an audio time object with the specified host time, sample time, and sample rate.
//   - [IAVAudioTime.InitWithSampleTimeAtRate]: Creates an audio time object with the specified timestamp and sample rate.
//   - [IAVAudioTime.ExtrapolateTimeFromAnchor]: Creates an audio time object by converting between host time and sample time.
//
// # Manipulating Host Time
//
//   - [IAVAudioTime.HostTime]: The host time.
//   - [IAVAudioTime.HostTimeValid]: A Boolean value that indicates whether the host time value is valid.
//
// # Getting Sample Rate Information
//
//   - [IAVAudioTime.SampleRate]: The sampling rate that the sample time property expresses.
//   - [IAVAudioTime.SampleTime]: The time as a number of audio samples that the current audio device tracks.
//   - [IAVAudioTime.SampleTimeValid]: A Boolean value that indicates whether the sample time and sample rate properties are in a valid state.
//
// # Getting the Core Audio Time Stamp
//
//   - [IAVAudioTime.AudioTimeStamp]: The time as an audio timestamp.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime
type IAVAudioTime interface {
	objectivec.IObject

	// Topic: Creating an Audio Time Instance

	// Creates an audio time object with the specified timestamp and sample rate.
	InitWithAudioTimeStampSampleRate(ts objectivec.IObject, sampleRate float64) AVAudioTime
	// Creates an audio time object with the specified host time.
	InitWithHostTime(hostTime uint64) AVAudioTime
	// Creates an audio time object with the specified host time, sample time, and sample rate.
	InitWithHostTimeSampleTimeAtRate(hostTime uint64, sampleTime AVAudioFramePosition, sampleRate float64) AVAudioTime
	// Creates an audio time object with the specified timestamp and sample rate.
	InitWithSampleTimeAtRate(sampleTime AVAudioFramePosition, sampleRate float64) AVAudioTime
	// Creates an audio time object by converting between host time and sample time.
	ExtrapolateTimeFromAnchor(anchorTime IAVAudioTime) IAVAudioTime

	// Topic: Manipulating Host Time

	// The host time.
	HostTime() uint64
	// A Boolean value that indicates whether the host time value is valid.
	HostTimeValid() bool

	// Topic: Getting Sample Rate Information

	// The sampling rate that the sample time property expresses.
	SampleRate() float64
	// The time as a number of audio samples that the current audio device tracks.
	SampleTime() AVAudioFramePosition
	// A Boolean value that indicates whether the sample time and sample rate properties are in a valid state.
	SampleTimeValid() bool

	// Topic: Getting the Core Audio Time Stamp

	// The time as an audio timestamp.
	AudioTimeStamp() objectivec.IObject
}

// Init initializes the instance.
func (a AVAudioTime) Init() AVAudioTime {
	rv := objc.Send[AVAudioTime](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioTime) Autorelease() AVAudioTime {
	rv := objc.Send[AVAudioTime](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioTime creates a new AVAudioTime instance.
func NewAVAudioTime() AVAudioTime {
	class := getAVAudioTimeClass()
	rv := objc.Send[AVAudioTime](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an audio time object with the specified timestamp and sample rate.
//
// ts: The timestamp.
//
// sampleRate: The sample rate.
//
// # Return Value
// 
// A new [AVAudioTime] instance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/init(audioTimeStamp:sampleRate:)
func NewAudioTimeWithAudioTimeStampSampleRate(ts objectivec.IObject, sampleRate float64) AVAudioTime {
	instance := getAVAudioTimeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAudioTimeStamp:sampleRate:"), ts, sampleRate)
	return AVAudioTimeFromID(rv)
}

// Creates an audio time object with the specified host time.
//
// hostTime: The host time.
//
// # Return Value
// 
// A new [AVAudioTime] instance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/init(hostTime:)
func NewAudioTimeWithHostTime(hostTime uint64) AVAudioTime {
	instance := getAVAudioTimeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithHostTime:"), hostTime)
	return AVAudioTimeFromID(rv)
}

// Creates an audio time object with the specified host time, sample time, and
// sample rate.
//
// hostTime: The host time.
//
// sampleTime: The sample time.
//
// sampleRate: The sample rate.
//
// # Return Value
// 
// A new [AVAudioTime] instance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/init(hostTime:sampleTime:atRate:)
func NewAudioTimeWithHostTimeSampleTimeAtRate(hostTime uint64, sampleTime AVAudioFramePosition, sampleRate float64) AVAudioTime {
	instance := getAVAudioTimeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithHostTime:sampleTime:atRate:"), hostTime, sampleTime, sampleRate)
	return AVAudioTimeFromID(rv)
}

// Creates an audio time object with the specified timestamp and sample rate.
//
// sampleTime: The sample time.
//
// sampleRate: The sample rate.
//
// # Return Value
// 
// A new [AVAudioTime] instance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/init(sampleTime:atRate:)
func NewAudioTimeWithSampleTimeAtRate(sampleTime AVAudioFramePosition, sampleRate float64) AVAudioTime {
	instance := getAVAudioTimeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSampleTime:atRate:"), sampleTime, sampleRate)
	return AVAudioTimeFromID(rv)
}

// Creates an audio time object with the specified timestamp and sample rate.
//
// ts: The timestamp.
//
// sampleRate: The sample rate.
//
// # Return Value
// 
// A new [AVAudioTime] instance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/init(audioTimeStamp:sampleRate:)
func (a AVAudioTime) InitWithAudioTimeStampSampleRate(ts objectivec.IObject, sampleRate float64) AVAudioTime {
	rv := objc.Send[AVAudioTime](a.ID, objc.Sel("initWithAudioTimeStamp:sampleRate:"), ts, sampleRate)
	return rv
}
// Creates an audio time object with the specified host time.
//
// hostTime: The host time.
//
// # Return Value
// 
// A new [AVAudioTime] instance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/init(hostTime:)
func (a AVAudioTime) InitWithHostTime(hostTime uint64) AVAudioTime {
	rv := objc.Send[AVAudioTime](a.ID, objc.Sel("initWithHostTime:"), hostTime)
	return rv
}
// Creates an audio time object with the specified host time, sample time, and
// sample rate.
//
// hostTime: The host time.
//
// sampleTime: The sample time.
//
// sampleRate: The sample rate.
//
// # Return Value
// 
// A new [AVAudioTime] instance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/init(hostTime:sampleTime:atRate:)
func (a AVAudioTime) InitWithHostTimeSampleTimeAtRate(hostTime uint64, sampleTime AVAudioFramePosition, sampleRate float64) AVAudioTime {
	rv := objc.Send[AVAudioTime](a.ID, objc.Sel("initWithHostTime:sampleTime:atRate:"), hostTime, sampleTime, sampleRate)
	return rv
}
// Creates an audio time object with the specified timestamp and sample rate.
//
// sampleTime: The sample time.
//
// sampleRate: The sample rate.
//
// # Return Value
// 
// A new [AVAudioTime] instance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/init(sampleTime:atRate:)
func (a AVAudioTime) InitWithSampleTimeAtRate(sampleTime AVAudioFramePosition, sampleRate float64) AVAudioTime {
	rv := objc.Send[AVAudioTime](a.ID, objc.Sel("initWithSampleTime:atRate:"), sampleTime, sampleRate)
	return rv
}
// Creates an audio time object by converting between host time and sample
// time.
//
// anchorTime: An audio time instance with a more complete timestamp than that of the
// receiver (self).
//
// # Return Value
// 
// A new [AVAudioTime] instance.
//
// # Discussion
// 
// If `anchorTime` is an [AVAudioTime] instance where both host time and
// sample time are valid, and the receiver is another timestamp where only one
// of the two is valid, this method returns a new [AVAudioTime]. It copies it
// from the receiver, where the anchor provides additional valid fields.
// 
// The `anchorTime` value must have a valid host time and sample time, and
// self must have sample rate and at least one valid host time or sample time.
// Otherwise, this method returns `nil`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/extrapolateTime(fromAnchor:)
func (a AVAudioTime) ExtrapolateTimeFromAnchor(anchorTime IAVAudioTime) IAVAudioTime {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("extrapolateTimeFromAnchor:"), anchorTime)
	return AVAudioTimeFromID(rv)
}

// Converts seconds to host time.
//
// seconds: The number of seconds.
//
// # Return Value
// 
// The host time that represents the seconds you specify.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/hostTime(forSeconds:)
func (_AVAudioTimeClass AVAudioTimeClass) HostTimeForSeconds(seconds float64) uint64 {
	rv := objc.Send[uint64](objc.ID(_AVAudioTimeClass.class), objc.Sel("hostTimeForSeconds:"), seconds)
	return rv
}
// Converts host time to seconds.
//
// hostTime: The host time.
//
// # Return Value
// 
// The number of seconds that represent the host time you specify.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/seconds(forHostTime:)
func (_AVAudioTimeClass AVAudioTimeClass) SecondsForHostTime(hostTime uint64) float64 {
	rv := objc.Send[float64](objc.ID(_AVAudioTimeClass.class), objc.Sel("secondsForHostTime:"), hostTime)
	return rv
}
// Creates an audio time object with the specified timestamp and sample rate.
//
// ts: The timestamp.
//
// sampleRate: The sample rate.
//
// # Return Value
// 
// A new [AVAudioTime] instance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/timeWithAudioTimeStamp:sampleRate:
func (_AVAudioTimeClass AVAudioTimeClass) TimeWithAudioTimeStampSampleRate(ts objectivec.IObject, sampleRate float64) AVAudioTime {
	rv := objc.Send[objc.ID](objc.ID(_AVAudioTimeClass.class), objc.Sel("timeWithAudioTimeStamp:sampleRate:"), ts, sampleRate)
	return AVAudioTimeFromID(rv)
}
// Creates an audio time object with the specified host time.
//
// hostTime: The host time.
//
// # Return Value
// 
// A new [AVAudioTime] instance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/timeWithHostTime:
func (_AVAudioTimeClass AVAudioTimeClass) TimeWithHostTime(hostTime uint64) AVAudioTime {
	rv := objc.Send[objc.ID](objc.ID(_AVAudioTimeClass.class), objc.Sel("timeWithHostTime:"), hostTime)
	return AVAudioTimeFromID(rv)
}
// Creates an audio time object with the specified host time, sample time, and
// sample rate.
//
// hostTime: The host time.
//
// sampleTime: The sample time.
//
// sampleRate: The sample rate.
//
// # Return Value
// 
// A new [AVAudioTime] instance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/timeWithHostTime:sampleTime:atRate:
func (_AVAudioTimeClass AVAudioTimeClass) TimeWithHostTimeSampleTimeAtRate(hostTime uint64, sampleTime AVAudioFramePosition, sampleRate float64) AVAudioTime {
	rv := objc.Send[objc.ID](objc.ID(_AVAudioTimeClass.class), objc.Sel("timeWithHostTime:sampleTime:atRate:"), hostTime, sampleTime, sampleRate)
	return AVAudioTimeFromID(rv)
}
// Creates an audio time object with the specified sample time and sample
// rate.
//
// sampleTime: The sample time.
//
// sampleRate: The sample rate.
//
// # Return Value
// 
// A new [AVAudioTime] instance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/timeWithSampleTime:atRate:
func (_AVAudioTimeClass AVAudioTimeClass) TimeWithSampleTimeAtRate(sampleTime AVAudioFramePosition, sampleRate float64) AVAudioTime {
	rv := objc.Send[objc.ID](objc.ID(_AVAudioTimeClass.class), objc.Sel("timeWithSampleTime:atRate:"), sampleTime, sampleRate)
	return AVAudioTimeFromID(rv)
}

// The host time.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/hostTime
func (a AVAudioTime) HostTime() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("hostTime"))
	return rv
}
// A Boolean value that indicates whether the host time value is valid.
//
// # Discussion
// 
// This property returns [true] if the [HostTime] property is valid;
// otherwise, it returns [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/isHostTimeValid
func (a AVAudioTime) HostTimeValid() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isHostTimeValid"))
	return rv
}
// The sampling rate that the sample time property expresses.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/sampleRate
func (a AVAudioTime) SampleRate() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("sampleRate"))
	return rv
}
// The time as a number of audio samples that the current audio device tracks.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/sampleTime
func (a AVAudioTime) SampleTime() AVAudioFramePosition {
	rv := objc.Send[AVAudioFramePosition](a.ID, objc.Sel("sampleTime"))
	return AVAudioFramePosition(rv)
}
// A Boolean value that indicates whether the sample time and sample rate
// properties are in a valid state.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/isSampleTimeValid
func (a AVAudioTime) SampleTimeValid() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isSampleTimeValid"))
	return rv
}
// The time as an audio timestamp.
//
// # Discussion
// 
// This is useful for compatibility with lower-level Core Audio and Audio
// Toolbox API.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioTime/audioTimeStamp
func (a AVAudioTime) AudioTimeStamp() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("audioTimeStamp"))
	return objectivec.Object{ID: rv}
}

