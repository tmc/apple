// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

// Package soundanalysis provides Go bindings for the SoundAnalysis framework.
//
// Classify various sounds by analyzing audio files or streams.
//
// Identify specific sounds in your app, such as laughter or applause, by
// creating an [SNClassifySoundRequest] to analyze an audio file or stream.
// Sound requests can identify over 300 sounds. Alternatively, you identify a
// custom set of sounds by providing the sound request with a custom Core ML
// model. You train a custom sound classification model by creating an
// [MLSoundClassifier] with audio data in Create ML.
//
// # Audio analyzers
//
//   - [Classifying Sounds in an Audio File]: Identify individual sounds in a file, such as a recording, with an audio file analyzer.
//   - [SNAudioFileAnalyzer]: An analyzer that runs sound classification requests on an audio file. ([SNRequest], [SNResultsObserving])
//   - [Classifying Sounds in an Audio Stream]: Identify individual sounds in an audio data stream, such as from a microphone, with an audio stream analyzer.
//   - [SNAudioStreamAnalyzer]: An object you create to analyze a stream of audio data and provide the results to your app. ([SNRequest], [SNResultsObserving])
//
// # Sound classification requests
//
//   - [Classifying Live Audio Input with a Built-in Sound Classifier]: Detect and identify hundreds of sounds by using a trained classifier.
//   - [SNClassifySoundRequest]: A request that classifies sound using a Core ML model. ([SNClassifierIdentifier], [SNTimeDurationConstraint])
//   - [SNClassificationResult]: A result that contains the highest-ranking classifications in a time range. ([SNClassification])
//
// # Errors
//
//   - [SNErrorCode]: The enumerated error codes that the Sound Analysis framework produces.
//   - [SNErrorDomain]: A string that identifies the Sound Analysis error domain.//
//
// # Key Types
//
//   - [SNAudioFileAnalyzer] - An analyzer that runs sound classification requests on an audio file.
//   - [SNAudioStreamAnalyzer] - An object you create to analyze a stream of audio data and provide the results to your app.
//   - [SNClassifySoundRequest] - A request that classifies sound using a Core ML model.
//   - [SNTimeDurationConstraint] - Defines the time duration windows the request’s underlying sound classifier accepts with a range, or an array, of durations.
//   - [SNClassificationResult] - A result that contains the highest-ranking classifications in a time range.
//   - [SNClassification] - A type that pairs a sound classifier’s prediction with its confidence in that prediction.
//
// [Classifying Live Audio Input with a Built-in Sound Classifier]: https://developer.apple.com/documentation/soundanalysis/classifying-live-audio-input-with-a-built-in-sound-classifier
// [Classifying Sounds in an Audio File]: https://developer.apple.com/documentation/soundanalysis/classifying-sounds-in-an-audio-file
// [Classifying Sounds in an Audio Stream]: https://developer.apple.com/documentation/soundanalysis/classifying-sounds-in-an-audio-stream
package soundanalysis

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the SoundAnalysis library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/SoundAnalysis.framework/SoundAnalysis",
	"/usr/lib/libSoundAnalysis.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: SoundAnalysis: failed to load framework from any known path\n")
	}
}
