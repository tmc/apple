// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MIDICIProfileState] class.
var (
	_MIDICIProfileStateClass     MIDICIProfileStateClass
	_MIDICIProfileStateClassOnce sync.Once
)

func getMIDICIProfileStateClass() MIDICIProfileStateClass {
	_MIDICIProfileStateClassOnce.Do(func() {
		_MIDICIProfileStateClass = MIDICIProfileStateClass{class: objc.GetClass("MIDICIProfileState")}
	})
	return _MIDICIProfileStateClass
}

// GetMIDICIProfileStateClass returns the class object for MIDICIProfileState.
func GetMIDICIProfileStateClass() MIDICIProfileStateClass {
	return getMIDICIProfileStateClass()
}

type MIDICIProfileStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MIDICIProfileStateClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MIDICIProfileStateClass) Alloc() MIDICIProfileState {
	rv := objc.Send[MIDICIProfileState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides the enabled and disabled profiles for a MIDI
// channel or port on a device.
//
// # Creating a Profile State
//
//   - [MIDICIProfileState.InitWithEnabledProfilesDisabledProfiles]: Creates a new profile state object for the specified profiles.
//
// # Accessing the MIDI Channel
//
//   - [MIDICIProfileState.MidiChannel]: The MIDI channel to which this state applies.
//
// # Accessing Profiles
//
//   - [MIDICIProfileState.EnabledProfiles]: The object’s enabled profiles.
//   - [MIDICIProfileState.DisabledProfiles]: The object’s disabled profiles.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileState
type MIDICIProfileState struct {
	objectivec.Object
}

// MIDICIProfileStateFromID constructs a [MIDICIProfileState] from an objc.ID.
//
// An object that provides the enabled and disabled profiles for a MIDI
// channel or port on a device.
func MIDICIProfileStateFromID(id objc.ID) MIDICIProfileState {
	return MIDICIProfileState{objectivec.Object{ID: id}}
}

// NOTE: MIDICIProfileState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MIDICIProfileState] class.
//
// # Creating a Profile State
//
//   - [IMIDICIProfileState.InitWithEnabledProfilesDisabledProfiles]: Creates a new profile state object for the specified profiles.
//
// # Accessing the MIDI Channel
//
//   - [IMIDICIProfileState.MidiChannel]: The MIDI channel to which this state applies.
//
// # Accessing Profiles
//
//   - [IMIDICIProfileState.EnabledProfiles]: The object’s enabled profiles.
//   - [IMIDICIProfileState.DisabledProfiles]: The object’s disabled profiles.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileState
type IMIDICIProfileState interface {
	objectivec.IObject

	// Topic: Creating a Profile State

	// Creates a new profile state object for the specified profiles.
	InitWithEnabledProfilesDisabledProfiles(enabled []MIDICIProfile, disabled []MIDICIProfile) MIDICIProfileState

	// Topic: Accessing the MIDI Channel

	// The MIDI channel to which this state applies.
	MidiChannel() MIDIChannelNumber

	// Topic: Accessing Profiles

	// The object’s enabled profiles.
	EnabledProfiles() []MIDICIProfile
	// The object’s disabled profiles.
	DisabledProfiles() []MIDICIProfile

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (m MIDICIProfileState) Init() MIDICIProfileState {
	rv := objc.Send[MIDICIProfileState](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MIDICIProfileState) Autorelease() MIDICIProfileState {
	rv := objc.Send[MIDICIProfileState](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIProfileState creates a new MIDICIProfileState instance.
func NewMIDICIProfileState() MIDICIProfileState {
	class := getMIDICIProfileStateClass()
	rv := objc.Send[MIDICIProfileState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new profile state object for the specified MIDI channel and
// profiles.
//
// midiChannelNum: The MIDI channel.
//
// enabled: The enabled MIDI-CI profles.
//
// disabled: The disabled MIDI-CI profles.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileState/init(channel:enabledProfiles:disabledProfiles:)
func NewMIDICIProfileStateWithChannelEnabledProfilesDisabledProfiles(midiChannelNum MIDIChannelNumber, enabled []MIDICIProfile, disabled []MIDICIProfile) MIDICIProfileState {
	instance := getMIDICIProfileStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithChannel:enabledProfiles:disabledProfiles:"), midiChannelNum, objectivec.IObjectSliceToNSArray(enabled), objectivec.IObjectSliceToNSArray(disabled))
	return MIDICIProfileStateFromID(rv)
}

// Creates a new profile state object for the specified profiles.
//
// enabled: The enabled MIDI-CI profiles.
//
// disabled: The disabled MIDI-CI profiles.
//
// # Return Value
//
// A new profile state instance.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileState/init(enabledProfiles:disabledProfiles:)
func NewMIDICIProfileStateWithEnabledProfilesDisabledProfiles(enabled []MIDICIProfile, disabled []MIDICIProfile) MIDICIProfileState {
	instance := getMIDICIProfileStateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEnabledProfiles:disabledProfiles:"), objectivec.IObjectSliceToNSArray(enabled), objectivec.IObjectSliceToNSArray(disabled))
	return MIDICIProfileStateFromID(rv)
}

// Creates a new profile state object for the specified profiles.
//
// enabled: The enabled MIDI-CI profiles.
//
// disabled: The disabled MIDI-CI profiles.
//
// # Return Value
//
// A new profile state instance.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileState/init(enabledProfiles:disabledProfiles:)
func (m MIDICIProfileState) InitWithEnabledProfilesDisabledProfiles(enabled []MIDICIProfile, disabled []MIDICIProfile) MIDICIProfileState {
	rv := objc.Send[MIDICIProfileState](m.ID, objc.Sel("initWithEnabledProfiles:disabledProfiles:"), objectivec.IObjectSliceToNSArray(enabled), objectivec.IObjectSliceToNSArray(disabled))
	return rv
}
func (m MIDICIProfileState) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The MIDI channel to which this state applies.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileState/midiChannel
func (m MIDICIProfileState) MidiChannel() MIDIChannelNumber {
	rv := objc.Send[MIDIChannelNumber](m.ID, objc.Sel("midiChannel"))
	return MIDIChannelNumber(rv)
}

// The object’s enabled profiles.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileState/enabledProfiles
func (m MIDICIProfileState) EnabledProfiles() []MIDICIProfile {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("enabledProfiles"))
	return objc.ConvertSlice(rv, func(id objc.ID) MIDICIProfile {
		return MIDICIProfileFromID(id)
	})
}

// The object’s disabled profiles.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileState/disabledProfiles
func (m MIDICIProfileState) DisabledProfiles() []MIDICIProfile {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("disabledProfiles"))
	return objc.ConvertSlice(rv, func(id objc.ID) MIDICIProfile {
		return MIDICIProfileFromID(id)
	})
}
