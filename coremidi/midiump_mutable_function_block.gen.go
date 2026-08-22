// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MIDIUMPMutableFunctionBlock] class.
var (
	_MIDIUMPMutableFunctionBlockClass     MIDIUMPMutableFunctionBlockClass
	_MIDIUMPMutableFunctionBlockClassOnce sync.Once
)

func getMIDIUMPMutableFunctionBlockClass() MIDIUMPMutableFunctionBlockClass {
	_MIDIUMPMutableFunctionBlockClassOnce.Do(func() {
		_MIDIUMPMutableFunctionBlockClass = MIDIUMPMutableFunctionBlockClass{class: objc.GetClass("MIDIUMPMutableFunctionBlock")}
	})
	return _MIDIUMPMutableFunctionBlockClass
}

// GetMIDIUMPMutableFunctionBlockClass returns the class object for MIDIUMPMutableFunctionBlock.
func GetMIDIUMPMutableFunctionBlockClass() MIDIUMPMutableFunctionBlockClass {
	return getMIDIUMPMutableFunctionBlockClass()
}

type MIDIUMPMutableFunctionBlockClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MIDIUMPMutableFunctionBlockClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MIDIUMPMutableFunctionBlockClass) Alloc() MIDIUMPMutableFunctionBlock {
	rv := objc.Send[MIDIUMPMutableFunctionBlock](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MIDIUMPMutableFunctionBlock.InitWithNameDirectionFirstGroupTotalGroupsSpannedMaxSysEx8StreamsMIDI1InfoUIHintIsEnabled]
//
// # Instance Methods
//
//   - [MIDIUMPMutableFunctionBlock.ReconfigureWithFirstGroupDirectionMIDI1InfoUIHintError]
//   - [MIDIUMPMutableFunctionBlock.SetEnabledError]
//   - [MIDIUMPMutableFunctionBlock.SetNameError]
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableFunctionBlock
type MIDIUMPMutableFunctionBlock struct {
	MIDIUMPFunctionBlock
}

// MIDIUMPMutableFunctionBlockFromID constructs a [MIDIUMPMutableFunctionBlock] from an objc.ID.
func MIDIUMPMutableFunctionBlockFromID(id objc.ID) MIDIUMPMutableFunctionBlock {
	return MIDIUMPMutableFunctionBlock{MIDIUMPFunctionBlock: MIDIUMPFunctionBlockFromID(id)}
}

// NOTE: MIDIUMPMutableFunctionBlock adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MIDIUMPMutableFunctionBlock] class.
//
// # Initializers
//
//   - [IMIDIUMPMutableFunctionBlock.InitWithNameDirectionFirstGroupTotalGroupsSpannedMaxSysEx8StreamsMIDI1InfoUIHintIsEnabled]
//
// # Instance Methods
//
//   - [IMIDIUMPMutableFunctionBlock.ReconfigureWithFirstGroupDirectionMIDI1InfoUIHintError]
//   - [IMIDIUMPMutableFunctionBlock.SetEnabledError]
//   - [IMIDIUMPMutableFunctionBlock.SetNameError]
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableFunctionBlock
type IMIDIUMPMutableFunctionBlock interface {
	IMIDIUMPFunctionBlock

	// Topic: Initializers

	InitWithNameDirectionFirstGroupTotalGroupsSpannedMaxSysEx8StreamsMIDI1InfoUIHintIsEnabled(name string, direction MIDIUMPFunctionBlockDirection, firstGroup MIDIUMPGroupNumber, totalGroupsSpanned MIDIUInteger7, maxSysEx8Streams MIDIUInteger7, MIDI1Info MIDIUMPFunctionBlockMIDI1Info, UIHint MIDIUMPFunctionBlockUIHint, isEnabled bool) MIDIUMPMutableFunctionBlock

	// Topic: Instance Methods

	ReconfigureWithFirstGroupDirectionMIDI1InfoUIHintError(firstGroup MIDIUMPGroupNumber, direction MIDIUMPFunctionBlockDirection, MIDI1Info MIDIUMPFunctionBlockMIDI1Info, UIHint MIDIUMPFunctionBlockUIHint) (bool, error)
	SetEnabledError(isEnabled bool) (bool, error)
	SetNameError(name string) (bool, error)
}

// Init initializes the instance.
func (m MIDIUMPMutableFunctionBlock) Init() MIDIUMPMutableFunctionBlock {
	rv := objc.Send[MIDIUMPMutableFunctionBlock](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MIDIUMPMutableFunctionBlock) Autorelease() MIDIUMPMutableFunctionBlock {
	rv := objc.Send[MIDIUMPMutableFunctionBlock](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIUMPMutableFunctionBlock creates a new MIDIUMPMutableFunctionBlock instance.
func NewMIDIUMPMutableFunctionBlock() MIDIUMPMutableFunctionBlock {
	class := getMIDIUMPMutableFunctionBlockClass()
	rv := objc.Send[MIDIUMPMutableFunctionBlock](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableFunctionBlock/init(name:direction:firstGroup:totalGroupsSpanned:maxSysEx8Streams:midi1Info:uiHint:isEnabled:)
func NewMIDIUMPMutableFunctionBlockWithNameDirectionFirstGroupTotalGroupsSpannedMaxSysEx8StreamsMIDI1InfoUIHintIsEnabled(name string, direction MIDIUMPFunctionBlockDirection, firstGroup MIDIUMPGroupNumber, totalGroupsSpanned MIDIUInteger7, maxSysEx8Streams MIDIUInteger7, MIDI1Info MIDIUMPFunctionBlockMIDI1Info, UIHint MIDIUMPFunctionBlockUIHint, isEnabled bool) MIDIUMPMutableFunctionBlock {
	instance := getMIDIUMPMutableFunctionBlockClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:direction:firstGroup:totalGroupsSpanned:maxSysEx8Streams:MIDI1Info:UIHint:isEnabled:"), objc.String(name), direction, firstGroup, totalGroupsSpanned, maxSysEx8Streams, MIDI1Info, UIHint, isEnabled)
	return MIDIUMPMutableFunctionBlockFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableFunctionBlock/init(name:direction:firstGroup:totalGroupsSpanned:maxSysEx8Streams:midi1Info:uiHint:isEnabled:)
func (m MIDIUMPMutableFunctionBlock) InitWithNameDirectionFirstGroupTotalGroupsSpannedMaxSysEx8StreamsMIDI1InfoUIHintIsEnabled(name string, direction MIDIUMPFunctionBlockDirection, firstGroup MIDIUMPGroupNumber, totalGroupsSpanned MIDIUInteger7, maxSysEx8Streams MIDIUInteger7, MIDI1Info MIDIUMPFunctionBlockMIDI1Info, UIHint MIDIUMPFunctionBlockUIHint, isEnabled bool) MIDIUMPMutableFunctionBlock {
	rv := objc.Send[MIDIUMPMutableFunctionBlock](m.ID, objc.Sel("initWithName:direction:firstGroup:totalGroupsSpanned:maxSysEx8Streams:MIDI1Info:UIHint:isEnabled:"), objc.String(name), direction, firstGroup, totalGroupsSpanned, maxSysEx8Streams, MIDI1Info, UIHint, isEnabled)
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableFunctionBlock/reconfigure(firstGroup:direction:MIDI1Info:UIHint:)
func (m MIDIUMPMutableFunctionBlock) ReconfigureWithFirstGroupDirectionMIDI1InfoUIHintError(firstGroup MIDIUMPGroupNumber, direction MIDIUMPFunctionBlockDirection, MIDI1Info MIDIUMPFunctionBlockMIDI1Info, UIHint MIDIUMPFunctionBlockUIHint) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("reconfigureWithFirstGroup:direction:MIDI1Info:UIHint:error:"), firstGroup, direction, MIDI1Info, UIHint, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("reconfigureWithFirstGroup:direction:MIDI1Info:UIHint:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableFunctionBlock/setEnabled(_:)
func (m MIDIUMPMutableFunctionBlock) SetEnabledError(isEnabled bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("setEnabled:error:"), isEnabled, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setEnabled:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPMutableFunctionBlock/setName(_:)
func (m MIDIUMPMutableFunctionBlock) SetNameError(name string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("setName:error:"), objc.String(name), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setName:error: returned NO with nil NSError")
	}
	return rv, nil

}
