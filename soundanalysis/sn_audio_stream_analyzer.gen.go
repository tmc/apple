// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/avfaudio"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SNAudioStreamAnalyzer] class.
var (
	_SNAudioStreamAnalyzerClass     SNAudioStreamAnalyzerClass
	_SNAudioStreamAnalyzerClassOnce sync.Once
)

func getSNAudioStreamAnalyzerClass() SNAudioStreamAnalyzerClass {
	_SNAudioStreamAnalyzerClassOnce.Do(func() {
		_SNAudioStreamAnalyzerClass = SNAudioStreamAnalyzerClass{class: objc.GetClass("SNAudioStreamAnalyzer")}
	})
	return _SNAudioStreamAnalyzerClass
}

// GetSNAudioStreamAnalyzerClass returns the class object for SNAudioStreamAnalyzer.
func GetSNAudioStreamAnalyzerClass() SNAudioStreamAnalyzerClass {
	return getSNAudioStreamAnalyzerClass()
}

type SNAudioStreamAnalyzerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SNAudioStreamAnalyzerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SNAudioStreamAnalyzerClass) Alloc() SNAudioStreamAnalyzer {
	rv := objc.Send[SNAudioStreamAnalyzer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An object you create to analyze a stream of audio data and provide the
// results to your app.
//
// # Overview
//
// Run an [SNRequest] on an audio stream by creating an
// [SNAudioStreamAnalyzer]. You can run the same sound analysis request on
// multiple stream analyzers, and each analyzer can process multiple requests.
// An audio file analyzer generates an [SNResult] each time any of its active
// requests recognizes a sound.
//
// # Creating an Analyzer
//
//   - [SNAudioStreamAnalyzer.InitWithFormat]: Creates a new audio stream analyzer.
//
// # Managing Requests
//
//   - [SNAudioStreamAnalyzer.AddRequestWithObserverError]: Adds a new analysis request to the audio stream analyzer.
//   - [SNAudioStreamAnalyzer.RemoveRequest]: Removes an existing request from the audio stream analyzer.
//   - [SNAudioStreamAnalyzer.RemoveAllRequests]: Removes all the sound analysis requests from the audio stream analyzer.
//
// # Analyzing Data
//
//   - [SNAudioStreamAnalyzer.AnalyzeAudioBufferAtAudioFramePosition]: Adds a new audio buffer to the analyzer’s larger stream buffer.
//   - [SNAudioStreamAnalyzer.CompleteAnalysis]: Notifies the analyzer when it receives the final audio buffer.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNAudioStreamAnalyzer
type SNAudioStreamAnalyzer struct {
	objectivec.Object
}

// SNAudioStreamAnalyzerFromID constructs a [SNAudioStreamAnalyzer] from an objc.ID.
//
// An object you create to analyze a stream of audio data and provide the
// results to your app.
func SNAudioStreamAnalyzerFromID(id objc.ID) SNAudioStreamAnalyzer {
	return SNAudioStreamAnalyzer{objectivec.Object{ID: id}}
}

// NOTE: SNAudioStreamAnalyzer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SNAudioStreamAnalyzer] class.
//
// # Creating an Analyzer
//
//   - [ISNAudioStreamAnalyzer.InitWithFormat]: Creates a new audio stream analyzer.
//
// # Managing Requests
//
//   - [ISNAudioStreamAnalyzer.AddRequestWithObserverError]: Adds a new analysis request to the audio stream analyzer.
//   - [ISNAudioStreamAnalyzer.RemoveRequest]: Removes an existing request from the audio stream analyzer.
//   - [ISNAudioStreamAnalyzer.RemoveAllRequests]: Removes all the sound analysis requests from the audio stream analyzer.
//
// # Analyzing Data
//
//   - [ISNAudioStreamAnalyzer.AnalyzeAudioBufferAtAudioFramePosition]: Adds a new audio buffer to the analyzer’s larger stream buffer.
//   - [ISNAudioStreamAnalyzer.CompleteAnalysis]: Notifies the analyzer when it receives the final audio buffer.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNAudioStreamAnalyzer
type ISNAudioStreamAnalyzer interface {
	objectivec.IObject

	// Topic: Creating an Analyzer

	// Creates a new audio stream analyzer.
	InitWithFormat(format avfaudio.AVAudioFormat) SNAudioStreamAnalyzer

	// Topic: Managing Requests

	// Adds a new analysis request to the audio stream analyzer.
	AddRequestWithObserverError(request SNRequest, observer SNResultsObserving) (bool, error)
	// Removes an existing request from the audio stream analyzer.
	RemoveRequest(request SNRequest)
	// Removes all the sound analysis requests from the audio stream analyzer.
	RemoveAllRequests()

	// Topic: Analyzing Data

	// Adds a new audio buffer to the analyzer’s larger stream buffer.
	AnalyzeAudioBufferAtAudioFramePosition(audioBuffer avfaudio.AVAudioBuffer, audioFramePosition avfaudio.AVAudioFramePosition)
	// Notifies the analyzer when it receives the final audio buffer.
	CompleteAnalysis()
}

// Init initializes the instance.
func (a SNAudioStreamAnalyzer) Init() SNAudioStreamAnalyzer {
	rv := objc.Send[SNAudioStreamAnalyzer](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a SNAudioStreamAnalyzer) Autorelease() SNAudioStreamAnalyzer {
	rv := objc.Send[SNAudioStreamAnalyzer](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewSNAudioStreamAnalyzer creates a new SNAudioStreamAnalyzer instance.
func NewSNAudioStreamAnalyzer() SNAudioStreamAnalyzer {
	class := getSNAudioStreamAnalyzerClass()
	rv := objc.Send[SNAudioStreamAnalyzer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new audio stream analyzer.
//
// format: The audio format of an audio stream.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNAudioStreamAnalyzer/init(format:)
func NewAudioStreamAnalyzerWithFormat(format avfaudio.AVAudioFormat) SNAudioStreamAnalyzer {
	instance := getSNAudioStreamAnalyzerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:"), format)
	return SNAudioStreamAnalyzerFromID(rv)
}

// Creates a new audio stream analyzer.
//
// format: The audio format of an audio stream.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNAudioStreamAnalyzer/init(format:)
func (a SNAudioStreamAnalyzer) InitWithFormat(format avfaudio.AVAudioFormat) SNAudioStreamAnalyzer {
	rv := objc.Send[SNAudioStreamAnalyzer](a.ID, objc.Sel("initWithFormat:"), format)
	return rv
}

// Adds a new analysis request to the audio stream analyzer.
//
// request: A sound analysis request.
//
// observer: An [SNResultsObserving] instance that receives the analyzer’s results.
// The analyzer maintains a weak reference to the observer.
//
// # Discussion
//
// You can add requests to an analyzer that’s actively analyzing an audio
// stream. The analyzer throws an error (Swift) or returns NO (Objective-C) if
// it can’t accept the new request, such as a request with an audio format
// that doesn’t match the analyzer’s.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNAudioStreamAnalyzer/add(_:withObserver:)
func (a SNAudioStreamAnalyzer) AddRequestWithObserverError(request SNRequest, observer SNResultsObserving) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("addRequest:withObserver:error:"), request, observer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addRequest:withObserver:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Removes an existing request from the audio stream analyzer.
//
// request: A sound analysis request.
//
// # Discussion
//
// You can remove a request while the analyzer is processing it. The analyzer
// stops sending results to the observer after the method removes the request.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNAudioStreamAnalyzer/remove(_:)
func (a SNAudioStreamAnalyzer) RemoveRequest(request SNRequest) {
	objc.Send[objc.ID](a.ID, objc.Sel("removeRequest:"), request)
}

// Removes all the sound analysis requests from the audio stream analyzer.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNAudioStreamAnalyzer/removeAllRequests()
func (a SNAudioStreamAnalyzer) RemoveAllRequests() {
	objc.Send[objc.ID](a.ID, objc.Sel("removeAllRequests"))
}

// Adds a new audio buffer to the analyzer’s larger stream buffer.
//
// audioBuffer: The audio data buffer.
//
// audioFramePosition: The frame position of the data in the buffer. The analyzer expects this
// parameter to monotonically increase with each call. Otherwise, the analyzer
// may reset its internal state to account for the jump in time.
//
// # Discussion
//
// Serialize all calls to this method on a dedicated dispatch queue to prevent
// blocking the calling thread. The audio stream analyzer sends errors to each
// request’s results observer.
//
// The method handles audio buffers that vary in size. If necessary, the
// analyzer regroups the data to a block size the underlying Core ML model
// expects. The analyzer uses a fixed-size audio block from some audio
// analysis types and may call a request’s results observer one time or many
// times. The factors that affect the number of calls are:
//
// - The input buffer size - The underlying model’s native analysis block
// size - The analyzer’s current state
//
// By default, the analyzer processes data on the first audio channel in the
// audio stream. The analyzer converts the sample rate of the data if it
// doesn’t match the underlying model’s requirements.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNAudioStreamAnalyzer/analyze(_:atAudioFramePosition:)
func (a SNAudioStreamAnalyzer) AnalyzeAudioBufferAtAudioFramePosition(audioBuffer avfaudio.AVAudioBuffer, audioFramePosition avfaudio.AVAudioFramePosition) {
	objc.Send[objc.ID](a.ID, objc.Sel("analyzeAudioBuffer:atAudioFramePosition:"), audioBuffer, audioFramePosition)
}

// Notifies the analyzer when it receives the final audio buffer.
//
// # Discussion
//
// Use this method for requests that provide final results when a stream
// reaches its end. The analyzer ignores any further calls to the
// [SNAudioStreamAnalyzer.AnalyzeAudioBufferAtAudioFramePosition] method.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNAudioStreamAnalyzer/completeAnalysis()
func (a SNAudioStreamAnalyzer) CompleteAnalysis() {
	objc.Send[objc.ID](a.ID, objc.Sel("completeAnalysis"))
}
