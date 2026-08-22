// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MIDICIProfile] class.
var (
	_MIDICIProfileClass     MIDICIProfileClass
	_MIDICIProfileClassOnce sync.Once
)

func getMIDICIProfileClass() MIDICIProfileClass {
	_MIDICIProfileClassOnce.Do(func() {
		_MIDICIProfileClass = MIDICIProfileClass{class: objc.GetClass("MIDICIProfile")}
	})
	return _MIDICIProfileClass
}

// GetMIDICIProfileClass returns the class object for MIDICIProfile.
func GetMIDICIProfileClass() MIDICIProfileClass {
	return getMIDICIProfileClass()
}

type MIDICIProfileClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MIDICIProfileClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MIDICIProfileClass) Alloc() MIDICIProfile {
	rv := objc.Send[MIDICIProfile](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A mapping of MIDI messages to specific sounds and synthesis behaviors, such
// as General MIDI, a drawbar organ, and so on.
//
// # Creating a Profile
//
//   - [MIDICIProfile.InitWithData]: Creates a MIDI profile for the specified data.
//   - [MIDICIProfile.InitWithDataName]: Creates a named MIDI profile for the specified data.
//
// # Inspecting a Profile
//
//   - [MIDICIProfile.Name]: A string that describes the profile.
//   - [MIDICIProfile.ProfileID]: The unique five-byte profile identifier that represents the profile.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfile
type MIDICIProfile struct {
	objectivec.Object
}

// MIDICIProfileFromID constructs a [MIDICIProfile] from an objc.ID.
//
// A mapping of MIDI messages to specific sounds and synthesis behaviors, such
// as General MIDI, a drawbar organ, and so on.
func MIDICIProfileFromID(id objc.ID) MIDICIProfile {
	return MIDICIProfile{objectivec.Object{ID: id}}
}

// NOTE: MIDICIProfile adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MIDICIProfile] class.
//
// # Creating a Profile
//
//   - [IMIDICIProfile.InitWithData]: Creates a MIDI profile for the specified data.
//   - [IMIDICIProfile.InitWithDataName]: Creates a named MIDI profile for the specified data.
//
// # Inspecting a Profile
//
//   - [IMIDICIProfile.Name]: A string that describes the profile.
//   - [IMIDICIProfile.ProfileID]: The unique five-byte profile identifier that represents the profile.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfile
type IMIDICIProfile interface {
	objectivec.IObject

	// Topic: Creating a Profile

	// Creates a MIDI profile for the specified data.
	InitWithData(data foundation.NSData) MIDICIProfile
	// Creates a named MIDI profile for the specified data.
	InitWithDataName(data foundation.NSData, inName string) MIDICIProfile

	// Topic: Inspecting a Profile

	// A string that describes the profile.
	Name() string
	// The unique five-byte profile identifier that represents the profile.
	ProfileID() foundation.NSData

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (m MIDICIProfile) Init() MIDICIProfile {
	rv := objc.Send[MIDICIProfile](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MIDICIProfile) Autorelease() MIDICIProfile {
	rv := objc.Send[MIDICIProfile](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIProfile creates a new MIDICIProfile instance.
func NewMIDICIProfile() MIDICIProfile {
	class := getMIDICIProfileClass()
	rv := objc.Send[MIDICIProfile](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a MIDI profile for the specified data.
//
// data: The profile’s unique byte sequence.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfile/init(data:)
func NewMIDICIProfileWithData(data foundation.NSData) MIDICIProfile {
	instance := getMIDICIProfileClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), data)
	return MIDICIProfileFromID(rv)
}

// Creates a named MIDI profile for the specified data.
//
// data: The profile’s unique byte sequence.
//
// inName: The profile name.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfile/init(data:name:)
func NewMIDICIProfileWithDataName(data foundation.NSData, inName string) MIDICIProfile {
	instance := getMIDICIProfileClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:name:"), data, objc.String(inName))
	return MIDICIProfileFromID(rv)
}

// Creates a MIDI profile for the specified data.
//
// data: The profile’s unique byte sequence.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfile/init(data:)
func (m MIDICIProfile) InitWithData(data foundation.NSData) MIDICIProfile {
	rv := objc.Send[MIDICIProfile](m.ID, objc.Sel("initWithData:"), data)
	return rv
}

// Creates a named MIDI profile for the specified data.
//
// data: The profile’s unique byte sequence.
//
// inName: The profile name.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfile/init(data:name:)
func (m MIDICIProfile) InitWithDataName(data foundation.NSData, inName string) MIDICIProfile {
	rv := objc.Send[MIDICIProfile](m.ID, objc.Sel("initWithData:name:"), data, objc.String(inName))
	return rv
}
func (m MIDICIProfile) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A string that describes the profile.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfile/name
func (m MIDICIProfile) Name() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The unique five-byte profile identifier that represents the profile.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfile/profileID
func (m MIDICIProfile) ProfileID() foundation.NSData {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("profileID"))
	return foundation.NSDataFromID(objc.ID(rv))
}
