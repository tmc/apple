
// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

// Package avfaudio provides Go bindings for the AVFAudio framework.
//
// Play, record, and process audio; configure your app’s system audio behavior.
//
// # Essentials
//
//   - AVFAudio updates: Learn about important changes to AVFAudio.
//
// # System audio
//
//   - Handling audio interruptions: Observe audio session notifications to ensure that your app responds appropriately to interruptions.
//   - Responding to audio route changes: Observe audio session notifications to ensure that your app responds appropriately to route changes.
//   - Routing audio to specific devices in multidevice sessions: Map audio channels to specific devices in multiroute sessions for recording and playback.
//   - Adding synthesized speech to calls: Provide a more accessible experience by adding your app’s audio to a call.
//   - Capturing stereo audio from built-In microphones: Configure an iOS device’s built-in microphones to add stereo recording capabilities to your app.
//   - AVAudioSession: An object that communicates to the system how you intend to use audio in your app. ([AVAudioSessionSpatialExperience], [AVAudioSessionActivationOptions])
//   - AVAudioApplication: An object that manages one or more audio sessions that belong to an app.
//   - AVAudioRoutingArbiter: An object for configuring macOS apps to participate in AirPods Automatic Switching.
//
// # Basic playback and recording
//
//   - AVAudioPlayer: An object that plays audio data from a file or buffer. ([AVAudioPlayerDelegate])
//   - AVAudioRecorder: An object that records audio data to a file. ([AVAudioRecorderDelegate])
//   - AVMIDIPlayer: An object that plays MIDI data through a system sound module. ([AVMIDIPlayerCompletionHandler])
//
// # Advanced audio processing
//
//   - Audio Engine: Perform advanced real-time and offline audio processing, implement 3D spatialization, and work with MIDI and samplers. ([AVAudioEngine], [AVAudioNode], [AVAudioInputNode], [AVAudioOutputNode], [AVAudioIONode])
//
// # Speech synthesis
//
//   - Speech synthesis: Configure voices to speak strings of text. ([AVSpeechUtterance], [AVSpeechSynthesisVoice], [AVSpeechSynthesizer], [AVSpeechSynthesisProviderAudioUnit])
//
// # Macros
//
//   - Macros
//
// # Key Types
//
//   - [AVAudioPlayerNode] - An object for scheduling the playback of buffers or segments of audio files.
//   - [AVAudioEngine] - An object that manages a graph of audio nodes, controls playback, and configures real-time rendering constraints.
//   - [AVAudioEnvironmentNode] - An object that simulates a 3D audio environment.
//   - [AVAudioUnitMIDIInstrument] - An object that represents music devices or remote instruments.
//   - [AVAudioPlayer] - An object that plays audio data from a file or buffer.
//   - [AVAudioUnitComponent] - An object that provides details about an audio unit.
//   - [AVAudioUnitSampler] - An object that you configure with one or more instrument samples, based on Apple’s Sampler audio unit.
//   - [AVAudioInputNode] - An object that connects to the system’s audio input.
//   - [AVAudioConverter] - An object that converts streams of audio between formats.
//   - [AVAudioSequencer] - An object that plays audio from a collection of MIDI events the system organizes into music tracks.
//
// [AVFAudio Documentation]: https://developer.apple.com/documentation/AVFAudio
package avfaudio

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/AVFAudio.framework/AVFAudio"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: AVFAudio: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

