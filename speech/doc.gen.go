// Code generated from Apple documentation for Speech. DO NOT EDIT.

// Package speech provides Go bindings for the Speech framework.
//
// Perform speech recognition on live or prerecorded audio, and receive
// transcriptions, alternative interpretations, and confidence levels of the
// results.
//
// Use the Speech framework to recognize spoken words in recorded or live
// audio. The keyboard’s dictation support uses speech recognition to
// translate audio content into text. This framework provides a similar
// behavior, except that you can use it without the presence of the keyboard.
// For example, you might use speech recognition to recognize verbal commands
// or to handle text dictation in other parts of your app.
//
// # Essentials
//
//   - [Speech Recognition in Objective-C]: Use these classes to perform speech recognition in Objective-C code. ([SFSpeechRecognizer], [SFSpeechRecognizerDelegate], [SFSpeechRecognitionTaskHint], [SFSpeechRecognizerAuthorizationStatus], [SFSpeechURLRecognitionRequest])
//
// # Custom vocabulary
//
//   - [SFSpeechLanguageModel]: A language model built from custom training data.
//   - [SFSpeechLanguageModelConfiguration]: An object describing the location of a custom language model and specialized vocabulary.//
//
// # Key Types
//
//   - [SFSpeechRecognizer] - An object you use to check for the availability of the speech recognition service, and to initiate the speech recognition process.
//   - [SFSpeechLanguageModelConfiguration] - An object describing the location of a custom language model and specialized vocabulary.
//   - [SFSpeechRecognitionRequest] - An abstract class that represents a request to recognize speech from an audio source.
//   - [SFSpeechRecognitionTask] - A task object for monitoring the speech recognition progress.
//   - [SFTranscriptionSegment] - A discrete part of an entire transcription, as identified by the speech recognizer.
//   - [SFSpeechRecognitionMetadata] - The metadata of speech in the audio of a speech recognition request.
//   - [SFSpeechAudioBufferRecognitionRequest] - A request to recognize speech from captured audio content, such as audio from the device’s microphone.
//   - [SFSpeechLanguageModel] - A language model built from custom training data.
//   - [SFSpeechRecognitionResult] - An object that contains the partial or final results of a speech recognition request.
//   - [SFVoiceAnalytics] - A collection of vocal analysis metrics.
//
// [Speech Recognition in Objective-C]: https://developer.apple.com/documentation/speech/speech-recognition-in-objc
package speech

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the Speech library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/Speech.framework/Speech",
	"/usr/lib/libSpeech.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: Speech: failed to load framework from any known path\n")
	}
}
