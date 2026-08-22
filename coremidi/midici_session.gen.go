// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MIDICISession] class.
var (
	_MIDICISessionClass     MIDICISessionClass
	_MIDICISessionClassOnce sync.Once
)

func getMIDICISessionClass() MIDICISessionClass {
	_MIDICISessionClassOnce.Do(func() {
		_MIDICISessionClass = MIDICISessionClass{class: objc.GetClass("MIDICISession")}
	})
	return _MIDICISessionClass
}

// GetMIDICISessionClass returns the class object for MIDICISession.
func GetMIDICISessionClass() MIDICISessionClass {
	return getMIDICISessionClass()
}

type MIDICISessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MIDICISessionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MIDICISessionClass) Alloc() MIDICISession {
	rv := objc.Send[MIDICISession](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a MIDI-CI session.
//
// # Overview
//
// A MIDI-CI session is a bidirectional communication path between a MIDI
// source and destination identified using MIDI-CI discovery. Use a session to
// manipulate MIDI-CI profiles and to discover device capabilities.
//
// # Configuring a Session
//
//   - [MIDICISession.ProfileChangedCallback]: An optional block the system calls after it enables or disables a profile.
//   - [MIDICISession.SetProfileChangedCallback]
//   - [MIDICISession.ProfileSpecificDataHandler]: An optional block the system calls when a device sends profile-specific data to the session.
//   - [MIDICISession.SetProfileSpecificDataHandler]
//
// # Inspecting a Session
//
//   - [MIDICISession.DeviceInfo]: Information about a MIDI-CI device.
//   - [MIDICISession.MaxSysExSize]: The maximum size of System Exclusive (SysEx) messages.
//   - [MIDICISession.MidiDestination]: The MIDI destination with which the session is communicating.
//   - [MIDICISession.MaxPropertyRequests]: The maximum number of simultaneous property exchange requests, if supported.
//   - [MIDICISession.SupportsProfileCapability]: A Boolean value that indicates whether the entity supports the MIDI-CI profile’s capability.
//   - [MIDICISession.SupportsPropertyCapability]: A Boolean value that indicates whether the entity supports the MIDI-CI property exchange capability.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICISession
type MIDICISession struct {
	objectivec.Object
}

// MIDICISessionFromID constructs a [MIDICISession] from an objc.ID.
//
// An object that represents a MIDI-CI session.
func MIDICISessionFromID(id objc.ID) MIDICISession {
	return MIDICISession{objectivec.Object{ID: id}}
}

// NOTE: MIDICISession adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MIDICISession] class.
//
// # Configuring a Session
//
//   - [IMIDICISession.ProfileChangedCallback]: An optional block the system calls after it enables or disables a profile.
//   - [IMIDICISession.SetProfileChangedCallback]
//   - [IMIDICISession.ProfileSpecificDataHandler]: An optional block the system calls when a device sends profile-specific data to the session.
//   - [IMIDICISession.SetProfileSpecificDataHandler]
//
// # Inspecting a Session
//
//   - [IMIDICISession.DeviceInfo]: Information about a MIDI-CI device.
//   - [IMIDICISession.MaxSysExSize]: The maximum size of System Exclusive (SysEx) messages.
//   - [IMIDICISession.MidiDestination]: The MIDI destination with which the session is communicating.
//   - [IMIDICISession.MaxPropertyRequests]: The maximum number of simultaneous property exchange requests, if supported.
//   - [IMIDICISession.SupportsProfileCapability]: A Boolean value that indicates whether the entity supports the MIDI-CI profile’s capability.
//   - [IMIDICISession.SupportsPropertyCapability]: A Boolean value that indicates whether the entity supports the MIDI-CI property exchange capability.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICISession
type IMIDICISession interface {
	objectivec.IObject

	// Topic: Configuring a Session

	// An optional block the system calls after it enables or disables a profile.
	ProfileChangedCallback() MIDICISessionUint32MIDICIProfileInt32Handler
	SetProfileChangedCallback(value MIDICISessionUint32MIDICIProfileInt32Handler)
	// An optional block the system calls when a device sends profile-specific data to the session.
	ProfileSpecificDataHandler() MIDICISessionUint32MIDICIProfileDataHandler
	SetProfileSpecificDataHandler(value MIDICISessionUint32MIDICIProfileDataHandler)

	// Topic: Inspecting a Session

	// Information about a MIDI-CI device.
	DeviceInfo() IMIDICIDeviceInfo
	// The maximum size of System Exclusive (SysEx) messages.
	MaxSysExSize() foundation.NSNumber
	// The MIDI destination with which the session is communicating.
	MidiDestination() MIDIEntityRef
	// The maximum number of simultaneous property exchange requests, if supported.
	MaxPropertyRequests() foundation.NSNumber
	// A Boolean value that indicates whether the entity supports the MIDI-CI profile’s capability.
	SupportsProfileCapability() bool
	// A Boolean value that indicates whether the entity supports the MIDI-CI property exchange capability.
	SupportsPropertyCapability() bool
}

// Init initializes the instance.
func (m MIDICISession) Init() MIDICISession {
	rv := objc.Send[MIDICISession](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MIDICISession) Autorelease() MIDICISession {
	rv := objc.Send[MIDICISession](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICISession creates a new MIDICISession instance.
func NewMIDICISession() MIDICISession {
	class := getMIDICISessionClass()
	rv := objc.Send[MIDICISession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An optional block the system calls after it enables or disables a profile.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/profileChangedCallback
func (m MIDICISession) ProfileChangedCallback() MIDICISessionUint32MIDICIProfileInt32Handler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("profileChangedCallback"))
	_ = rv
	return nil
}
func (m MIDICISession) SetProfileChangedCallback(value MIDICISessionUint32MIDICIProfileInt32Handler) {
	block, cleanup := NewMIDICISessionUint32MIDICIProfileInt32Block(value)
	defer cleanup()
	objc.Send[struct{}](m.ID, objc.Sel("setProfileChangedCallback:"), block)
}

// An optional block the system calls when a device sends profile-specific
// data to the session.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/profileSpecificDataHandler
func (m MIDICISession) ProfileSpecificDataHandler() MIDICISessionUint32MIDICIProfileDataHandler {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("profileSpecificDataHandler"))
	_ = rv
	return nil
}
func (m MIDICISession) SetProfileSpecificDataHandler(value MIDICISessionUint32MIDICIProfileDataHandler) {
	block, cleanup := NewMIDICISessionUint32MIDICIProfileDataBlock(value)
	defer cleanup()
	objc.Send[struct{}](m.ID, objc.Sel("setProfileSpecificDataHandler:"), block)
}

// Information about a MIDI-CI device.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/deviceInfo
func (m MIDICISession) DeviceInfo() IMIDICIDeviceInfo {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("deviceInfo"))
	return MIDICIDeviceInfoFromID(objc.ID(rv))
}

// The maximum size of System Exclusive (SysEx) messages.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/maxSysExSize
func (m MIDICISession) MaxSysExSize() foundation.NSNumber {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("maxSysExSize"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// The MIDI destination with which the session is communicating.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/midiDestination
func (m MIDICISession) MidiDestination() MIDIEntityRef {
	rv := objc.Send[MIDIEntityRef](m.ID, objc.Sel("midiDestination"))
	return MIDIEntityRef(rv)
}

// The maximum number of simultaneous property exchange requests, if
// supported.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/maxPropertyRequests
func (m MIDICISession) MaxPropertyRequests() foundation.NSNumber {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("maxPropertyRequests"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the entity supports the MIDI-CI
// profile’s capability.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/supportsProfileCapability
func (m MIDICISession) SupportsProfileCapability() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportsProfileCapability"))
	return rv
}

// A Boolean value that indicates whether the entity supports the MIDI-CI
// property exchange capability.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICISession/supportsPropertyCapability
func (m MIDICISession) SupportsPropertyCapability() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportsPropertyCapability"))
	return rv
}
