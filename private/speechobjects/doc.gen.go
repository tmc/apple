
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

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/PrivateFrameworks/SpeechObjects.framework/SpeechObjects"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: speechobjects: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

