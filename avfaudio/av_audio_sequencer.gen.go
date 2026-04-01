// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioSequencer] class.
var (
	_AVAudioSequencerClass     AVAudioSequencerClass
	_AVAudioSequencerClassOnce sync.Once
)

func getAVAudioSequencerClass() AVAudioSequencerClass {
	_AVAudioSequencerClassOnce.Do(func() {
		_AVAudioSequencerClass = AVAudioSequencerClass{class: objc.GetClass("AVAudioSequencer")}
	})
	return _AVAudioSequencerClass
}

// GetAVAudioSequencerClass returns the class object for AVAudioSequencer.
func GetAVAudioSequencerClass() AVAudioSequencerClass {
	return getAVAudioSequencerClass()
}

type AVAudioSequencerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioSequencerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioSequencerClass) Alloc() AVAudioSequencer {
	rv := objc.Send[AVAudioSequencer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that plays audio from a collection of MIDI events the system
// organizes into music tracks.
//
// # Creating an Audio Sequencer
//
//   - [AVAudioSequencer.InitWithAudioEngine]: Creates an audio sequencer that the framework attaches to an audio engine instance.
//
// # Writing to a MIDI File
//
//   - [AVAudioSequencer.WriteToURLSMPTEResolutionReplaceExistingError]: Creates and writes a MIDI file from the events in the sequence.
//
// # Handling Music Tracks
//
//   - [AVAudioSequencer.CreateAndAppendTrack]: Creates a new music track and appends it to the sequencer’s list.
//   - [AVAudioSequencer.ReverseEvents]: Reverses the order of all events in all music tracks, including the tempo track.
//   - [AVAudioSequencer.RemoveTrack]: Removes the music track from the sequencer.
//
// # Managing Sequence Load Options
//
//   - [AVAudioSequencer.LoadFromDataOptionsError]: Parses the data and adds its events to the sequence.
//   - [AVAudioSequencer.LoadFromURLOptionsError]: Loads the file the URL references and adds the events to the sequence.
//
// # Operating an Audio Sequencer
//
//   - [AVAudioSequencer.PrepareToPlay]: Gets ready to play the sequence by prerolling all events.
//   - [AVAudioSequencer.StartAndReturnError]: Starts the sequencer’s player.
//   - [AVAudioSequencer.Stop]: Stops the sequencer’s player.
//
// # Managing Time Stamps
//
//   - [AVAudioSequencer.HostTimeForBeatsError]: Gets the host time the sequence plays at the specified position.
//   - [AVAudioSequencer.SecondsForBeats]: Gets the time for the specified beat position (timestamp) in the track, in seconds.
//
// # Handling Beat Range
//
//   - [AVAudioSequencer.BeatsForHostTimeError]: Gets the beat the system plays at the specified host time.
//   - [AVAudioSequencer.BeatsForSeconds]: Gets the beat position (timestamp) for the specified time in the track.
//   - [AVAudioSequencer.AVMusicTimeStampEndOfTrack]: A timestamp you use to access all events in a music track through a beat range.
//   - [AVAudioSequencer.SetAVMusicTimeStampEndOfTrack]
//
// # Setting the User Callback
//
//   - [AVAudioSequencer.SetUserCallback]: Adds a callback that the sequencer calls each time it encounters a user event during playback.
//
// # Getting Sequence Properties
//
//   - [AVAudioSequencer.Playing]: A Boolean value that indicates whether the sequencer’s player is in a playing state.
//   - [AVAudioSequencer.Rate]: The playback rate of the sequencer’s player.
//   - [AVAudioSequencer.SetRate]
//   - [AVAudioSequencer.Tracks]: An array that contains all the tracks in the sequence.
//   - [AVAudioSequencer.CurrentPositionInBeats]: The current playback position, in beats.
//   - [AVAudioSequencer.SetCurrentPositionInBeats]
//   - [AVAudioSequencer.CurrentPositionInSeconds]: The current playback position, in seconds.
//   - [AVAudioSequencer.SetCurrentPositionInSeconds]
//   - [AVAudioSequencer.TempoTrack]: The track that contains tempo information about the sequence.
//   - [AVAudioSequencer.UserInfo]: A dictionary that contains metadata from a sequence.
//   - [AVAudioSequencer.DataWithSMPTEResolutionError]: Gets a data object that contains the events from the sequence.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer
type AVAudioSequencer struct {
	objectivec.Object
}

// AVAudioSequencerFromID constructs a [AVAudioSequencer] from an objc.ID.
//
// An object that plays audio from a collection of MIDI events the system
// organizes into music tracks.
func AVAudioSequencerFromID(id objc.ID) AVAudioSequencer {
	return AVAudioSequencer{objectivec.Object{ID: id}}
}

// NOTE: AVAudioSequencer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioSequencer] class.
//
// # Creating an Audio Sequencer
//
//   - [IAVAudioSequencer.InitWithAudioEngine]: Creates an audio sequencer that the framework attaches to an audio engine instance.
//
// # Writing to a MIDI File
//
//   - [IAVAudioSequencer.WriteToURLSMPTEResolutionReplaceExistingError]: Creates and writes a MIDI file from the events in the sequence.
//
// # Handling Music Tracks
//
//   - [IAVAudioSequencer.CreateAndAppendTrack]: Creates a new music track and appends it to the sequencer’s list.
//   - [IAVAudioSequencer.ReverseEvents]: Reverses the order of all events in all music tracks, including the tempo track.
//   - [IAVAudioSequencer.RemoveTrack]: Removes the music track from the sequencer.
//
// # Managing Sequence Load Options
//
//   - [IAVAudioSequencer.LoadFromDataOptionsError]: Parses the data and adds its events to the sequence.
//   - [IAVAudioSequencer.LoadFromURLOptionsError]: Loads the file the URL references and adds the events to the sequence.
//
// # Operating an Audio Sequencer
//
//   - [IAVAudioSequencer.PrepareToPlay]: Gets ready to play the sequence by prerolling all events.
//   - [IAVAudioSequencer.StartAndReturnError]: Starts the sequencer’s player.
//   - [IAVAudioSequencer.Stop]: Stops the sequencer’s player.
//
// # Managing Time Stamps
//
//   - [IAVAudioSequencer.HostTimeForBeatsError]: Gets the host time the sequence plays at the specified position.
//   - [IAVAudioSequencer.SecondsForBeats]: Gets the time for the specified beat position (timestamp) in the track, in seconds.
//
// # Handling Beat Range
//
//   - [IAVAudioSequencer.BeatsForHostTimeError]: Gets the beat the system plays at the specified host time.
//   - [IAVAudioSequencer.BeatsForSeconds]: Gets the beat position (timestamp) for the specified time in the track.
//   - [IAVAudioSequencer.AVMusicTimeStampEndOfTrack]: A timestamp you use to access all events in a music track through a beat range.
//   - [IAVAudioSequencer.SetAVMusicTimeStampEndOfTrack]
//
// # Setting the User Callback
//
//   - [IAVAudioSequencer.SetUserCallback]: Adds a callback that the sequencer calls each time it encounters a user event during playback.
//
// # Getting Sequence Properties
//
//   - [IAVAudioSequencer.Playing]: A Boolean value that indicates whether the sequencer’s player is in a playing state.
//   - [IAVAudioSequencer.Rate]: The playback rate of the sequencer’s player.
//   - [IAVAudioSequencer.SetRate]
//   - [IAVAudioSequencer.Tracks]: An array that contains all the tracks in the sequence.
//   - [IAVAudioSequencer.CurrentPositionInBeats]: The current playback position, in beats.
//   - [IAVAudioSequencer.SetCurrentPositionInBeats]
//   - [IAVAudioSequencer.CurrentPositionInSeconds]: The current playback position, in seconds.
//   - [IAVAudioSequencer.SetCurrentPositionInSeconds]
//   - [IAVAudioSequencer.TempoTrack]: The track that contains tempo information about the sequence.
//   - [IAVAudioSequencer.UserInfo]: A dictionary that contains metadata from a sequence.
//   - [IAVAudioSequencer.DataWithSMPTEResolutionError]: Gets a data object that contains the events from the sequence.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer
type IAVAudioSequencer interface {
	objectivec.IObject

	// Topic: Creating an Audio Sequencer

	// Creates an audio sequencer that the framework attaches to an audio engine instance.
	InitWithAudioEngine(engine IAVAudioEngine) AVAudioSequencer

	// Topic: Writing to a MIDI File

	// Creates and writes a MIDI file from the events in the sequence.
	WriteToURLSMPTEResolutionReplaceExistingError(fileURL foundation.INSURL, resolution int, replace bool) (bool, error)

	// Topic: Handling Music Tracks

	// Creates a new music track and appends it to the sequencer’s list.
	CreateAndAppendTrack() IAVMusicTrack
	// Reverses the order of all events in all music tracks, including the tempo track.
	ReverseEvents()
	// Removes the music track from the sequencer.
	RemoveTrack(track IAVMusicTrack) bool

	// Topic: Managing Sequence Load Options

	// Parses the data and adds its events to the sequence.
	LoadFromDataOptionsError(data foundation.INSData, options AVMusicSequenceLoadOptions) (bool, error)
	// Loads the file the URL references and adds the events to the sequence.
	LoadFromURLOptionsError(fileURL foundation.INSURL, options AVMusicSequenceLoadOptions) (bool, error)

	// Topic: Operating an Audio Sequencer

	// Gets ready to play the sequence by prerolling all events.
	PrepareToPlay()
	// Starts the sequencer’s player.
	StartAndReturnError() (bool, error)
	// Stops the sequencer’s player.
	Stop()

	// Topic: Managing Time Stamps

	// Gets the host time the sequence plays at the specified position.
	HostTimeForBeatsError(inBeats AVMusicTimeStamp) (uint64, error)
	// Gets the time for the specified beat position (timestamp) in the track, in seconds.
	SecondsForBeats(beats AVMusicTimeStamp) float64

	// Topic: Handling Beat Range

	// Gets the beat the system plays at the specified host time.
	BeatsForHostTimeError(inHostTime uint64) (AVMusicTimeStamp, error)
	// Gets the beat position (timestamp) for the specified time in the track.
	BeatsForSeconds(seconds float64) AVMusicTimeStamp
	// A timestamp you use to access all events in a music track through a beat range.
	AVMusicTimeStampEndOfTrack() float64
	SetAVMusicTimeStampEndOfTrack(value float64)

	// Topic: Setting the User Callback

	// Adds a callback that the sequencer calls each time it encounters a user event during playback.
	SetUserCallback(userCallback AVAudioSequencerUserCallback)

	// Topic: Getting Sequence Properties

	// A Boolean value that indicates whether the sequencer’s player is in a playing state.
	Playing() bool
	// The playback rate of the sequencer’s player.
	Rate() float32
	SetRate(value float32)
	// An array that contains all the tracks in the sequence.
	Tracks() []AVMusicTrack
	// The current playback position, in beats.
	CurrentPositionInBeats() float64
	SetCurrentPositionInBeats(value float64)
	// The current playback position, in seconds.
	CurrentPositionInSeconds() float64
	SetCurrentPositionInSeconds(value float64)
	// The track that contains tempo information about the sequence.
	TempoTrack() IAVMusicTrack
	// A dictionary that contains metadata from a sequence.
	UserInfo() foundation.INSDictionary
	// Gets a data object that contains the events from the sequence.
	DataWithSMPTEResolutionError(SMPTEResolution int) (foundation.INSData, error)
}

// Init initializes the instance.
func (a AVAudioSequencer) Init() AVAudioSequencer {
	rv := objc.Send[AVAudioSequencer](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioSequencer) Autorelease() AVAudioSequencer {
	rv := objc.Send[AVAudioSequencer](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioSequencer creates a new AVAudioSequencer instance.
func NewAVAudioSequencer() AVAudioSequencer {
	class := getAVAudioSequencerClass()
	rv := objc.Send[AVAudioSequencer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an audio sequencer that the framework attaches to an audio engine
// instance.
//
// engine: The engine to attach to.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/init(audioEngine:)
func NewAudioSequencerWithAudioEngine(engine IAVAudioEngine) AVAudioSequencer {
	instance := getAVAudioSequencerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAudioEngine:"), engine)
	return AVAudioSequencerFromID(rv)
}

// Creates an audio sequencer that the framework attaches to an audio engine
// instance.
//
// engine: The engine to attach to.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/init(audioEngine:)
func (a AVAudioSequencer) InitWithAudioEngine(engine IAVAudioEngine) AVAudioSequencer {
	rv := objc.Send[AVAudioSequencer](a.ID, objc.Sel("initWithAudioEngine:"), engine)
	return rv
}

// Creates and writes a MIDI file from the events in the sequence.
//
// fileURL: The URL of the file you want to write to.
//
// resolution: The relationship between tick and quarter note for saving to a Standard
// MIDI File. Passing zero uses the default value set using the tempo track.
//
// replace: When `true`, the framework overwrites an existing file at `fileURL`.
// Otherwise, the call fails with a permission error if a file at the
// specified path exists.
//
// # Discussion
//
// The framework writes only MIDI events when writing to the MIDI file. MIDI
// files are normally beat-based, but can also have an SMPTE (or real-time,
// rather than beat time) representation. The relationship between tick and
// quarter note for saving to a Standard MIDI File is the current value for
// the tempo track.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/write(to:smpteResolution:replaceExisting:)
func (a AVAudioSequencer) WriteToURLSMPTEResolutionReplaceExistingError(fileURL foundation.INSURL, resolution int, replace bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("writeToURL:SMPTEResolution:replaceExisting:error:"), fileURL, resolution, replace, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeToURL:SMPTEResolution:replaceExisting:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Creates a new music track and appends it to the sequencer’s list.
//
// # Return Value
//
// A new music track appended to the sequencer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/createAndAppendTrack()
func (a AVAudioSequencer) CreateAndAppendTrack() IAVMusicTrack {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("createAndAppendTrack"))
	return AVMusicTrackFromID(rv)
}

// Reverses the order of all events in all music tracks, including the tempo
// track.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/reverseEvents()
func (a AVAudioSequencer) ReverseEvents() {
	objc.Send[objc.ID](a.ID, objc.Sel("reverseEvents"))
}

// Removes the music track from the sequencer.
//
// track: The music track to remove.
//
// # Return Value
//
// A Boolean value that indicates whether the call succeeds.
//
// # Discussion
//
// This method doesn’t destroy the method track since you can reuse it.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/removeTrack(_:)
func (a AVAudioSequencer) RemoveTrack(track IAVMusicTrack) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("removeTrack:"), track)
	return rv
}

// Parses the data and adds its events to the sequence.
//
// data: The data to load from.
//
// options: Determines how the contents map to the tracks inside the sequence.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/load(from:options:)-8o58w
func (a AVAudioSequencer) LoadFromDataOptionsError(data foundation.INSData, options AVMusicSequenceLoadOptions) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("loadFromData:options:error:"), data, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("loadFromData:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Loads the file the URL references and adds the events to the sequence.
//
// fileURL: The URL to the file.
//
// options: Determines how the contents map to the tracks inside the sequence.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/load(from:options:)-9kb6m
func (a AVAudioSequencer) LoadFromURLOptionsError(fileURL foundation.INSURL, options AVMusicSequenceLoadOptions) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("loadFromURL:options:error:"), fileURL, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("loadFromURL:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Gets ready to play the sequence by prerolling all events.
//
// # Discussion
//
// The framework invokes this method automatically on play if you don’t call
// it, but it may delay startup.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/prepareToPlay()
func (a AVAudioSequencer) PrepareToPlay() {
	objc.Send[objc.ID](a.ID, objc.Sel("prepareToPlay"))
}

// Starts the sequencer’s player.
//
// # Discussion
//
// If you don’t call [PrepareToPlay], the framework calls it and then starts
// the player.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/start()
func (a AVAudioSequencer) StartAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("startAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// Stops the sequencer’s player.
//
// # Discussion
//
// Stopping the player leaves it in an unprerolled state, but stores the
// playback position so that a subsequent call to [StartAndReturnError]
// resumes where it stops. This action doesn’t stop an audio engine you
// associate with it.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/stop()
func (a AVAudioSequencer) Stop() {
	objc.Send[objc.ID](a.ID, objc.Sel("stop"))
}

// Gets the host time the sequence plays at the specified position.
//
// inBeats: The timestamp for the beat position.
//
// outError: On exit, if an error occurs, a description of the error.
//
// # Discussion
//
// This call is valid when the player is in a playing state. It returns `0`
// with an error, otherwise, or if the starting position of the player is
// after the specified beat. The method uses the sequence’s tempo map to
// translate a beat time from the starting time and the beat of the player.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/hostTime(forBeats:error:)
func (a AVAudioSequencer) HostTimeForBeatsError(inBeats AVMusicTimeStamp) (uint64, error) {
	var errorPtr objc.ID
	rv := objc.Send[uint64](a.ID, objc.Sel("hostTimeForBeats:error:"), inBeats, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// Gets the time for the specified beat position (timestamp) in the track, in
// seconds.
//
// beats: The timestamp for the beat position.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/seconds(forBeats:)
func (a AVAudioSequencer) SecondsForBeats(beats AVMusicTimeStamp) float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("secondsForBeats:"), beats)
	return rv
}

// Gets the beat the system plays at the specified host time.
//
// inHostTime: The host time for the beat position.
//
// outError: On exit, if an error occurs, a description of the error.
//
// # Discussion
//
// This call is valid when the player is in a playing state. It returns `0`
// with an error, otherwise, or if the starting position of the player is
// after the specified host time. This method uses the sequence’s tempo map
// to retrieve a beat time from the specified host time.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/beats(forHostTime:error:)
func (a AVAudioSequencer) BeatsForHostTimeError(inHostTime uint64) (AVMusicTimeStamp, error) {
	var errorPtr objc.ID
	rv := objc.Send[AVMusicTimeStamp](a.ID, objc.Sel("beatsForHostTime:error:"), inHostTime, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(AVMusicTimeStamp), foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// Gets the beat position (timestamp) for the specified time in the track.
//
// seconds: The time to retrieve the beat timestamp for.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/beats(forSeconds:)
func (a AVAudioSequencer) BeatsForSeconds(seconds float64) AVMusicTimeStamp {
	rv := objc.Send[AVMusicTimeStamp](a.ID, objc.Sel("beatsForSeconds:"), seconds)
	return AVMusicTimeStamp(rv)
}

// Adds a callback that the sequencer calls each time it encounters a user
// event during playback.
//
// userCallback: The user callback that the system calls.
//
// # Discussion
//
// The system calls the same callback for events that occur on any track in
// the sequencer. Set the callback to `nil` to disable it.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/setUserCallback(_:)
func (a AVAudioSequencer) SetUserCallback(userCallback AVAudioSequencerUserCallback) {
	objc.Send[objc.ID](a.ID, objc.Sel("setUserCallback:"), userCallback)
}

// Gets a data object that contains the events from the sequence.
//
// SMPTEResolution: The relationship between tick and quarter note for saving to a Standard
// MIDI File. Pass `0` to use the default.
//
// outError: On exit, if an error occurs, a description of the error.
//
// # Discussion
//
// The client controls the lifetime of the data value this method returns.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/data(withSMPTEResolution:error:)
func (a AVAudioSequencer) DataWithSMPTEResolutionError(SMPTEResolution int) (foundation.INSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("dataWithSMPTEResolution:error:"), SMPTEResolution, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSData{}, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDataFromID(rv), nil

}

// A Boolean value that indicates whether the sequencer’s player is in a
// playing state.
//
// # Discussion
//
// This value returns true if the sequencer’s player is in a started state.
// The framework considers it to be playing until it explicitly stops,
// including when playing past the end of the events in a sequence.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/isPlaying
func (a AVAudioSequencer) Playing() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isPlaying"))
	return rv
}

// The playback rate of the sequencer’s player.
//
// # Discussion
//
// The default playback rate is `1.0`, and must be greater than `0.0`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/rate
func (a AVAudioSequencer) Rate() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("rate"))
	return rv
}
func (a AVAudioSequencer) SetRate(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setRate:"), value)
}

// An array that contains all the tracks in the sequence.
//
// # Discussion
//
// The track indices start at `0`, and don’t include the tempo track.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/tracks
func (a AVAudioSequencer) Tracks() []AVMusicTrack {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("tracks"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMusicTrack {
		return AVMusicTrackFromID(id)
	})
}

// The current playback position, in beats.
//
// # Discussion
//
// Setting this property positions the sequencer’s player to the specified
// beat. You can update this property while the player is in a playing state,
// in which case, playback resumes at the new position.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/currentPositionInBeats
func (a AVAudioSequencer) CurrentPositionInBeats() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("currentPositionInBeats"))
	return rv
}
func (a AVAudioSequencer) SetCurrentPositionInBeats(value float64) {
	objc.Send[struct{}](a.ID, objc.Sel("setCurrentPositionInBeats:"), value)
}

// The current playback position, in seconds.
//
// # Discussion
//
// This property positions the sequencer’s player to the specified time. You
// can update this property while the player is in a playing state, in which
// case, playback resumes at the new position.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/currentPositionInSeconds
func (a AVAudioSequencer) CurrentPositionInSeconds() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("currentPositionInSeconds"))
	return rv
}
func (a AVAudioSequencer) SetCurrentPositionInSeconds(value float64) {
	objc.Send[struct{}](a.ID, objc.Sel("setCurrentPositionInSeconds:"), value)
}

// The track that contains tempo information about the sequence.
//
// # Discussion
//
// Each sequence has a single tempo track. The framework places all tempo
// events into this track along with other appropriate events, such as the
// time signature from a MIDI file.
//
// You can edit the tempo track like any other track. The framework ignores
// nontempo events in the track.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/tempoTrack
func (a AVAudioSequencer) TempoTrack() IAVMusicTrack {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("tempoTrack"))
	return AVMusicTrackFromID(objc.ID(rv))
}

// A dictionary that contains metadata from a sequence.
//
// # Discussion
//
// This property contains one or more of the values from
// [AVAudioSequencerInfoDictionaryKey].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/userInfo
func (a AVAudioSequencer) UserInfo() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("userInfo"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// A timestamp you use to access all events in a music track through a beat
// range.
//
// See: https://developer.apple.com/documentation/avfaudio/avmusictimestampendoftrack
func (a AVAudioSequencer) AVMusicTimeStampEndOfTrack() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("AVMusicTimeStampEndOfTrack"))
	return rv
}
func (a AVAudioSequencer) SetAVMusicTimeStampEndOfTrack(value float64) {
	objc.Send[struct{}](a.ID, objc.Sel("setAVMusicTimeStampEndOfTrack:"), value)
}
