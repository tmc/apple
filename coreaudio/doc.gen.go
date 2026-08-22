// Code generated from Apple documentation for CoreAudio. DO NOT EDIT.

// Package coreaudio provides Go bindings for the CoreAudio framework.
//
// Use the Core Audio framework to interact with device’s audio hardware.
//
// # Drivers
//
//   - [Creating an Audio Server Driver Plug-in]: Build a virtual audio device by creating a custom driver plug-in.
//   - [Building an Audio Server Plug-in and Driver Extension]: Create a plug-in and driver extension to support an audio device in macOS.
//   - [Capturing system audio with Core Audio taps]: Use a Core Audio tap to capture outgoing audio from a process or group of processes.
//
// # Classes
//
//   - [CATapDescription]//
//
// # Key Types
//
//   - [CATapDescription]
//
// [Building an Audio Server Plug-in and Driver Extension]: https://developer.apple.com/documentation/coreaudio/building-an-audio-server-plug-in-and-driver-extension
// [Capturing system audio with Core Audio taps]: https://developer.apple.com/documentation/coreaudio/capturing-system-audio-with-core-audio-taps
// [Creating an Audio Server Driver Plug-in]: https://developer.apple.com/documentation/coreaudio/creating-an-audio-server-driver-plug-in
package coreaudio

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the CoreAudio library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/CoreAudio.framework/CoreAudio",
	"/usr/lib/libCoreAudio.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: CoreAudio: failed to load framework from any known path\n")
	}
}
