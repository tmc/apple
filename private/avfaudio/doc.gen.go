
// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

// Package avfaudio provides Go bindings for the AVFAudio framework.
//
// # Key Types
//
//   - [AVVoiceController]
//   - [AVSpeechSynthesisVoice]
//   - [AVVCAudioDeviceManager]
//   - [AVVoiceTriggerClient]
//   - [AVSpeechSynthesizer]
//   - [AVVCMetricsManager]
//   - [AVAudioApplication]
//   - [AVAudioNode]
//   - [AVSpeechSynthesisProviderVoice]
//   - [AVVoiceTriggerClientPortManager]
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

