// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMIDIPlayer] class.
var (
	_AVMIDIPlayerClass     AVMIDIPlayerClass
	_AVMIDIPlayerClassOnce sync.Once
)

func getAVMIDIPlayerClass() AVMIDIPlayerClass {
	_AVMIDIPlayerClassOnce.Do(func() {
		_AVMIDIPlayerClass = AVMIDIPlayerClass{class: objc.GetClass("AVMIDIPlayer")}
	})
	return _AVMIDIPlayerClass
}

// GetAVMIDIPlayerClass returns the class object for AVMIDIPlayer.
func GetAVMIDIPlayerClass() AVMIDIPlayerClass {
	return getAVMIDIPlayerClass()
}

type AVMIDIPlayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMIDIPlayerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMIDIPlayerClass) Alloc() AVMIDIPlayer {
	rv := objc.Send[AVMIDIPlayer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that plays MIDI data through a system sound module.
//
// # Overview
// 
// For more information about preparing your app to play audio, see
// [Configuring your app for media playback].
//
// [Configuring your app for media playback]: https://developer.apple.com/documentation/AVFoundation/configuring-your-app-for-media-playback
//
// # Creating a MIDI player
//
//   - [AVMIDIPlayer.InitWithContentsOfURLSoundBankURLError]: Creates a player to play a MIDI file with the specified soundbank.
//   - [AVMIDIPlayer.InitWithDataSoundBankURLError]: Creates a player to play MIDI data with the specified soundbank.
//
// # Controlling playback
//
//   - [AVMIDIPlayer.PrepareToPlay]: Prepares the player to play the sequence by prerolling all events.
//   - [AVMIDIPlayer.Play]: Plays the MIDI sequence.
//   - [AVMIDIPlayer.Stop]: Stops playing the sequence.
//   - [AVMIDIPlayer.Playing]: A Boolean value that indicates whether the sequence is playing.
//
// # Configuring playback settings
//
//   - [AVMIDIPlayer.Rate]: The playback rate of the player.
//   - [AVMIDIPlayer.SetRate]
//
// # Accessing player timing
//
//   - [AVMIDIPlayer.CurrentPosition]: The current playback position, in seconds.
//   - [AVMIDIPlayer.SetCurrentPosition]
//   - [AVMIDIPlayer.Duration]: The duration, in seconds, of the currently loaded file.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer
type AVMIDIPlayer struct {
	objectivec.Object
}

// AVMIDIPlayerFromID constructs a [AVMIDIPlayer] from an objc.ID.
//
// An object that plays MIDI data through a system sound module.
func AVMIDIPlayerFromID(id objc.ID) AVMIDIPlayer {
	return AVMIDIPlayer{objectivec.Object{ID: id}}
}
// NOTE: AVMIDIPlayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMIDIPlayer] class.
//
// # Creating a MIDI player
//
//   - [IAVMIDIPlayer.InitWithContentsOfURLSoundBankURLError]: Creates a player to play a MIDI file with the specified soundbank.
//   - [IAVMIDIPlayer.InitWithDataSoundBankURLError]: Creates a player to play MIDI data with the specified soundbank.
//
// # Controlling playback
//
//   - [IAVMIDIPlayer.PrepareToPlay]: Prepares the player to play the sequence by prerolling all events.
//   - [IAVMIDIPlayer.Play]: Plays the MIDI sequence.
//   - [IAVMIDIPlayer.Stop]: Stops playing the sequence.
//   - [IAVMIDIPlayer.Playing]: A Boolean value that indicates whether the sequence is playing.
//
// # Configuring playback settings
//
//   - [IAVMIDIPlayer.Rate]: The playback rate of the player.
//   - [IAVMIDIPlayer.SetRate]
//
// # Accessing player timing
//
//   - [IAVMIDIPlayer.CurrentPosition]: The current playback position, in seconds.
//   - [IAVMIDIPlayer.SetCurrentPosition]
//   - [IAVMIDIPlayer.Duration]: The duration, in seconds, of the currently loaded file.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer
type IAVMIDIPlayer interface {
	objectivec.IObject

	// Topic: Creating a MIDI player

	// Creates a player to play a MIDI file with the specified soundbank.
	InitWithContentsOfURLSoundBankURLError(inURL foundation.INSURL, bankURL foundation.INSURL) (AVMIDIPlayer, error)
	// Creates a player to play MIDI data with the specified soundbank.
	InitWithDataSoundBankURLError(data foundation.INSData, bankURL foundation.INSURL) (AVMIDIPlayer, error)

	// Topic: Controlling playback

	// Prepares the player to play the sequence by prerolling all events.
	PrepareToPlay()
	// Plays the MIDI sequence.
	Play(completionHandler ErrorHandler)
	// Stops playing the sequence.
	Stop()
	// A Boolean value that indicates whether the sequence is playing.
	Playing() bool

	// Topic: Configuring playback settings

	// The playback rate of the player.
	Rate() float32
	SetRate(value float32)

	// Topic: Accessing player timing

	// The current playback position, in seconds.
	CurrentPosition() float64
	SetCurrentPosition(value float64)
	// The duration, in seconds, of the currently loaded file.
	Duration() float64
}

// Init initializes the instance.
func (m AVMIDIPlayer) Init() AVMIDIPlayer {
	rv := objc.Send[AVMIDIPlayer](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMIDIPlayer) Autorelease() AVMIDIPlayer {
	rv := objc.Send[AVMIDIPlayer](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMIDIPlayer creates a new AVMIDIPlayer instance.
func NewAVMIDIPlayer() AVMIDIPlayer {
	class := getAVMIDIPlayerClass()
	rv := objc.Send[AVMIDIPlayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a player to play a MIDI file with the specified soundbank.
//
// inURL: The URL of the file to play.
//
// bankURL: The URL of the sound bank. The sound bank must be in SoundFont2 or DLS
// format. In macOS, you can pass [nil] for the bank URL argument to use the
// default sound bank. In iOS, you must always pass a valid bank file.
// //
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// # Return Value
// 
// A new MIDI player, or [nil] if an error occurred.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/init(contentsOf:soundBankURL:)
func NewMIDIPlayerWithContentsOfURLSoundBankURLError(inURL foundation.INSURL, bankURL foundation.INSURL) (AVMIDIPlayer, error) {
	var errorPtr objc.ID
	instance := getAVMIDIPlayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:soundBankURL:error:"), inURL, bankURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVMIDIPlayer{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVMIDIPlayerFromID(rv), nil
}

// Creates a player to play MIDI data with the specified soundbank.
//
// data: The data to play.
//
// bankURL: The URL of the sound bank. The sound bank must be a SoundFont2 or DLS bank.
// In macOS, you can pass [nil] for the bank URL argument to use the default
// sound bank. In iOS, you must always pass a valid bank file.
// //
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// # Return Value
// 
// A new MIDI player, or [nil] if an error occurred.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/init(data:soundBankURL:)
func NewMIDIPlayerWithDataSoundBankURLError(data foundation.INSData, bankURL foundation.INSURL) (AVMIDIPlayer, error) {
	var errorPtr objc.ID
	instance := getAVMIDIPlayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:soundBankURL:error:"), data, bankURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVMIDIPlayer{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVMIDIPlayerFromID(rv), nil
}

// Creates a player to play a MIDI file with the specified soundbank.
//
// inURL: The URL of the file to play.
//
// bankURL: The URL of the sound bank. The sound bank must be in SoundFont2 or DLS
// format. In macOS, you can pass [nil] for the bank URL argument to use the
// default sound bank. In iOS, you must always pass a valid bank file.
// //
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// # Return Value
// 
// A new MIDI player, or [nil] if an error occurred.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/init(contentsOf:soundBankURL:)
func (m AVMIDIPlayer) InitWithContentsOfURLSoundBankURLError(inURL foundation.INSURL, bankURL foundation.INSURL) (AVMIDIPlayer, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithContentsOfURL:soundBankURL:error:"), inURL, bankURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVMIDIPlayer{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVMIDIPlayerFromID(rv), nil

}
// Creates a player to play MIDI data with the specified soundbank.
//
// data: The data to play.
//
// bankURL: The URL of the sound bank. The sound bank must be a SoundFont2 or DLS bank.
// In macOS, you can pass [nil] for the bank URL argument to use the default
// sound bank. In iOS, you must always pass a valid bank file.
// //
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// # Return Value
// 
// A new MIDI player, or [nil] if an error occurred.
//
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/init(data:soundBankURL:)
func (m AVMIDIPlayer) InitWithDataSoundBankURLError(data foundation.INSData, bankURL foundation.INSURL) (AVMIDIPlayer, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithData:soundBankURL:error:"), data, bankURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVMIDIPlayer{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVMIDIPlayerFromID(rv), nil

}
// Prepares the player to play the sequence by prerolling all events.
//
// # Discussion
// 
// The system automatically calls this method on playback, but calling it in
// advance minimizes the delay between calling [Play] and the start of sound
// output.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/prepareToPlay()
func (m AVMIDIPlayer) PrepareToPlay() {
	objc.Send[objc.ID](m.ID, objc.Sel("prepareToPlay"))
}
// Plays the MIDI sequence.
//
// completionHandler: A closure the system calls when playback completes.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/play(_:)
func (m AVMIDIPlayer) Play(completionHandler ErrorHandler) {
_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](m.ID, objc.Sel("play:"), _block0)
}
// Stops playing the sequence.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/stop()
func (m AVMIDIPlayer) Stop() {
	objc.Send[objc.ID](m.ID, objc.Sel("stop"))
}

// A Boolean value that indicates whether the sequence is playing.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/isPlaying
func (m AVMIDIPlayer) Playing() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isPlaying"))
	return rv
}
// The playback rate of the player.
//
// # Discussion
// 
// The default value is `1.0,` the standard playback rate.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/rate
func (m AVMIDIPlayer) Rate() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("rate"))
	return rv
}
func (m AVMIDIPlayer) SetRate(value float32) {
	objc.Send[struct{}](m.ID, objc.Sel("setRate:"), value)
}
// The current playback position, in seconds.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/currentPosition
func (m AVMIDIPlayer) CurrentPosition() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("currentPosition"))
	return rv
}
func (m AVMIDIPlayer) SetCurrentPosition(value float64) {
	objc.Send[struct{}](m.ID, objc.Sel("setCurrentPosition:"), value)
}
// The duration, in seconds, of the currently loaded file.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/duration
func (m AVMIDIPlayer) Duration() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("duration"))
	return rv
}

