// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

// Package avfaudio provides Go bindings for the AVFAudio framework.
//
// Play, record, and process audio; configure your app’s system audio
// behavior.
//
// # Essentials
//
//   - [AVFAudio updates]: Learn about important changes to AVFAudio.
//
// # System audio
//
//   - [Handling audio interruptions]: Observe audio session notifications to ensure that your app responds appropriately to interruptions.
//   - [Responding to audio route changes]: Observe audio session notifications to ensure that your app responds appropriately to route changes.
//   - [Routing audio to specific devices in multidevice sessions]: Map audio channels to specific devices in multiroute sessions for recording and playback.
//   - [Adding synthesized speech to calls]: Provide a more accessible experience by adding your app’s audio to a call.
//   - [Capturing stereo audio from built-In microphones]: Configure an iOS device’s built-in microphones to add stereo recording capabilities to your app.
//   - AVAudioSession: An object that communicates to the system how you intend to use audio in your app. ([AVAudioSessionSpatialExperience], [AVAudioSessionActivationOptions])
//   - [AVAudioApplication]: An object that manages one or more audio sessions that belong to an app.
//   - [AVAudioRoutingArbiter]: An object for configuring macOS apps to participate in AirPods Automatic Switching.
//
// # Basic playback and recording
//
//   - [AVAudioPlayer]: An object that plays audio data from a file or buffer. ([AVAudioPlayerDelegate])
//   - [AVAudioRecorder]: An object that records audio data to a file. ([AVAudioRecorderDelegate])
//   - [AVMIDIPlayer]: An object that plays MIDI data through a system sound module. ([AVMIDIPlayerCompletionHandler])
//
// # Advanced audio processing
//
//   - [Audio Engine]: Perform advanced real-time and offline audio processing, implement 3D spatialization, and work with MIDI and samplers. ([AVAudioEngine], [AVAudioNode], [AVAudioInputNode], [AVAudioOutputNode], [AVAudioIONode])
//
// # Speech synthesis
//
//   - [Speech synthesis]: Configure voices to speak strings of text. ([AVSpeechUtterance], [AVSpeechSynthesisVoice], [AVSpeechSynthesizer], [AVSpeechSynthesisProviderAudioUnit])//
//
// # Key Types
//
//   - [AVAudioEngine] - An object that manages a graph of audio nodes, controls playback, and configures real-time rendering constraints.
//   - [AVAudioPlayerNode] - An object for scheduling the playback of buffers or segments of audio files.
//   - [AVAudioUnitMIDIInstrument] - An object that represents music devices or remote instruments.
//   - [AVAudioEnvironmentNode] - An object that simulates a 3D audio environment.
//   - [AVAudioPlayer] - An object that plays audio data from a file or buffer.
//   - [AVAudioUnitSampler] - An object that you configure with one or more instrument samples, based on Apple’s Sampler audio unit.
//   - [AVAudioInputNode] - An object that connects to the system’s audio input.
//   - [AVAudioConverter] - An object that converts streams of audio between formats.
//   - [AVAudioMixerNode] - An object that takes any number of inputs and converts them into a single output.
//   - [AVAudioSequencer] - An object that plays audio from a collection of MIDI events the system organizes into music tracks.
//
// [AVFAudio updates]: https://developer.apple.com/documentation/Updates/AVFAudio
// [Adding synthesized speech to calls]: https://developer.apple.com/documentation/avfaudio/adding-synthesized-speech-to-calls
// [Audio Engine]: https://developer.apple.com/documentation/avfaudio/audio-engine
// [Capturing stereo audio from built-In microphones]: https://developer.apple.com/documentation/avfaudio/capturing-stereo-audio-from-built-in-microphones
// [Handling audio interruptions]: https://developer.apple.com/documentation/avfaudio/handling-audio-interruptions
// [Responding to audio route changes]: https://developer.apple.com/documentation/avfaudio/responding-to-audio-route-changes
// [Routing audio to specific devices in multidevice sessions]: https://developer.apple.com/documentation/avfaudio/routing-audio-to-specific-devices-in-multidevice-sessions
// [Speech synthesis]: https://developer.apple.com/documentation/avfaudio/speech-synthesis
package avfaudio

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the AVFAudio library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/AVFAudio.framework/AVFAudio",
	"/usr/lib/libAVFAudio.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: AVFAudio: failed to load framework from any known path\n")
	}
}
