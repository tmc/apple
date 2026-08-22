// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MIDICIDeviceInfo] class.
var (
	_MIDICIDeviceInfoClass     MIDICIDeviceInfoClass
	_MIDICIDeviceInfoClassOnce sync.Once
)

func getMIDICIDeviceInfoClass() MIDICIDeviceInfoClass {
	_MIDICIDeviceInfoClassOnce.Do(func() {
		_MIDICIDeviceInfoClass = MIDICIDeviceInfoClass{class: objc.GetClass("MIDICIDeviceInfo")}
	})
	return _MIDICIDeviceInfoClass
}

// GetMIDICIDeviceInfoClass returns the class object for MIDICIDeviceInfo.
func GetMIDICIDeviceInfoClass() MIDICIDeviceInfoClass {
	return getMIDICIDeviceInfoClass()
}

type MIDICIDeviceInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MIDICIDeviceInfoClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MIDICIDeviceInfoClass) Alloc() MIDICIDeviceInfo {
	rv := objc.Send[MIDICIDeviceInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides basic information about a MIDI-CI device.
//
// # Inspecting a Device
//
//   - [MIDICIDeviceInfo.ManufacturerID]: The MIDI System Exclusive (SysEx) ID of the device manufacturer.
//   - [MIDICIDeviceInfo.Family]: The family to which the device belongs.
//   - [MIDICIDeviceInfo.ModelNumber]: The model number of the device.
//   - [MIDICIDeviceInfo.RevisionLevel]: The revision number of the device model number.
//   - [MIDICIDeviceInfo.MidiDestination]: The MIDI destination the device’s MIDI entity uses for capability inquiries.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceInfo
type MIDICIDeviceInfo struct {
	objectivec.Object
}

// MIDICIDeviceInfoFromID constructs a [MIDICIDeviceInfo] from an objc.ID.
//
// An object that provides basic information about a MIDI-CI device.
func MIDICIDeviceInfoFromID(id objc.ID) MIDICIDeviceInfo {
	return MIDICIDeviceInfo{objectivec.Object{ID: id}}
}

// NOTE: MIDICIDeviceInfo adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MIDICIDeviceInfo] class.
//
// # Inspecting a Device
//
//   - [IMIDICIDeviceInfo.ManufacturerID]: The MIDI System Exclusive (SysEx) ID of the device manufacturer.
//   - [IMIDICIDeviceInfo.Family]: The family to which the device belongs.
//   - [IMIDICIDeviceInfo.ModelNumber]: The model number of the device.
//   - [IMIDICIDeviceInfo.RevisionLevel]: The revision number of the device model number.
//   - [IMIDICIDeviceInfo.MidiDestination]: The MIDI destination the device’s MIDI entity uses for capability inquiries.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceInfo
type IMIDICIDeviceInfo interface {
	objectivec.IObject

	// Topic: Inspecting a Device

	// The MIDI System Exclusive (SysEx) ID of the device manufacturer.
	ManufacturerID() foundation.NSData
	// The family to which the device belongs.
	Family() foundation.NSData
	// The model number of the device.
	ModelNumber() foundation.NSData
	// The revision number of the device model number.
	RevisionLevel() foundation.NSData
	// The MIDI destination the device’s MIDI entity uses for capability inquiries.
	MidiDestination() MIDIEndpointRef

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (m MIDICIDeviceInfo) Init() MIDICIDeviceInfo {
	rv := objc.Send[MIDICIDeviceInfo](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MIDICIDeviceInfo) Autorelease() MIDICIDeviceInfo {
	rv := objc.Send[MIDICIDeviceInfo](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIDeviceInfo creates a new MIDICIDeviceInfo instance.
func NewMIDICIDeviceInfo() MIDICIDeviceInfo {
	class := getMIDICIDeviceInfoClass()
	rv := objc.Send[MIDICIDeviceInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new device information instance.
//
// midiDestination: The MIDI destination to use for capability inquiry.
//
// manufacturer: The device manufacturer.
//
// family: The family to which this device belongs.
//
// modelNumber: The device’s model number.
//
// revisionLevel: The version of the device’s model number.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceInfo/init(destination:manufacturer:family:model:revision:)
func NewMIDICIDeviceInfoWithDestinationManufacturerFamilyModelRevision(midiDestination MIDIEntityRef, manufacturer foundation.NSData, family foundation.NSData, modelNumber foundation.NSData, revisionLevel foundation.NSData) MIDICIDeviceInfo {
	instance := getMIDICIDeviceInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDestination:manufacturer:family:model:revision:"), midiDestination, manufacturer, family, modelNumber, revisionLevel)
	return MIDICIDeviceInfoFromID(rv)
}

func (m MIDICIDeviceInfo) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The MIDI System Exclusive (SysEx) ID of the device manufacturer.
//
// # Discussion
//
// This value is 3 bytes long.
//
// The framework pads single-byte System Exclusive (SysEx) IDs with trailing
// zeros. For example, Apple’s SysEx ID, 0x11, is `0x110000`.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceInfo/manufacturerID
func (m MIDICIDeviceInfo) ManufacturerID() foundation.NSData {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("manufacturerID"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// The family to which the device belongs.
//
// # Discussion
//
// This value is 2 bytes long.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceInfo/family
func (m MIDICIDeviceInfo) Family() foundation.NSData {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("family"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// The model number of the device.
//
// # Discussion
//
// This value is 2 bytes long.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceInfo/modelNumber
func (m MIDICIDeviceInfo) ModelNumber() foundation.NSData {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelNumber"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// The revision number of the device model number.
//
// # Discussion
//
// This value is 2 bytes long.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceInfo/revisionLevel
func (m MIDICIDeviceInfo) RevisionLevel() foundation.NSData {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("revisionLevel"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// The MIDI destination the device’s MIDI entity uses for capability
// inquiries.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceInfo/midiDestination
func (m MIDICIDeviceInfo) MidiDestination() MIDIEndpointRef {
	rv := objc.Send[MIDIEndpointRef](m.ID, objc.Sel("midiDestination"))
	return MIDIEndpointRef(rv)
}
