// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

// Package speechobjects provides Go bindings for the speechobjects framework.
//
// # Key Types
//
//   - [SOSpeechItem]
//   - [SOVoiceObject]
//   - [SOSpeechInstallationManager]
//   - [CustomizeVoicesWindowController]
//   - [SOUtteranceResultsFile]
//   - [SOVoicePopUpButton]
//   - [SOSRLanguagePopUpButton]
//   - [SOUtteranceResult]
//   - [SOCustomizeSRLanguagesWindowController]
//   - [VoiceSettingsWindowController]
package speechobjects

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the speechobjects library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{"/System/Library/PrivateFrameworks/SpeechObjects.framework/SpeechObjects"}

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
	fmt.Fprintf(os.Stderr, "warning: speechobjects: failed to load framework from any known path\n")
}
