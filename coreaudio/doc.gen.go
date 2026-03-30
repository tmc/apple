// Code generated from Apple documentation for CoreAudio. DO NOT EDIT.

// Package coreaudio provides Go bindings for the CoreAudio framework.
//
// Use the Core Audio framework to interact with device’s audio hardware.
//
// # Drivers
//
//   - Creating an Audio Server Driver Plug-in: Build a virtual audio device by creating a custom driver plug-in.
//   - Building an Audio Server Plug-in and Driver Extension: Create a plug-in and driver extension to support an audio device in macOS.
//   - Capturing system audio with Core Audio taps: Use a Core Audio tap to capture outgoing audio from a process or group of processes.
//
// # Classes
//
//   - AudioHardwareAggregateDevice: Instances of the AudioHardwareAggregateDevice class encapsulate a single audio aggregate device, which is a virtual device that combines the input and output streams of multiple real devices or taps.
//   - AudioHardwareBox: Instances of the AudioHardwareBox class encapsulate a single audio box, which is a container for other objects (typically device objects).
//   - AudioHardwareClock: Instances of the AudioHardwareClock class encapsulate individual audio clocks.
//   - AudioHardwareControl: Instances of the AudioHardwareControl class encapsulate a single audio control, which provides properties that describe/manipulate a particular aspect of the owning object such as gain, mute, data source selection, etc.
//   - AudioHardwareDevice: Instances of the AudioHardwareDevice class encapsulate individual audio devices.
//   - AudioHardwareObject: The audio HAL provides an abstraction through which applications can access audio hardware.
//   - AudioHardwarePlugin: Instances of the AudioHardwarePlugin class encapsulate a single audio HAL plugin, which is a CFBundle loaded by the HAL as a driver to implement device-specific properties and routines.
//   - AudioHardwareProcess: Instances of the AudioHardwareProcess class encapsulate a single audio process, which contains information about a client process connected to the HAL.
//   - AudioHardwareStream: Instances of the AudioHardwareStream class encapsulate a single audio stream, which represents a single buffer of data for transferring across the user/kernel boundary.
//   - AudioHardwareSystem: The audio objects in the HAL are arranged in a containment hierarchy.
//   - AudioHardwareTap: Instances of the AudioHardwareTap class encapsulate a single audio tap, which can capture outgoing audio from a process or group of processes, and be used as an input stream source in an aggregate device.
//   - CATapDescription
//
// # Protocols
//
//   - PropertyListenerDelegate: A delegate protocol for receiving notifications when properties registered with AudioHardwareObject.addPropertyListener change.
//
// # Variables
//
//   - kAudioDevicePropertyWantsControlsRestored
//   - kAudioDevicePropertyWantsStreamFormatsRestored
//
// # Key Types
//
//   - [CATapDescription]
//
// [CoreAudio Documentation]: https://developer.apple.com/documentation/CoreAudio
package coreaudio

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the CoreAudio library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
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
	fmt.Fprintf(os.Stderr, "warning: CoreAudio: failed to load framework from any known path\n")
}
