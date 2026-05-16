// Code generated from Apple documentation for ApplicationServices. DO NOT EDIT.

// Package applicationservices provides Go bindings for the ApplicationServices framework.
//
// Perform common application tasks.
//
// This collection of documents provides the API reference for the Application Services framework, which includes several services that are essential to Carbon applications. The Application Services framework also includes support for a number of legacy technologies—such as QuickDraw and the Font Manager—that have been superseded with newer technologies like Quartz 2D and ATSUI.
//
// # Managers
//
//   - Apple Event Manager
//   - ColorSync Manager ([CMFlattenProcPtr], [CM2Profile], [CMDeviceInfo], [CMDeviceProfileArray], [CMDeviceScope])
//   - Speech Synthesis Manager ([SpeechDoneProcPtr], [SpeechErrorProcPtr], [SpeechErrorCFProcPtr], [SpeechPhonemeProcPtr], [SpeechSyncProcPtr])
//
// # Classes
//
//   - ColorSyncCMM
//   - ColorSyncMutableProfile
//   - ColorSyncProfile
//   - ColorSyncTransform
//   - HIMutableShape
//   - HIShape
//   - Pasteboard
//   - Translation
//   - AXTextMarker
//   - AXTextMarkerRange
//
// # Protocols
//
//   - PDEPanel
//   - PDEPlugIn
//   - PDEPlugInCallbackProtocol
//
// [ApplicationServices Documentation]: https://developer.apple.com/documentation/applicationservices
package applicationservices

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the ApplicationServices library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/ApplicationServices.framework/ApplicationServices",
	"/usr/lib/libApplicationServices.dylib",
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
	fmt.Fprintf(os.Stderr, "warning: ApplicationServices: failed to load framework from any known path\n")
}
