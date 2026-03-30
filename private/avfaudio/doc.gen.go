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

// frameworkPaths lists paths to try when loading the AVFAudio library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{"/System/Library/Frameworks/AVFAudio.framework/AVFAudio"}

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
	fmt.Fprintf(os.Stderr, "warning: AVFAudio: failed to load framework from any known path\n")
}
