// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioRecorder] class.
var (
	_AVAudioRecorderClass     AVAudioRecorderClass
	_AVAudioRecorderClassOnce sync.Once
)

func getAVAudioRecorderClass() AVAudioRecorderClass {
	_AVAudioRecorderClassOnce.Do(func() {
		_AVAudioRecorderClass = AVAudioRecorderClass{class: objc.GetClass("AVAudioRecorder")}
	})
	return _AVAudioRecorderClass
}

// GetAVAudioRecorderClass returns the class object for AVAudioRecorder.
func GetAVAudioRecorderClass() AVAudioRecorderClass {
	return getAVAudioRecorderClass()
}

type AVAudioRecorderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioRecorderClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioRecorderClass) Alloc() AVAudioRecorder {
	rv := objc.Send[AVAudioRecorder](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that records audio data to a file.
//
// # Overview
// 
// Use an audio recorder to:
// 
// - Record audio from the system’s active input device - Record for a
// specified duration or until the user stops recording - Pause and resume a
// recording - Access recording-level metering data
// 
// To record audio in iOS or tvOS, configure your audio session to use the
// [record] or [playAndRecord] category.
//
// [playAndRecord]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/Category-swift.struct/playAndRecord
// [record]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/Category-swift.struct/record
//
// # Creating an audio recorder
//
//   - [AVAudioRecorder.InitWithURLSettingsError]: Creates an audio recorder with settings.
//   - [AVAudioRecorder.InitWithURLFormatError]: Creates an audio recorder with an audio format.
//
// # Controlling recording
//
//   - [AVAudioRecorder.PrepareToRecord]: Creates an audio file and prepares the system for recording.
//   - [AVAudioRecorder.RecordAtTime]: Records audio starting at a specific time.
//   - [AVAudioRecorder.RecordForDuration]: Records audio for the indicated duration of time.
//   - [AVAudioRecorder.RecordAtTimeForDuration]: Records audio starting at a specific time for the indicated duration.
//   - [AVAudioRecorder.Pause]: Pauses an audio recording.
//   - [AVAudioRecorder.Stop]: Stops recording and closes the audio file.
//   - [AVAudioRecorder.Recording]: A Boolean value that indicates whether the audio recorder is recording.
//   - [AVAudioRecorder.DeleteRecording]: Deletes a recorded audio file.
//
// # Accessing recorder timing
//
//   - [AVAudioRecorder.CurrentTime]: The time, in seconds, since the beginning of the recording.
//   - [AVAudioRecorder.DeviceCurrentTime]: The time, in seconds, of the host audio device.
//
// # Managing audio-level metering
//
//   - [AVAudioRecorder.MeteringEnabled]: A Boolean value that indicates whether you’ve enabled the recorder to generate audio-level metering data.
//   - [AVAudioRecorder.SetMeteringEnabled]
//   - [AVAudioRecorder.UpdateMeters]: Refreshes the average and peak power values for all channels of an audio recorder.
//   - [AVAudioRecorder.AveragePowerForChannel]: Returns the average power, in decibels full-scale (dBFS), for an audio channel.
//   - [AVAudioRecorder.PeakPowerForChannel]: Returns the peak power, in decibels full-scale (dBFS), for an audio channel.
//
// # Responding to recorder events
//
//   - [AVAudioRecorder.Delegate]: The delegate object for the audio recorder.
//   - [AVAudioRecorder.SetDelegate]
//
// # Inspecting the audio data
//
//   - [AVAudioRecorder.Url]: The URL to which the recorder writes its data.
//   - [AVAudioRecorder.Format]: The format of the recorded audio.
//   - [AVAudioRecorder.Settings]: The settings that describe the format of the recorded audio.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder
type AVAudioRecorder struct {
	objectivec.Object
}

// AVAudioRecorderFromID constructs a [AVAudioRecorder] from an objc.ID.
//
// An object that records audio data to a file.
func AVAudioRecorderFromID(id objc.ID) AVAudioRecorder {
	return AVAudioRecorder{objectivec.Object{ID: id}}
}
// NOTE: AVAudioRecorder adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioRecorder] class.
//
// # Creating an audio recorder
//
//   - [IAVAudioRecorder.InitWithURLSettingsError]: Creates an audio recorder with settings.
//   - [IAVAudioRecorder.InitWithURLFormatError]: Creates an audio recorder with an audio format.
//
// # Controlling recording
//
//   - [IAVAudioRecorder.PrepareToRecord]: Creates an audio file and prepares the system for recording.
//   - [IAVAudioRecorder.RecordAtTime]: Records audio starting at a specific time.
//   - [IAVAudioRecorder.RecordForDuration]: Records audio for the indicated duration of time.
//   - [IAVAudioRecorder.RecordAtTimeForDuration]: Records audio starting at a specific time for the indicated duration.
//   - [IAVAudioRecorder.Pause]: Pauses an audio recording.
//   - [IAVAudioRecorder.Stop]: Stops recording and closes the audio file.
//   - [IAVAudioRecorder.Recording]: A Boolean value that indicates whether the audio recorder is recording.
//   - [IAVAudioRecorder.DeleteRecording]: Deletes a recorded audio file.
//
// # Accessing recorder timing
//
//   - [IAVAudioRecorder.CurrentTime]: The time, in seconds, since the beginning of the recording.
//   - [IAVAudioRecorder.DeviceCurrentTime]: The time, in seconds, of the host audio device.
//
// # Managing audio-level metering
//
//   - [IAVAudioRecorder.MeteringEnabled]: A Boolean value that indicates whether you’ve enabled the recorder to generate audio-level metering data.
//   - [IAVAudioRecorder.SetMeteringEnabled]
//   - [IAVAudioRecorder.UpdateMeters]: Refreshes the average and peak power values for all channels of an audio recorder.
//   - [IAVAudioRecorder.AveragePowerForChannel]: Returns the average power, in decibels full-scale (dBFS), for an audio channel.
//   - [IAVAudioRecorder.PeakPowerForChannel]: Returns the peak power, in decibels full-scale (dBFS), for an audio channel.
//
// # Responding to recorder events
//
//   - [IAVAudioRecorder.Delegate]: The delegate object for the audio recorder.
//   - [IAVAudioRecorder.SetDelegate]
//
// # Inspecting the audio data
//
//   - [IAVAudioRecorder.Url]: The URL to which the recorder writes its data.
//   - [IAVAudioRecorder.Format]: The format of the recorded audio.
//   - [IAVAudioRecorder.Settings]: The settings that describe the format of the recorded audio.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder
type IAVAudioRecorder interface {
	objectivec.IObject

	// Topic: Creating an audio recorder

	// Creates an audio recorder with settings.
	InitWithURLSettingsError(url foundation.INSURL, settings foundation.INSDictionary) (AVAudioRecorder, error)
	// Creates an audio recorder with an audio format.
	InitWithURLFormatError(url foundation.INSURL, format IAVAudioFormat) (AVAudioRecorder, error)

	// Topic: Controlling recording

	// Creates an audio file and prepares the system for recording.
	PrepareToRecord() bool
	// Records audio starting at a specific time.
	RecordAtTime(time float64) bool
	// Records audio for the indicated duration of time.
	RecordForDuration(duration float64) bool
	// Records audio starting at a specific time for the indicated duration.
	RecordAtTimeForDuration(time float64, duration float64) bool
	// Pauses an audio recording.
	Pause()
	// Stops recording and closes the audio file.
	Stop()
	// A Boolean value that indicates whether the audio recorder is recording.
	Recording() bool
	// Deletes a recorded audio file.
	DeleteRecording() bool

	// Topic: Accessing recorder timing

	// The time, in seconds, since the beginning of the recording.
	CurrentTime() float64
	// The time, in seconds, of the host audio device.
	DeviceCurrentTime() float64

	// Topic: Managing audio-level metering

	// A Boolean value that indicates whether you’ve enabled the recorder to generate audio-level metering data.
	MeteringEnabled() bool
	SetMeteringEnabled(value bool)
	// Refreshes the average and peak power values for all channels of an audio recorder.
	UpdateMeters()
	// Returns the average power, in decibels full-scale (dBFS), for an audio channel.
	AveragePowerForChannel(channelNumber uint) float32
	// Returns the peak power, in decibels full-scale (dBFS), for an audio channel.
	PeakPowerForChannel(channelNumber uint) float32

	// Topic: Responding to recorder events

	// The delegate object for the audio recorder.
	Delegate() AVAudioRecorderDelegate
	SetDelegate(value AVAudioRecorderDelegate)

	// Topic: Inspecting the audio data

	// The URL to which the recorder writes its data.
	Url() foundation.INSURL
	// The format of the recorded audio.
	Format() IAVAudioFormat
	// The settings that describe the format of the recorded audio.
	Settings() foundation.INSDictionary

	// The category for recording (input) and playback (output) of audio, such as for a Voice over Internet Protocol (VoIP) app.
	PlayAndRecord() objc.ID
	// The category for recording audio while also silencing playback audio.
	Record() objc.ID
}

// Init initializes the instance.
func (a AVAudioRecorder) Init() AVAudioRecorder {
	rv := objc.Send[AVAudioRecorder](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioRecorder) Autorelease() AVAudioRecorder {
	rv := objc.Send[AVAudioRecorder](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioRecorder creates a new AVAudioRecorder instance.
func NewAVAudioRecorder() AVAudioRecorder {
	class := getAVAudioRecorderClass()
	rv := objc.Send[AVAudioRecorder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an audio recorder with an audio format.
//
// url: The file system location to record to.
//
// format: The audio format to use for the recording.
//
// # Return Value
// 
// A new audio recorder, or [nil] if an error occurred.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/init(url:format:)
func NewAudioRecorderWithURLFormatError(url foundation.INSURL, format IAVAudioFormat) (AVAudioRecorder, error) {
	var errorPtr objc.ID
	instance := getAVAudioRecorderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:format:error:"), url, format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioRecorder{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioRecorderFromID(rv), nil
}

// Creates an audio recorder with settings.
//
// url: The file system location to record to.
//
// settings: The audio settings to use for the recording.
//
// # Return Value
// 
// A new audio recorder, or [nil] if an error occurred.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// # Discussion
// 
// The system supports the following keys when defining the format settings:
// 
// [Table data omitted]
// 
// The system supports additional configuration options based on your selected
// audio format. See [Linear PCM format settings] for information about
// customizing Linear PCM formats and [Encoder settings] for compressed
// formats.
//
// [Encoder settings]: https://developer.apple.com/documentation/AVFoundation/encoder-settings
// [Linear PCM format settings]: https://developer.apple.com/documentation/AVFoundation/linear-pcm-format-settings
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/init(url:settings:)
func NewAudioRecorderWithURLSettingsError(url foundation.INSURL, settings foundation.INSDictionary) (AVAudioRecorder, error) {
	var errorPtr objc.ID
	instance := getAVAudioRecorderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:settings:error:"), url, settings, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioRecorder{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioRecorderFromID(rv), nil
}

// Creates an audio recorder with settings.
//
// url: The file system location to record to.
//
// settings: The audio settings to use for the recording.
//
// # Return Value
// 
// A new audio recorder, or [nil] if an error occurred.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// # Discussion
// 
// The system supports the following keys when defining the format settings:
// 
// [Table data omitted]
// 
// The system supports additional configuration options based on your selected
// audio format. See [Linear PCM format settings] for information about
// customizing Linear PCM formats and [Encoder settings] for compressed
// formats.
//
// [Encoder settings]: https://developer.apple.com/documentation/AVFoundation/encoder-settings
// [Linear PCM format settings]: https://developer.apple.com/documentation/AVFoundation/linear-pcm-format-settings
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/init(url:settings:)
func (a AVAudioRecorder) InitWithURLSettingsError(url foundation.INSURL, settings foundation.INSDictionary) (AVAudioRecorder, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithURL:settings:error:"), url, settings, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioRecorder{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioRecorderFromID(rv), nil

}
// Creates an audio recorder with an audio format.
//
// url: The file system location to record to.
//
// format: The audio format to use for the recording.
//
// # Return Value
// 
// A new audio recorder, or [nil] if an error occurred.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/init(url:format:)
func (a AVAudioRecorder) InitWithURLFormatError(url foundation.INSURL, format IAVAudioFormat) (AVAudioRecorder, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithURL:format:error:"), url, format, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioRecorder{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioRecorderFromID(rv), nil

}
// Creates an audio file and prepares the system for recording.
//
// # Return Value
// 
// [true] if successful; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Calling this method creates an audio file at the URL you used to create the
// recorder. If a file already exists at that location, this method overwrites
// it.
// 
// Call this method to start recording as quickly as possible upon calling
// [Record].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/prepareToRecord()
func (a AVAudioRecorder) PrepareToRecord() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("prepareToRecord"))
	return rv
}
// Records audio starting at a specific time.
//
// time: The time at which to start recording, relative to [DeviceCurrentTime].
//
// # Return Value
// 
// [true] if recording starts successfully; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// You can call this method on a single recorder, or use it to synchronize the
// recording of multiple players as shown below.
// 
// Calling this method implicitly calls [PrepareToRecord], which creates an
// audio file and prepares the system for recording.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/record(atTime:)
func (a AVAudioRecorder) RecordAtTime(time float64) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("recordAtTime:"), time)
	return rv
}
// Records audio for the indicated duration of time.
//
// duration: The duration of time to record, in seconds.
//
// # Return Value
// 
// [true] if successful; otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The recorder stops recording when it reaches the indicated duration.
// 
// Calling this method implicitly calls [PrepareToRecord], which creates an
// audio file and prepares the system for recording.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/record(forDuration:)
func (a AVAudioRecorder) RecordForDuration(duration float64) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("recordForDuration:"), duration)
	return rv
}
// Records audio starting at a specific time for the indicated duration.
//
// time: The time at which to start recording, relative to [DeviceCurrentTime].
//
// duration: The duration of time to record, in seconds.
//
// # Return Value
// 
// [true] if recording was successful; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The recorder automatically stops recording when it reaches the indicated
// duration. You may also use it to synchronize recording of multiple
// recorders as shown below.
// 
// Calling this method implicitly calls [PrepareToRecord], which creates an
// audio file and prepares the system for recording.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/record(atTime:forDuration:)
func (a AVAudioRecorder) RecordAtTimeForDuration(time float64, duration float64) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("recordAtTime:forDuration:"), time, duration)
	return rv
}
// Pauses an audio recording.
//
// # Discussion
// 
// Call [Record] to resume recording.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/pause()
func (a AVAudioRecorder) Pause() {
	objc.Send[objc.ID](a.ID, objc.Sel("pause"))
}
// Stops recording and closes the audio file.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/stop()
func (a AVAudioRecorder) Stop() {
	objc.Send[objc.ID](a.ID, objc.Sel("stop"))
}
// Deletes a recorded audio file.
//
// # Return Value
// 
// [true] if the system deleted the file; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// You must stop the audio recorder before calling this method.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/deleteRecording()
func (a AVAudioRecorder) DeleteRecording() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("deleteRecording"))
	return rv
}
// Refreshes the average and peak power values for all channels of an audio
// recorder.
//
// # Discussion
// 
// Call this method to update the level meter data before calling
// [AveragePowerForChannel] or [PeakPowerForChannel].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/updateMeters()
func (a AVAudioRecorder) UpdateMeters() {
	objc.Send[objc.ID](a.ID, objc.Sel("updateMeters"))
}
// Returns the average power, in decibels full-scale (dBFS), for an audio
// channel.
//
// channelNumber: The number of the channel that you want the average power value for.
//
// # Return Value
// 
// The audio channel’s current average power.
//
// # Discussion
// 
// Before asking the player for its average power value, you must call
// [UpdateMeters] to generate the latest data. The returned value ranges from
// `–160` dBFS, indicating minimum power, to 0 dBFS, indicating maximum
// power.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/averagePower(forChannel:)
func (a AVAudioRecorder) AveragePowerForChannel(channelNumber uint) float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("averagePowerForChannel:"), channelNumber)
	return rv
}
// Returns the peak power, in decibels full-scale (dBFS), for an audio
// channel.
//
// channelNumber: The number of the channel that you want the peak power value for.
//
// # Return Value
// 
// The audio channel’s current peak power.
//
// # Discussion
// 
// Before asking the player for its peak power value, you must call
// [UpdateMeters] to generate the latest data. The returned value ranges from
// `–160` dBFS, indicating minimum power, to 0 dBFS, indicating maximum
// power.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/peakPower(forChannel:)
func (a AVAudioRecorder) PeakPowerForChannel(channelNumber uint) float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("peakPowerForChannel:"), channelNumber)
	return rv
}

// A Boolean value that indicates whether the audio recorder is recording.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/isRecording
func (a AVAudioRecorder) Recording() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isRecording"))
	return rv
}
// The time, in seconds, since the beginning of the recording.
//
// # Discussion
// 
// The value of this property is `0` when you call it on a stopped audio
// recorder.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/currentTime
func (a AVAudioRecorder) CurrentTime() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("currentTime"))
	return rv
}
// The time, in seconds, of the host audio device.
//
// # Discussion
// 
// Use this property value to schedule audio recording using the
// [RecordAtTime] and [RecordAtTimeForDuration] methods.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/deviceCurrentTime
func (a AVAudioRecorder) DeviceCurrentTime() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("deviceCurrentTime"))
	return rv
}
// A Boolean value that indicates whether you’ve enabled the recorder to
// generate audio-level metering data.
//
// # Discussion
// 
// By default, the recorder doesn’t generate audio-level metering data.
// Because metering uses computing resources, enable it only if you intend to
// use it.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/isMeteringEnabled
func (a AVAudioRecorder) MeteringEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isMeteringEnabled"))
	return rv
}
func (a AVAudioRecorder) SetMeteringEnabled(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setMeteringEnabled:"), value)
}
// The delegate object for the audio recorder.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/delegate
func (a AVAudioRecorder) Delegate() AVAudioRecorderDelegate {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("delegate"))
	return AVAudioRecorderDelegateObjectFromID(rv)
}
func (a AVAudioRecorder) SetDelegate(value AVAudioRecorderDelegate) {
	objc.Send[struct{}](a.ID, objc.Sel("setDelegate:"), value)
}
// The URL to which the recorder writes its data.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/url
func (a AVAudioRecorder) Url() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// The format of the recorded audio.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/format
func (a AVAudioRecorder) Format() IAVAudioFormat {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("format"))
	return AVAudioFormatFromID(objc.ID(rv))
}
// The settings that describe the format of the recorded audio.
//
// # Discussion
// 
// See [InitWithURLSettingsError] for supported keys and values.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRecorder/settings
func (a AVAudioRecorder) Settings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("settings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// The category for recording (input) and playback (output) of audio, such as
// for a Voice over Internet Protocol (VoIP) app.
//
// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/category-swift.struct/playandrecord
func (a AVAudioRecorder) PlayAndRecord() objc.ID {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVAudioSessionCategoryPlayAndRecord"))
	return rv
}
// The category for recording audio while also silencing playback audio.
//
// See: https://developer.apple.com/documentation/avfaudio/avaudiosession/category-swift.struct/record
func (a AVAudioRecorder) Record() objc.ID {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVAudioSessionCategoryRecord"))
	return rv
}

