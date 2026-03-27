
// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

// Package texttospeech provides Go bindings for the texttospeech framework.
//
// # Key Types
//
//   - [TTSSpeechManager]
//   - [TTSSpeechSynthesizer]
//   - [TTSSpeechAction]
//   - [TTSWrappedAudioQueue]
//   - [TTSSpeechRequest]
//   - [TTSAXResource]
//   - [TextToSpeechCoreSynthesizer]
//   - [TTSSubstitution]
//   - [TTSSiriAssetManager]
//   - [TTSSpeechString]
package texttospeech

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/PrivateFrameworks/TextToSpeech.framework/TextToSpeech"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: texttospeech: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

