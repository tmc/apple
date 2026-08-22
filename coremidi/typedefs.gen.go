// Code generated from Apple documentation. DO NOT EDIT.

package coremidi

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
)

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceID
type MIDICIDeviceID = uint8

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceManager/DictionaryKey
type MIDICIDeviceManagerDictionaryKey = string

// MIDICIDiscoveryResponseBlock is a block the system calls when a MIDI-CI node discovery request completes.
//
// Deprecated: Deprecated since macOS 15.0. No longer supported for CoreMIDI
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveryResponseBlock
type MIDICIDiscoveryResponseBlock = func(discoveredNodes []MIDICIDiscoveredNode)

// MIDICIInitiatiorMUID is the unique MIDI-CI negotiation identifier to use for a responder connection.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIInitiatiorMUID
type MIDICIInitiatiorMUID = *foundation.NSNumber

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIMUID
type MIDICIMUID = uint32

// MIDICIProfileChangedBlock is a block the system calls to indicate it has enabled or disabled a profile.
//
// Deprecated: Deprecated since macOS 15.0. No longer supported for CoreMIDI
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileChangedBlock
type MIDICIProfileChangedBlock = func(session MIDICISession, channel uint32, profile MIDICIProfile, enabled int32)

// MIDICIProfileSpecificDataBlock is a block the system calls when a MIDI-CI session or responder receives profile-specific data.
//
// Deprecated: Deprecated since macOS 15.0. No longer supported for CoreMIDI
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileSpecificDataBlock
type MIDICIProfileSpecificDataBlock = func(session MIDICISession, channel uint32, profile MIDICIProfile, profileSpecificData foundation.NSData)

// MIDICIProfileStateList is an array of profile state objects that describes the profile configuration for all channels of a reachable MIDI-CI node.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileStateList
type MIDICIProfileStateList = []*MIDICIProfileState

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIPropertyExchangeRequestID
type MIDICIPropertyExchangeRequestID = uint8

// MIDICISessionDisconnectBlock is a block the system calls when a MIDI-CI session disconnects.
//
// Deprecated: Deprecated since macOS 15.0. No longer supported for CoreMIDI
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICISessionDisconnectBlock
type MIDICISessionDisconnectBlock = func(session MIDICISession, error_ foundation.NSError)

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIChannelNumber
type MIDIChannelNumber = uint8

// MIDIClientRef is an object that maintains per-client state.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIClientRef
type MIDIClientRef = uint32

// MIDICompletionProc is a function the system calls after it completely sends a system-exclusive (SysEx) event.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICompletionProc
type MIDICompletionProc = func(request uintptr)

// MIDICompletionProcUMP is a function the system calls after it completely sends a UMP system-exclusive (SysEx) or SysEx 8-bit event.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICompletionProcUMP
type MIDICompletionProcUMP = func(request uintptr)

// MIDIDeviceListRef is a list of MIDI devices.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceListRef
type MIDIDeviceListRef = uint32

// MIDIDeviceRef is a MIDI device that contains entities.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIDeviceRef
type MIDIDeviceRef = uint32

// MIDIDriverRef is a MIDI driver object.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIDriverRef
type MIDIDriverRef = **MIDIDriverInterface

// MIDIEndpointRef is a MIDI source or destination an entity owns.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIEndpointRef
type MIDIEndpointRef = uint32

// MIDIEntityRef is an entity that a device owns and that contains endpoints.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIEntityRef
type MIDIEntityRef = uint32

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIEventVisitor
type MIDIEventVisitor = func(unsafe.Pointer, uint64, MIDIUniversalMessage)

// MIDIMessage_32 is a 32-bit MIDI message.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIMessage_32
type MIDIMessage_32 = uint32

// MIDINotifyBlock is a callback block for notifying clients of state changes.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINotifyBlock
type MIDINotifyBlock = func(message *MIDINotification)

// MIDINotifyProc is a callback function for notifying clients of state changes.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDINotifyProc
type MIDINotifyProc = func(message uintptr, refCon unsafe.Pointer)

// MIDIObjectRef is the common base class for many of the framework’s objects.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIObjectRef
type MIDIObjectRef = uint32

// MIDIPortRef is a MIDI connection that a client maintains.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIPortRef
type MIDIPortRef = uint32

// MIDIReceiveBlock is a block receiving MIDI input that includes the incoming messages and a refCon to identify the source.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIReceiveBlock
type MIDIReceiveBlock = func(evtlist *MIDIEventList, srcConnRefCon unsafe.Pointer)

// MIDISetupRef is a type that represents the global state of the MIDI system, that contains lists of the devices and serial port owners.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDISetupRef
type MIDISetupRef = uint32

// MIDIThruConnectionRef is an opaque reference to a play-through connection.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIThruConnectionRef
type MIDIThruConnectionRef = uint32

// MIDITimeStamp is the time on the host clock when the event occurred.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDITimeStamp
type MIDITimeStamp = uint64

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUInteger14
type MIDIUInteger14 = uint16

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUInteger2
type MIDIUInteger2 = uint8

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUInteger28
type MIDIUInteger28 = uint32

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUInteger4
type MIDIUInteger4 = uint8

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUInteger7
type MIDIUInteger7 = uint8

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpointManager/DictionaryKey
type MIDIUMPEndpointManagerDictionaryKey = string

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlockID
type MIDIUMPFunctionBlockID = uint8

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPGroupNumber
type MIDIUMPGroupNumber = uint8

// MIDIUniqueID is a MIDI object’s unique identifier.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUniqueID
type MIDIUniqueID = int32

// MIDIMessage32 is a Go-name alias for MIDIMessage_32.
type MIDIMessage32 = MIDIMessage_32
