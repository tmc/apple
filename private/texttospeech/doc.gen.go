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

// frameworkPaths lists paths to try when loading the texttospeech library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{"/System/Library/PrivateFrameworks/TextToSpeech.framework/TextToSpeech"}

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
	fmt.Fprintf(os.Stderr, "warning: texttospeech: failed to load framework from any known path\n")
}
