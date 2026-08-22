// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

// Package coremidi provides Go bindings for the CoreMIDI framework.
//
// Communicate with MIDI devices such as hardware keyboards and synthesizers.
//
// The Core MIDI framework provides APIs to communicate with MIDI (Musical
// Instrument Digital Interface) devices, including hardware keyboards and
// synthesizers. Connect from an iOS device using the dock connector or a
// network. For more information about using the dock connector, see the [MFi
// Program].
//
// # Services
//
//   - [MIDI Services]: Communicate with hardware using Universal MIDI Packets. ([MIDIObjectRef], [MIDIClientRef], [MIDIDeviceRef], [MIDIEntityRef], [MIDIPortRef])
//   - [MIDI System Setup]: Configure the global MIDI system.
//   - [MIDI Bluetooth]: Connect to Bluetooth Low Energy MIDI peripherals.
//   - [MIDI Messages]: Create and configure messages. ([MIDICVStatus], [MIDIProtocolID], [MIDISysExStatus], [MIDISystemStatus], [MIDIMessage_128])
//   - [MIDI Thru Connection]: Create play-through connections between sources and destinations. ([MIDIThruConnectionRef], [MIDIThruConnectionEndpoint], [MIDIThruConnectionParams], [MIDIValueMap], [MIDIControlTransform])
//   - [MIDI Networking]: Create and manage devices connected over a local network. ([MIDINetworkHost], [MIDINetworkConnection], [MIDINetworkSession])
//   - [MIDI Drivers]: Create driver plug-ins. ([MIDIDeviceRef], [MIDIDeviceListRef], [MIDIDriverInterface], [MIDIDriverRef])
//   - [MIDI Capability Inquiry]: Provide support for bidirectional discovery and configuration of devices. ([MIDICIDiscoveryManager], [MIDICISession], [MIDICIProfile], [MIDICIProfileState], [MIDICIResponder])
//
// # Classes
//
//   - [MIDI2DeviceInfo]
//   - [MIDICIDevice]
//   - [MIDICIDeviceManager]
//   - [MIDIUMPCIProfile]
//   - [MIDIUMPEndpoint]
//   - [MIDIUMPEndpointManager]
//   - [MIDIUMPFunctionBlock]
//   - [MIDIUMPMutableEndpoint]
//   - [MIDIUMPMutableFunctionBlock]
//
// # Variables
//
//   - [KMIDIPropertyUMPEnabled]: kMIDIPropertyUMPEnabled//
//
// # Key Types
//
//   - [MIDINetworkSession] - An object that represents a pairing of a source and destination.
//   - [MIDIUMPEndpoint]
//   - [MIDIUMPFunctionBlock]
//   - [MIDICIDevice]
//   - [MIDINetworkHost] - An object that represents the host’s network address.
//   - [MIDIUMPCIProfile]
//   - [MIDIUMPMutableEndpoint]
//   - [MIDI2DeviceInfo]
//   - [MIDICIProfileState] - An object that provides the enabled and disabled profiles for a MIDI channel or port on a device.
//   - [MIDIUMPMutableFunctionBlock]
//
// [MIDI Bluetooth]: https://developer.apple.com/documentation/coremidi/midi-bluetooth
// [MIDI Capability Inquiry]: https://developer.apple.com/documentation/coremidi/midi-capability-inquiry
// [MIDI Drivers]: https://developer.apple.com/documentation/coremidi/midi-drivers
// [MIDI Messages]: https://developer.apple.com/documentation/coremidi/midi-messages
// [MIDI Networking]: https://developer.apple.com/documentation/coremidi/midi-networking
// [MIDI Services]: https://developer.apple.com/documentation/coremidi/midi-services
// [MIDI System Setup]: https://developer.apple.com/documentation/coremidi/midi-system-setup
// [MIDI Thru Connection]: https://developer.apple.com/documentation/coremidi/midi-thru-connection
package coremidi

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the CoreMIDI library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/CoreMIDI.framework/CoreMIDI",
	"/usr/lib/libCoreMIDI.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: CoreMIDI: failed to load framework from any known path\n")
	}
}
