// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioPlayer] class.
var (
	_AVAudioPlayerClass     AVAudioPlayerClass
	_AVAudioPlayerClassOnce sync.Once
)

func getAVAudioPlayerClass() AVAudioPlayerClass {
	_AVAudioPlayerClassOnce.Do(func() {
		_AVAudioPlayerClass = AVAudioPlayerClass{class: objc.GetClass("AVAudioPlayer")}
	})
	return _AVAudioPlayerClass
}

// GetAVAudioPlayerClass returns the class object for AVAudioPlayer.
func GetAVAudioPlayerClass() AVAudioPlayerClass {
	return getAVAudioPlayerClass()
}

type AVAudioPlayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioPlayerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioPlayerClass) Alloc() AVAudioPlayer {
	rv := objc.Send[AVAudioPlayer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that plays audio data from a file or buffer.
//
// # Overview
// 
// Use an audio player to:
// 
// - Play audio of any duration from a file or buffer - Control the volume,
// panning, rate, and looping behavior of the played audio - Access
// playback-level metering data - Play multiple sounds simultaneously by
// synchronizing the playback of multiple players
// 
// For more information about preparing your app to play audio, see
// [Configuring your app for media playback].
//
// [Configuring your app for media playback]: https://developer.apple.com/documentation/AVFoundation/configuring-your-app-for-media-playback
//
// # Creating an audio player
//
//   - [AVAudioPlayer.InitWithContentsOfURLError]: Creates a player to play audio from a file.
//   - [AVAudioPlayer.InitWithContentsOfURLFileTypeHintError]: Creates a player to play audio from a file of a particular type.
//   - [AVAudioPlayer.InitWithDataError]: Creates a player to play in-memory audio data.
//   - [AVAudioPlayer.InitWithDataFileTypeHintError]: Creates a player to play in-memory audio data of a particular type.
//
// # Controlling playback
//
//   - [AVAudioPlayer.PrepareToPlay]: Prepares the player for audio playback.
//   - [AVAudioPlayer.Play]: Plays audio asynchronously.
//   - [AVAudioPlayer.PlayAtTime]: Plays audio asynchronously, starting at a specified point in the audio output device’s timeline.
//   - [AVAudioPlayer.Pause]: Pauses audio playback.
//   - [AVAudioPlayer.Stop]: Stops playback and undoes the setup the system requires for playback.
//   - [AVAudioPlayer.Playing]: A Boolean value that indicates whether the player is currently playing audio.
//
// # Configuring playback settings
//
//   - [AVAudioPlayer.Volume]: The audio player’s volume relative to other audio output.
//   - [AVAudioPlayer.SetVolume]
//   - [AVAudioPlayer.SetVolumeFadeDuration]: Changes the audio player’s volume over a duration of time.
//   - [AVAudioPlayer.Pan]: The audio player’s stereo pan position.
//   - [AVAudioPlayer.SetPan]
//   - [AVAudioPlayer.EnableRate]: A Boolean value that indicates whether you can adjust the playback rate of the audio player.
//   - [AVAudioPlayer.SetEnableRate]
//   - [AVAudioPlayer.Rate]: The audio player’s playback rate.
//   - [AVAudioPlayer.SetRate]
//   - [AVAudioPlayer.NumberOfLoops]: The number of times the audio repeats playback.
//   - [AVAudioPlayer.SetNumberOfLoops]
//
// # Accessing player timing
//
//   - [AVAudioPlayer.CurrentTime]: The current playback time, in seconds, within the audio timeline.
//   - [AVAudioPlayer.SetCurrentTime]
//   - [AVAudioPlayer.Duration]: The total duration, in seconds, of the player’s audio.
//
// # Configuring the Spatial Audio experience
//
//   - [AVAudioPlayer.IntendedSpatialExperience]: The intended spatial experience for this player.
//   - [AVAudioPlayer.SetIntendedSpatialExperience]
//
// # Managing audio channels
//
//   - [AVAudioPlayer.NumberOfChannels]: The number of audio channels in the player’s audio.
//
// # Managing audio-level metering
//
//   - [AVAudioPlayer.MeteringEnabled]: A Boolean value that indicates whether the player is able to generate audio-level metering data.
//   - [AVAudioPlayer.SetMeteringEnabled]
//   - [AVAudioPlayer.UpdateMeters]: Refreshes the average and peak power values for all channels of an audio player.
//   - [AVAudioPlayer.AveragePowerForChannel]: Returns the average power, in decibels full-scale (dBFS), for an audio channel.
//   - [AVAudioPlayer.PeakPowerForChannel]: Returns the peak power, in decibels full-scale (dBFS), for an audio channel.
//
// # Responding to player events
//
//   - [AVAudioPlayer.Delegate]: The delegate object for the audio player.
//   - [AVAudioPlayer.SetDelegate]
//
// # Inspecting the audio data
//
//   - [AVAudioPlayer.Url]: The URL of the audio file.
//   - [AVAudioPlayer.Data]: The audio data associated with the player.
//   - [AVAudioPlayer.Format]: The format of the player’s audio data.
//   - [AVAudioPlayer.Settings]: A dictionary that provides information about the player’s audio data.
//
// # Accessing device information
//
//   - [AVAudioPlayer.CurrentDevice]: The unique identifier of the current audio player.
//   - [AVAudioPlayer.SetCurrentDevice]
//   - [AVAudioPlayer.DeviceCurrentTime]: The time value, in seconds, of the audio output device’s clock.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer
type AVAudioPlayer struct {
	objectivec.Object
}

// AVAudioPlayerFromID constructs a [AVAudioPlayer] from an objc.ID.
//
// An object that plays audio data from a file or buffer.
func AVAudioPlayerFromID(id objc.ID) AVAudioPlayer {
	return AVAudioPlayer{objectivec.Object{ID: id}}
}
// NOTE: AVAudioPlayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioPlayer] class.
//
// # Creating an audio player
//
//   - [IAVAudioPlayer.InitWithContentsOfURLError]: Creates a player to play audio from a file.
//   - [IAVAudioPlayer.InitWithContentsOfURLFileTypeHintError]: Creates a player to play audio from a file of a particular type.
//   - [IAVAudioPlayer.InitWithDataError]: Creates a player to play in-memory audio data.
//   - [IAVAudioPlayer.InitWithDataFileTypeHintError]: Creates a player to play in-memory audio data of a particular type.
//
// # Controlling playback
//
//   - [IAVAudioPlayer.PrepareToPlay]: Prepares the player for audio playback.
//   - [IAVAudioPlayer.Play]: Plays audio asynchronously.
//   - [IAVAudioPlayer.PlayAtTime]: Plays audio asynchronously, starting at a specified point in the audio output device’s timeline.
//   - [IAVAudioPlayer.Pause]: Pauses audio playback.
//   - [IAVAudioPlayer.Stop]: Stops playback and undoes the setup the system requires for playback.
//   - [IAVAudioPlayer.Playing]: A Boolean value that indicates whether the player is currently playing audio.
//
// # Configuring playback settings
//
//   - [IAVAudioPlayer.Volume]: The audio player’s volume relative to other audio output.
//   - [IAVAudioPlayer.SetVolume]
//   - [IAVAudioPlayer.SetVolumeFadeDuration]: Changes the audio player’s volume over a duration of time.
//   - [IAVAudioPlayer.Pan]: The audio player’s stereo pan position.
//   - [IAVAudioPlayer.SetPan]
//   - [IAVAudioPlayer.EnableRate]: A Boolean value that indicates whether you can adjust the playback rate of the audio player.
//   - [IAVAudioPlayer.SetEnableRate]
//   - [IAVAudioPlayer.Rate]: The audio player’s playback rate.
//   - [IAVAudioPlayer.SetRate]
//   - [IAVAudioPlayer.NumberOfLoops]: The number of times the audio repeats playback.
//   - [IAVAudioPlayer.SetNumberOfLoops]
//
// # Accessing player timing
//
//   - [IAVAudioPlayer.CurrentTime]: The current playback time, in seconds, within the audio timeline.
//   - [IAVAudioPlayer.SetCurrentTime]
//   - [IAVAudioPlayer.Duration]: The total duration, in seconds, of the player’s audio.
//
// # Configuring the Spatial Audio experience
//
//   - [IAVAudioPlayer.IntendedSpatialExperience]: The intended spatial experience for this player.
//   - [IAVAudioPlayer.SetIntendedSpatialExperience]
//
// # Managing audio channels
//
//   - [IAVAudioPlayer.NumberOfChannels]: The number of audio channels in the player’s audio.
//
// # Managing audio-level metering
//
//   - [IAVAudioPlayer.MeteringEnabled]: A Boolean value that indicates whether the player is able to generate audio-level metering data.
//   - [IAVAudioPlayer.SetMeteringEnabled]
//   - [IAVAudioPlayer.UpdateMeters]: Refreshes the average and peak power values for all channels of an audio player.
//   - [IAVAudioPlayer.AveragePowerForChannel]: Returns the average power, in decibels full-scale (dBFS), for an audio channel.
//   - [IAVAudioPlayer.PeakPowerForChannel]: Returns the peak power, in decibels full-scale (dBFS), for an audio channel.
//
// # Responding to player events
//
//   - [IAVAudioPlayer.Delegate]: The delegate object for the audio player.
//   - [IAVAudioPlayer.SetDelegate]
//
// # Inspecting the audio data
//
//   - [IAVAudioPlayer.Url]: The URL of the audio file.
//   - [IAVAudioPlayer.Data]: The audio data associated with the player.
//   - [IAVAudioPlayer.Format]: The format of the player’s audio data.
//   - [IAVAudioPlayer.Settings]: A dictionary that provides information about the player’s audio data.
//
// # Accessing device information
//
//   - [IAVAudioPlayer.CurrentDevice]: The unique identifier of the current audio player.
//   - [IAVAudioPlayer.SetCurrentDevice]
//   - [IAVAudioPlayer.DeviceCurrentTime]: The time value, in seconds, of the audio output device’s clock.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer
type IAVAudioPlayer interface {
	objectivec.IObject

	// Topic: Creating an audio player

	// Creates a player to play audio from a file.
	InitWithContentsOfURLError(url foundation.INSURL) (AVAudioPlayer, error)
	// Creates a player to play audio from a file of a particular type.
	InitWithContentsOfURLFileTypeHintError(url foundation.INSURL, utiString string) (AVAudioPlayer, error)
	// Creates a player to play in-memory audio data.
	InitWithDataError(data foundation.INSData) (AVAudioPlayer, error)
	// Creates a player to play in-memory audio data of a particular type.
	InitWithDataFileTypeHintError(data foundation.INSData, utiString string) (AVAudioPlayer, error)

	// Topic: Controlling playback

	// Prepares the player for audio playback.
	PrepareToPlay() bool
	// Plays audio asynchronously.
	Play() bool
	// Plays audio asynchronously, starting at a specified point in the audio output device’s timeline.
	PlayAtTime(time float64) bool
	// Pauses audio playback.
	Pause()
	// Stops playback and undoes the setup the system requires for playback.
	Stop()
	// A Boolean value that indicates whether the player is currently playing audio.
	Playing() bool

	// Topic: Configuring playback settings

	// The audio player’s volume relative to other audio output.
	Volume() float32
	SetVolume(value float32)
	// Changes the audio player’s volume over a duration of time.
	SetVolumeFadeDuration(volume float32, duration float64)
	// The audio player’s stereo pan position.
	Pan() float32
	SetPan(value float32)
	// A Boolean value that indicates whether you can adjust the playback rate of the audio player.
	EnableRate() bool
	SetEnableRate(value bool)
	// The audio player’s playback rate.
	Rate() float32
	SetRate(value float32)
	// The number of times the audio repeats playback.
	NumberOfLoops() int
	SetNumberOfLoops(value int)

	// Topic: Accessing player timing

	// The current playback time, in seconds, within the audio timeline.
	CurrentTime() float64
	SetCurrentTime(value float64)
	// The total duration, in seconds, of the player’s audio.
	Duration() float64

	// Topic: Configuring the Spatial Audio experience

	// The intended spatial experience for this player.
	IntendedSpatialExperience() objectivec.IObject
	SetIntendedSpatialExperience(value objectivec.IObject)

	// Topic: Managing audio channels

	// The number of audio channels in the player’s audio.
	NumberOfChannels() uint

	// Topic: Managing audio-level metering

	// A Boolean value that indicates whether the player is able to generate audio-level metering data.
	MeteringEnabled() bool
	SetMeteringEnabled(value bool)
	// Refreshes the average and peak power values for all channels of an audio player.
	UpdateMeters()
	// Returns the average power, in decibels full-scale (dBFS), for an audio channel.
	AveragePowerForChannel(channelNumber uint) float32
	// Returns the peak power, in decibels full-scale (dBFS), for an audio channel.
	PeakPowerForChannel(channelNumber uint) float32

	// Topic: Responding to player events

	// The delegate object for the audio player.
	Delegate() AVAudioPlayerDelegate
	SetDelegate(value AVAudioPlayerDelegate)

	// Topic: Inspecting the audio data

	// The URL of the audio file.
	Url() foundation.INSURL
	// The audio data associated with the player.
	Data() foundation.INSData
	// The format of the player’s audio data.
	Format() IAVAudioFormat
	// A dictionary that provides information about the player’s audio data.
	Settings() foundation.INSDictionary

	// Topic: Accessing device information

	// The unique identifier of the current audio player.
	CurrentDevice() string
	SetCurrentDevice(value string)
	// The time value, in seconds, of the audio output device’s clock.
	DeviceCurrentTime() float64
}

// Init initializes the instance.
func (a AVAudioPlayer) Init() AVAudioPlayer {
	rv := objc.Send[AVAudioPlayer](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioPlayer) Autorelease() AVAudioPlayer {
	rv := objc.Send[AVAudioPlayer](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioPlayer creates a new AVAudioPlayer instance.
func NewAVAudioPlayer() AVAudioPlayer {
	class := getAVAudioPlayerClass()
	rv := objc.Send[AVAudioPlayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a player to play audio from a file.
//
// url: A URL that identifies the local audio file to play.
//
// # Return Value
// 
// A new audio player instance, or [nil] if an error occurs.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// # Discussion
// 
// The audio data must be in a format that Core Audio supports.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/init(contentsOf:)
func NewAudioPlayerWithContentsOfURLError(url foundation.INSURL) (AVAudioPlayer, error) {
	var errorPtr objc.ID
	instance := getAVAudioPlayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioPlayer{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioPlayerFromID(rv), nil
}

// Creates a player to play audio from a file of a particular type.
//
// url: A URL that identifies the local audio file to play.
//
// utiString: The uniform type identifier (UTI) string of the file format.
//
// # Return Value
// 
// A new audio player instance, or [nil] if there is an error.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// # Discussion
// 
// The audio data must be in a format that Core Audio supports. Passing a file
// type hint helps the system parse the data if it can’t determine the file
// type or if the data is corrupt. See [AVFileType] for supported values.
//
// [AVFileType]: https://developer.apple.com/documentation/AVFoundation/AVFileType
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/init(contentsOf:fileTypeHint:)
func NewAudioPlayerWithContentsOfURLFileTypeHintError(url foundation.INSURL, utiString string) (AVAudioPlayer, error) {
	var errorPtr objc.ID
	instance := getAVAudioPlayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:fileTypeHint:error:"), url, objc.String(utiString), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioPlayer{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioPlayerFromID(rv), nil
}

// Creates a player to play in-memory audio data.
//
// data: A buffer with the audio data to play.
//
// # Return Value
// 
// A new audio player instance, or [nil] if an error occurs.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// # Discussion
// 
// The audio data must be in a format that Core Audio supports.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/init(data:)
func NewAudioPlayerWithDataError(data foundation.INSData) (AVAudioPlayer, error) {
	var errorPtr objc.ID
	instance := getAVAudioPlayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioPlayer{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioPlayerFromID(rv), nil
}

// Creates a player to play in-memory audio data of a particular type.
//
// data: A buffer with the audio data to play.
//
// utiString: The uniform type identifier (UTI) string of the file format.
//
// # Return Value
// 
// A new audio player instance, or [nil] if an error occurs.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// # Discussion
// 
// The audio data must be in a format that Core Audio supports. Passing a file
// type hint helps the system parse the data if it can’t determine the file
// type or if the data is corrupt. See [AVFileType] for supported values.
//
// [AVFileType]: https://developer.apple.com/documentation/AVFoundation/AVFileType
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/init(data:fileTypeHint:)
func NewAudioPlayerWithDataFileTypeHintError(data foundation.INSData, utiString string) (AVAudioPlayer, error) {
	var errorPtr objc.ID
	instance := getAVAudioPlayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:fileTypeHint:error:"), data, objc.String(utiString), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioPlayer{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioPlayerFromID(rv), nil
}

// Creates a player to play audio from a file.
//
// url: A URL that identifies the local audio file to play.
//
// # Return Value
// 
// A new audio player instance, or [nil] if an error occurs.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// # Discussion
// 
// The audio data must be in a format that Core Audio supports.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/init(contentsOf:)
func (a AVAudioPlayer) InitWithContentsOfURLError(url foundation.INSURL) (AVAudioPlayer, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioPlayer{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioPlayerFromID(rv), nil

}
// Creates a player to play audio from a file of a particular type.
//
// url: A URL that identifies the local audio file to play.
//
// utiString: The uniform type identifier (UTI) string of the file format.
//
// # Return Value
// 
// A new audio player instance, or [nil] if there is an error.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// # Discussion
// 
// The audio data must be in a format that Core Audio supports. Passing a file
// type hint helps the system parse the data if it can’t determine the file
// type or if the data is corrupt. See [AVFileType] for supported values.
//
// [AVFileType]: https://developer.apple.com/documentation/AVFoundation/AVFileType
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/init(contentsOf:fileTypeHint:)
func (a AVAudioPlayer) InitWithContentsOfURLFileTypeHintError(url foundation.INSURL, utiString string) (AVAudioPlayer, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithContentsOfURL:fileTypeHint:error:"), url, objc.String(utiString), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioPlayer{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioPlayerFromID(rv), nil

}
// Creates a player to play in-memory audio data.
//
// data: A buffer with the audio data to play.
//
// # Return Value
// 
// A new audio player instance, or [nil] if an error occurs.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// # Discussion
// 
// The audio data must be in a format that Core Audio supports.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/init(data:)
func (a AVAudioPlayer) InitWithDataError(data foundation.INSData) (AVAudioPlayer, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioPlayer{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioPlayerFromID(rv), nil

}
// Creates a player to play in-memory audio data of a particular type.
//
// data: A buffer with the audio data to play.
//
// utiString: The uniform type identifier (UTI) string of the file format.
//
// # Return Value
// 
// A new audio player instance, or [nil] if an error occurs.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// # Discussion
// 
// The audio data must be in a format that Core Audio supports. Passing a file
// type hint helps the system parse the data if it can’t determine the file
// type or if the data is corrupt. See [AVFileType] for supported values.
//
// [AVFileType]: https://developer.apple.com/documentation/AVFoundation/AVFileType
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/init(data:fileTypeHint:)
func (a AVAudioPlayer) InitWithDataFileTypeHintError(data foundation.INSData, utiString string) (AVAudioPlayer, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithData:fileTypeHint:error:"), data, objc.String(utiString), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVAudioPlayer{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVAudioPlayerFromID(rv), nil

}
// Prepares the player for audio playback.
//
// # Return Value
// 
// [true] if the system successfully prepares the player; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Calling this method preloads audio buffers and acquires the audio hardware
// necessary for playback. This method activates the audio session, so pass
// [false] to [setActive:error:] if immediate playback isn’t necessary. For
// example, when using the category option
// [AudioSessionCategoryOptionDuckOthers], this method lowers the audio
// outside of the app.
// 
// The system calls this method when using the method [Play], but calling it
// in advance minimizes the delay between calling `play()` and the start of
// sound output.
// 
// Calling [Stop], or allowing a sound to finish playing, undoes this setup.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [setActive:error:]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession/setActive:error:
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/prepareToPlay()
func (a AVAudioPlayer) PrepareToPlay() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("prepareToPlay"))
	return rv
}
// Plays audio asynchronously.
//
// # Return Value
// 
// [true] if playback starts successfully; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Calling this method implicitly calls [PrepareToPlay] if the audio player is
// unprepared for playback.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/play()
func (a AVAudioPlayer) Play() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("play"))
	return rv
}
// Plays audio asynchronously, starting at a specified point in the audio
// output device’s timeline.
//
// time: The audio device time to begin playback. This time must be later than the
// device’s current time.
//
// # Return Value
// 
// [true] if playback starts successfully; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method to precisely synchronize the playback of two or more audio
// player objects as the following example shows:
// 
// Calling this method implicitly calls [PrepareToPlay] if the audio player is
// unprepared for playback.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/play(atTime:)
func (a AVAudioPlayer) PlayAtTime(time float64) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("playAtTime:"), time)
	return rv
}
// Pauses audio playback.
//
// # Discussion
// 
// Unlike calling [Stop], pausing playback doesn’t deallocate hardware
// resources. It leaves the audio ready to resume playback from where it
// stops.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/pause()
func (a AVAudioPlayer) Pause() {
	objc.Send[objc.ID](a.ID, objc.Sel("pause"))
}
// Stops playback and undoes the setup the system requires for playback.
//
// # Discussion
// 
// Calling this method undoes the resource allocation the system performs in
// [PrepareToPlay] or [Play]. It doesn’t reset the player’s [CurrentTime]
// value to `0`, so playback resumes from where it stops.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/stop()
func (a AVAudioPlayer) Stop() {
	objc.Send[objc.ID](a.ID, objc.Sel("stop"))
}
// Changes the audio player’s volume over a duration of time.
//
// volume: The target volume.
//
// duration: The duration, in seconds, over which to fade the volume.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/setVolume(_:fadeDuration:)
func (a AVAudioPlayer) SetVolumeFadeDuration(volume float32, duration float64) {
	objc.Send[objc.ID](a.ID, objc.Sel("setVolume:fadeDuration:"), volume, duration)
}
// Refreshes the average and peak power values for all channels of an audio
// player.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/updateMeters()
func (a AVAudioPlayer) UpdateMeters() {
	objc.Send[objc.ID](a.ID, objc.Sel("updateMeters"))
}
// Returns the average power, in decibels full-scale (dBFS), for an audio
// channel.
//
// channelNumber: The audio channel with the average power value you want to retrieve.
// Channel numbers are zero-indexed. A monaural signal, or the left channel of
// a stereo signal, has channel number `0`.
//
// # Return Value
// 
// A floating-point value, in dBFS, that indicates the audio channel’s
// current average power.
//
// # Discussion
// 
// Before asking the player for its average power value, you must call
// [UpdateMeters] to generate the latest data. The returned value ranges from
// `–160` dBFS, indicating minimum power, to 0 dBFS, indicating maximum
// power.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/averagePower(forChannel:)
func (a AVAudioPlayer) AveragePowerForChannel(channelNumber uint) float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("averagePowerForChannel:"), channelNumber)
	return rv
}
// Returns the peak power, in decibels full-scale (dBFS), for an audio
// channel.
//
// channelNumber: The audio channel with the peak power value you want to obtain. Channel
// numbers are zero-indexed. A monaural signal, or the left channel of a
// stereo signal, has channel number `0`.
//
// # Return Value
// 
// A floating-point value, in dBFS, that indicates the audio channel’s
// current peak power.
//
// # Discussion
// 
// Before asking the player for its peak power value, you must call
// [UpdateMeters] to generate the latest data. The returned value ranges from
// `–160` dBFS, indicating minimum power, to 0 dBFS, indicating maximum
// power.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/peakPower(forChannel:)
func (a AVAudioPlayer) PeakPowerForChannel(channelNumber uint) float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("peakPowerForChannel:"), channelNumber)
	return rv
}

// A Boolean value that indicates whether the player is currently playing
// audio.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/isPlaying
func (a AVAudioPlayer) Playing() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isPlaying"))
	return rv
}
// The audio player’s volume relative to other audio output.
//
// # Discussion
// 
// This property supports values ranging from `0.0` for silence to `1.0` for
// full volume.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/volume
func (a AVAudioPlayer) Volume() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("volume"))
	return rv
}
func (a AVAudioPlayer) SetVolume(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setVolume:"), value)
}
// The audio player’s stereo pan position.
//
// # Discussion
// 
// Set this property value to position the audio in the stereo field. Use a
// value of `-1.0` to indicate full left, `1.0` for full right, and `0.0` for
// center.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/pan
func (a AVAudioPlayer) Pan() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("pan"))
	return rv
}
func (a AVAudioPlayer) SetPan(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setPan:"), value)
}
// A Boolean value that indicates whether you can adjust the playback rate of
// the audio player.
//
// # Discussion
// 
// To enable modifying the player’s rate, set this property to [true] after
// you create the player, but before you call [PrepareToPlay].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/enableRate
func (a AVAudioPlayer) EnableRate() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("enableRate"))
	return rv
}
func (a AVAudioPlayer) SetEnableRate(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setEnableRate:"), value)
}
// The audio player’s playback rate.
//
// # Discussion
// 
// To set an audio player’s playback rate, you must first enable the rate
// adjustment by setting its [EnableRate] property to [true].
// 
// The default value of this property is `1.0`, which indicates that audio
// playback occurs at standard speed. This property supports values in the
// range of `0.5` for half-speed playback to `2.0` for double-speed playback.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/rate
func (a AVAudioPlayer) Rate() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("rate"))
	return rv
}
func (a AVAudioPlayer) SetRate(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setRate:"), value)
}
// The number of times the audio repeats playback.
//
// # Discussion
// 
// The default value of `0` results in the sound playing once. Set a positive
// integer value to specify the number of times to repeat the sound. For
// example, a value of `1` plays the sound twice: the original sound and one
// repetition. Set a negative integer value to loop the sound continuously
// until you call the [Stop] method.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/numberOfLoops
func (a AVAudioPlayer) NumberOfLoops() int {
	rv := objc.Send[int](a.ID, objc.Sel("numberOfLoops"))
	return rv
}
func (a AVAudioPlayer) SetNumberOfLoops(value int) {
	objc.Send[struct{}](a.ID, objc.Sel("setNumberOfLoops:"), value)
}
// The current playback time, in seconds, within the audio timeline.
//
// # Discussion
// 
// If the sound is playing, this property value is the offset, in seconds,
// from the start of the sound. If the sound isn’t playing, this property
// indicates the offset from where playback starts upon calling the [Play]
// method.
// 
// Use this property to seek to a specific time in the audio data or to
// implement audio fast-forward and rewind functions.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/currentTime
func (a AVAudioPlayer) CurrentTime() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("currentTime"))
	return rv
}
func (a AVAudioPlayer) SetCurrentTime(value float64) {
	objc.Send[struct{}](a.ID, objc.Sel("setCurrentTime:"), value)
}
// The total duration, in seconds, of the player’s audio.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/duration
func (a AVAudioPlayer) Duration() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("duration"))
	return rv
}
// The intended spatial experience for this player.
//
// See: https://developer.apple.com/documentation/avfaudio/avaudioplayer/intendedspatialexperience-27klj
func (a AVAudioPlayer) IntendedSpatialExperience() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("intendedSpatialExperience"))
	return objectivec.Object{ID: rv}
}
func (a AVAudioPlayer) SetIntendedSpatialExperience(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setIntendedSpatialExperience:"), value)
}
// The number of audio channels in the player’s audio.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/numberOfChannels
func (a AVAudioPlayer) NumberOfChannels() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("numberOfChannels"))
	return rv
}
// A Boolean value that indicates whether the player is able to generate
// audio-level metering data.
//
// # Discussion
// 
// By default, the player doesn’t generate audio-level metering data.
// Because metering uses computing resources, enable it only if you intend to
// use it.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/isMeteringEnabled
func (a AVAudioPlayer) MeteringEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isMeteringEnabled"))
	return rv
}
func (a AVAudioPlayer) SetMeteringEnabled(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setMeteringEnabled:"), value)
}
// The delegate object for the audio player.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/delegate
func (a AVAudioPlayer) Delegate() AVAudioPlayerDelegate {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("delegate"))
	return AVAudioPlayerDelegateObjectFromID(rv)
}
func (a AVAudioPlayer) SetDelegate(value AVAudioPlayerDelegate) {
	objc.Send[struct{}](a.ID, objc.Sel("setDelegate:"), value)
}
// The URL of the audio file.
//
// # Discussion
// 
// This property is [nil] if you don’t create the player with a URL.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/url
func (a AVAudioPlayer) Url() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// The audio data associated with the player.
//
// # Discussion
// 
// This property is [nil] if you don’t create the player with a data buffer.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/data
func (a AVAudioPlayer) Data() foundation.INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("data"))
	return foundation.NSDataFromID(objc.ID(rv))
}
// The format of the player’s audio data.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/format
func (a AVAudioPlayer) Format() IAVAudioFormat {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("format"))
	return AVAudioFormatFromID(objc.ID(rv))
}
// A dictionary that provides information about the player’s audio data.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/settings
func (a AVAudioPlayer) Settings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("settings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// The unique identifier of the current audio player.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/currentDevice
func (a AVAudioPlayer) CurrentDevice() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("currentDevice"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVAudioPlayer) SetCurrentDevice(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setCurrentDevice:"), objc.String(value))
}
// The time value, in seconds, of the audio output device’s clock.
//
// # Discussion
// 
// The value of this property increases monotonically while an audio player is
// playing or is in a paused state. If you connect more than one audio player
// to the audio output device, the time continues incrementing while at least
// one of the players is playing or is in a paused state. If the audio output
// device has no connected audio players that are either playing or are in a
// paused state, device time reverts to `0.0`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPlayer/deviceCurrentTime
func (a AVAudioPlayer) DeviceCurrentTime() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("deviceCurrentTime"))
	return rv
}

