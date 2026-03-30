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

// The class instance for the [AVAudioUnitSampler] class.
var (
	_AVAudioUnitSamplerClass     AVAudioUnitSamplerClass
	_AVAudioUnitSamplerClassOnce sync.Once
)

func getAVAudioUnitSamplerClass() AVAudioUnitSamplerClass {
	_AVAudioUnitSamplerClassOnce.Do(func() {
		_AVAudioUnitSamplerClass = AVAudioUnitSamplerClass{class: objc.GetClass("AVAudioUnitSampler")}
	})
	return _AVAudioUnitSamplerClass
}

// GetAVAudioUnitSamplerClass returns the class object for AVAudioUnitSampler.
func GetAVAudioUnitSamplerClass() AVAudioUnitSamplerClass {
	return getAVAudioUnitSamplerClass()
}

type AVAudioUnitSamplerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitSamplerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitSamplerClass) Alloc() AVAudioUnitSampler {
	rv := objc.Send[AVAudioUnitSampler](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that you configure with one or more instrument samples, based on
// Apple’s Sampler audio unit.
//
// # Overview
//
// An [AVAudioUnitSampler] is an [AVAudioUnit] for Apple’s Sampler audio
// unit.
//
// You configure the sampler by loading instruments from different types of
// files. These include an `aupreset` file, DLS, or SF2 sound bank; an EXS24
// instrument; a single audio file; or an array of audio files.
//
// The output of a [AVAudioUnitSampler] is a single stereo bus.
//
// # Configuring the Sampler Audio Unit
//
//   - [AVAudioUnitSampler.LoadInstrumentAtURLError]: Configures the sampler with the specified instrument file.
//   - [AVAudioUnitSampler.LoadAudioFilesAtURLsError]: Configures the sampler by loading the specified audio files.
//   - [AVAudioUnitSampler.LoadSoundBankInstrumentAtURLProgramBankMSBBankLSBError]: Loads a specific instrument from the specified soundbank.
//
// # Getting and Setting Sampler Values
//
//   - [AVAudioUnitSampler.GlobalTuning]: An adjustment for the tuning of all the played notes.
//   - [AVAudioUnitSampler.SetGlobalTuning]
//   - [AVAudioUnitSampler.OverallGain]: An adjustment for the gain of all the played notes, in decibels.
//   - [AVAudioUnitSampler.SetOverallGain]
//   - [AVAudioUnitSampler.StereoPan]: An adjustment for the stereo panning of all the played notes.
//   - [AVAudioUnitSampler.SetStereoPan]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler
type AVAudioUnitSampler struct {
	AVAudioUnitMIDIInstrument
}

// AVAudioUnitSamplerFromID constructs a [AVAudioUnitSampler] from an objc.ID.
//
// An object that you configure with one or more instrument samples, based on
// Apple’s Sampler audio unit.
func AVAudioUnitSamplerFromID(id objc.ID) AVAudioUnitSampler {
	return AVAudioUnitSampler{AVAudioUnitMIDIInstrument: AVAudioUnitMIDIInstrumentFromID(id)}
}

// NOTE: AVAudioUnitSampler adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioUnitSampler] class.
//
// # Configuring the Sampler Audio Unit
//
//   - [IAVAudioUnitSampler.LoadInstrumentAtURLError]: Configures the sampler with the specified instrument file.
//   - [IAVAudioUnitSampler.LoadAudioFilesAtURLsError]: Configures the sampler by loading the specified audio files.
//   - [IAVAudioUnitSampler.LoadSoundBankInstrumentAtURLProgramBankMSBBankLSBError]: Loads a specific instrument from the specified soundbank.
//
// # Getting and Setting Sampler Values
//
//   - [IAVAudioUnitSampler.GlobalTuning]: An adjustment for the tuning of all the played notes.
//   - [IAVAudioUnitSampler.SetGlobalTuning]
//   - [IAVAudioUnitSampler.OverallGain]: An adjustment for the gain of all the played notes, in decibels.
//   - [IAVAudioUnitSampler.SetOverallGain]
//   - [IAVAudioUnitSampler.StereoPan]: An adjustment for the stereo panning of all the played notes.
//   - [IAVAudioUnitSampler.SetStereoPan]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler
type IAVAudioUnitSampler interface {
	IAVAudioUnitMIDIInstrument

	// Topic: Configuring the Sampler Audio Unit

	// Configures the sampler with the specified instrument file.
	LoadInstrumentAtURLError(instrumentURL foundation.INSURL) (bool, error)
	// Configures the sampler by loading the specified audio files.
	LoadAudioFilesAtURLsError(audioFiles []foundation.NSURL) (bool, error)
	// Loads a specific instrument from the specified soundbank.
	LoadSoundBankInstrumentAtURLProgramBankMSBBankLSBError(bankURL foundation.INSURL, program uint8, bankMSB uint8, bankLSB uint8) (bool, error)

	// Topic: Getting and Setting Sampler Values

	// An adjustment for the tuning of all the played notes.
	GlobalTuning() float32
	SetGlobalTuning(value float32)
	// An adjustment for the gain of all the played notes, in decibels.
	OverallGain() float32
	SetOverallGain(value float32)
	// An adjustment for the stereo panning of all the played notes.
	StereoPan() float32
	SetStereoPan(value float32)
}

// Init initializes the instance.
func (a AVAudioUnitSampler) Init() AVAudioUnitSampler {
	rv := objc.Send[AVAudioUnitSampler](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitSampler) Autorelease() AVAudioUnitSampler {
	rv := objc.Send[AVAudioUnitSampler](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitSampler creates a new AVAudioUnitSampler instance.
func NewAVAudioUnitSampler() AVAudioUnitSampler {
	class := getAVAudioUnitSamplerClass()
	rv := objc.Send[AVAudioUnitSampler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a MIDI instrument audio unit with the component description you
// specify.
//
// description: The description of the audio component.
//
// # Return Value
//
// A new [AVAudioUnitMIDIInstrument] instance.
//
// # Discussion
//
// The component type must be `kAudioUnitType_MusicDevice` or
// `kAudioUnitType_RemoteInstrument`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/init(audioComponentDescription:)
// description is a [audiotoolbox.AudioComponentDescription].
func NewAudioUnitSamplerWithAudioComponentDescription(description objectivec.IObject) AVAudioUnitSampler {
	instance := getAVAudioUnitSamplerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAudioComponentDescription:"), description)
	return AVAudioUnitSamplerFromID(rv)
}

// Configures the sampler with the specified instrument file.
//
// instrumentURL: The URL of the file that contains the instrument.
//
// # Discussion
//
// The instrument can be one of the following types: Logic or GarageBand
// [EXS24], the sampler’s native `aupreset` file, or an audio file, such as
// `caf`, `aiff`, `wav`, or `mp3`.
//
// For a single audio file, the framework loads it into a new default
// instrument and uses any information in the audio file, such as the root key
// and key range, for its placement in the instrument.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler/loadInstrument(at:)
func (a AVAudioUnitSampler) LoadInstrumentAtURLError(instrumentURL foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("loadInstrumentAtURL:error:"), instrumentURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("loadInstrumentAtURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Configures the sampler by loading the specified audio files.
//
// audioFiles: An array of audio file URLs to load.
//
// # Discussion
//
// The framework loads the audio files into a new instrument with each audio
// file in its own sampler zone. The framework uses any information in the
// audio file for its placement in the instrument. For example, the root key
// and key range.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler/loadAudioFiles(at:)
func (a AVAudioUnitSampler) LoadAudioFilesAtURLsError(audioFiles []foundation.NSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("loadAudioFilesAtURLs:error:"), objectivec.IObjectSliceToNSArray(audioFiles), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("loadAudioFilesAtURLs:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Loads a specific instrument from the specified soundbank.
//
// bankURL: The URL for a soundbank file, either a DLS bank (`XCUIElementTypeDls`) or a
// SoundFont bank (`XCUIElementTypeSf2`).
//
// program: The program number for the instrument to load.
//
// bankMSB: The most significant bit for the bank number for the instrument to load.
// This is usually `0x79` for melodic instruments and `0x78` for percussion
// instruments.
//
// bankLSB: The least significant bit for the bank number for the instrument to load.
// This is often `0` and represents the bank variation.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler/loadSoundBankInstrument(at:program:bankMSB:bankLSB:)
func (a AVAudioUnitSampler) LoadSoundBankInstrumentAtURLProgramBankMSBBankLSBError(bankURL foundation.INSURL, program uint8, bankMSB uint8, bankLSB uint8) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("loadSoundBankInstrumentAtURL:program:bankMSB:bankLSB:error:"), bankURL, program, bankMSB, bankLSB, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("loadSoundBankInstrumentAtURL:program:bankMSB:bankLSB:error: returned NO with nil NSError")
	}
	return rv, nil

}

// An adjustment for the tuning of all the played notes.
//
// # Discussion
//
// The tuning unit is cents, and defaults to `0.0`. The range of valid values
// is `-2400` to `2400` cents.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler/globalTuning
func (a AVAudioUnitSampler) GlobalTuning() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("globalTuning"))
	return rv
}
func (a AVAudioUnitSampler) SetGlobalTuning(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setGlobalTuning:"), value)
}

// An adjustment for the gain of all the played notes, in decibels.
//
// # Discussion
//
// The default value is `0.0` dB, and the range of valid values is `-90.0` dB
// to `12.0` dB.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler/overallGain
func (a AVAudioUnitSampler) OverallGain() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("overallGain"))
	return rv
}
func (a AVAudioUnitSampler) SetOverallGain(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setOverallGain:"), value)
}

// An adjustment for the stereo panning of all the played notes.
//
// # Discussion
//
// The default value is `0.0`, and the range of valid values is `-100.0` to
// `100.0`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler/stereoPan
func (a AVAudioUnitSampler) StereoPan() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("stereoPan"))
	return rv
}
func (a AVAudioUnitSampler) SetStereoPan(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setStereoPan:"), value)
}
