// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MIDIUMPFunctionBlock] class.
var (
	_MIDIUMPFunctionBlockClass     MIDIUMPFunctionBlockClass
	_MIDIUMPFunctionBlockClassOnce sync.Once
)

func getMIDIUMPFunctionBlockClass() MIDIUMPFunctionBlockClass {
	_MIDIUMPFunctionBlockClassOnce.Do(func() {
		_MIDIUMPFunctionBlockClass = MIDIUMPFunctionBlockClass{class: objc.GetClass("MIDIUMPFunctionBlock")}
	})
	return _MIDIUMPFunctionBlockClass
}

// GetMIDIUMPFunctionBlockClass returns the class object for MIDIUMPFunctionBlock.
func GetMIDIUMPFunctionBlockClass() MIDIUMPFunctionBlockClass {
	return getMIDIUMPFunctionBlockClass()
}

type MIDIUMPFunctionBlockClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MIDIUMPFunctionBlockClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MIDIUMPFunctionBlockClass) Alloc() MIDIUMPFunctionBlock {
	rv := objc.Send[MIDIUMPFunctionBlock](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MIDIUMPFunctionBlock.Direction]
//   - [MIDIUMPFunctionBlock.FirstGroup]
//   - [MIDIUMPFunctionBlock.FunctionBlockID]
//   - [MIDIUMPFunctionBlock.IsEnabled]
//   - [MIDIUMPFunctionBlock.MaxSysEx8Streams]
//   - [MIDIUMPFunctionBlock.MIDI1Info]
//   - [MIDIUMPFunctionBlock.MidiCIDevice]
//   - [MIDIUMPFunctionBlock.Name]
//   - [MIDIUMPFunctionBlock.TotalGroupsSpanned]
//   - [MIDIUMPFunctionBlock.UIHint]
//   - [MIDIUMPFunctionBlock.UMPEndpoint]
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock
type MIDIUMPFunctionBlock struct {
	objectivec.Object
}

// MIDIUMPFunctionBlockFromID constructs a [MIDIUMPFunctionBlock] from an objc.ID.
func MIDIUMPFunctionBlockFromID(id objc.ID) MIDIUMPFunctionBlock {
	return MIDIUMPFunctionBlock{objectivec.Object{ID: id}}
}

// NOTE: MIDIUMPFunctionBlock adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MIDIUMPFunctionBlock] class.
//
// # Instance Properties
//
//   - [IMIDIUMPFunctionBlock.Direction]
//   - [IMIDIUMPFunctionBlock.FirstGroup]
//   - [IMIDIUMPFunctionBlock.FunctionBlockID]
//   - [IMIDIUMPFunctionBlock.IsEnabled]
//   - [IMIDIUMPFunctionBlock.MaxSysEx8Streams]
//   - [IMIDIUMPFunctionBlock.MIDI1Info]
//   - [IMIDIUMPFunctionBlock.MidiCIDevice]
//   - [IMIDIUMPFunctionBlock.Name]
//   - [IMIDIUMPFunctionBlock.TotalGroupsSpanned]
//   - [IMIDIUMPFunctionBlock.UIHint]
//   - [IMIDIUMPFunctionBlock.UMPEndpoint]
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock
type IMIDIUMPFunctionBlock interface {
	objectivec.IObject

	// Topic: Instance Properties

	Direction() MIDIUMPFunctionBlockDirection
	FirstGroup() MIDIUMPGroupNumber
	FunctionBlockID() MIDIUMPFunctionBlockID
	IsEnabled() bool
	MaxSysEx8Streams() uint8
	MIDI1Info() MIDIUMPFunctionBlockMIDI1Info
	MidiCIDevice() IMIDICIDevice
	Name() string
	TotalGroupsSpanned() MIDIUInteger7
	UIHint() MIDIUMPFunctionBlockUIHint
	UMPEndpoint() IMIDIUMPEndpoint
}

// Init initializes the instance.
func (m MIDIUMPFunctionBlock) Init() MIDIUMPFunctionBlock {
	rv := objc.Send[MIDIUMPFunctionBlock](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MIDIUMPFunctionBlock) Autorelease() MIDIUMPFunctionBlock {
	rv := objc.Send[MIDIUMPFunctionBlock](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIUMPFunctionBlock creates a new MIDIUMPFunctionBlock instance.
func NewMIDIUMPFunctionBlock() MIDIUMPFunctionBlock {
	class := getMIDIUMPFunctionBlockClass()
	rv := objc.Send[MIDIUMPFunctionBlock](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/direction
func (m MIDIUMPFunctionBlock) Direction() MIDIUMPFunctionBlockDirection {
	rv := objc.Send[MIDIUMPFunctionBlockDirection](m.ID, objc.Sel("direction"))
	return MIDIUMPFunctionBlockDirection(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/firstGroup
func (m MIDIUMPFunctionBlock) FirstGroup() MIDIUMPGroupNumber {
	rv := objc.Send[MIDIUMPGroupNumber](m.ID, objc.Sel("firstGroup"))
	return MIDIUMPGroupNumber(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/functionBlockID
func (m MIDIUMPFunctionBlock) FunctionBlockID() MIDIUMPFunctionBlockID {
	rv := objc.Send[MIDIUMPFunctionBlockID](m.ID, objc.Sel("functionBlockID"))
	return MIDIUMPFunctionBlockID(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/isEnabled
func (m MIDIUMPFunctionBlock) IsEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isEnabled"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/maxSysEx8Streams
func (m MIDIUMPFunctionBlock) MaxSysEx8Streams() uint8 {
	rv := objc.Send[uint8](m.ID, objc.Sel("maxSysEx8Streams"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/midi1Info
func (m MIDIUMPFunctionBlock) MIDI1Info() MIDIUMPFunctionBlockMIDI1Info {
	rv := objc.Send[MIDIUMPFunctionBlockMIDI1Info](m.ID, objc.Sel("MIDI1Info"))
	return MIDIUMPFunctionBlockMIDI1Info(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/midiCIDevice
func (m MIDIUMPFunctionBlock) MidiCIDevice() IMIDICIDevice {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("midiCIDevice"))
	return MIDICIDeviceFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/name
func (m MIDIUMPFunctionBlock) Name() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/totalGroupsSpanned
func (m MIDIUMPFunctionBlock) TotalGroupsSpanned() MIDIUInteger7 {
	rv := objc.Send[MIDIUInteger7](m.ID, objc.Sel("totalGroupsSpanned"))
	return MIDIUInteger7(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/uiHint
func (m MIDIUMPFunctionBlock) UIHint() MIDIUMPFunctionBlockUIHint {
	rv := objc.Send[MIDIUMPFunctionBlockUIHint](m.ID, objc.Sel("UIHint"))
	return MIDIUMPFunctionBlockUIHint(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPFunctionBlock/umpEndpoint
func (m MIDIUMPFunctionBlock) UMPEndpoint() IMIDIUMPEndpoint {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("UMPEndpoint"))
	return MIDIUMPEndpointFromID(objc.ID(rv))
}
