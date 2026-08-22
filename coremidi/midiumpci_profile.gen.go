// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MIDIUMPCIProfile] class.
var (
	_MIDIUMPCIProfileClass     MIDIUMPCIProfileClass
	_MIDIUMPCIProfileClassOnce sync.Once
)

func getMIDIUMPCIProfileClass() MIDIUMPCIProfileClass {
	_MIDIUMPCIProfileClassOnce.Do(func() {
		_MIDIUMPCIProfileClass = MIDIUMPCIProfileClass{class: objc.GetClass("MIDIUMPCIProfile")}
	})
	return _MIDIUMPCIProfileClass
}

// GetMIDIUMPCIProfileClass returns the class object for MIDIUMPCIProfile.
func GetMIDIUMPCIProfileClass() MIDIUMPCIProfileClass {
	return getMIDIUMPCIProfileClass()
}

type MIDIUMPCIProfileClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MIDIUMPCIProfileClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MIDIUMPCIProfileClass) Alloc() MIDIUMPCIProfile {
	rv := objc.Send[MIDIUMPCIProfile](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MIDIUMPCIProfile.EnabledChannelCount]
//   - [MIDIUMPCIProfile.FirstChannel]
//   - [MIDIUMPCIProfile.GroupOffset]
//   - [MIDIUMPCIProfile.IsEnabled]
//   - [MIDIUMPCIProfile.Name]
//   - [MIDIUMPCIProfile.ProfileID]
//   - [MIDIUMPCIProfile.ProfileType]
//   - [MIDIUMPCIProfile.TotalChannelCount]
//
// # Instance Methods
//
//   - [MIDIUMPCIProfile.SetProfileStateEnabledChannelCountError]
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile
type MIDIUMPCIProfile struct {
	objectivec.Object
}

// MIDIUMPCIProfileFromID constructs a [MIDIUMPCIProfile] from an objc.ID.
func MIDIUMPCIProfileFromID(id objc.ID) MIDIUMPCIProfile {
	return MIDIUMPCIProfile{objectivec.Object{ID: id}}
}

// NOTE: MIDIUMPCIProfile adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MIDIUMPCIProfile] class.
//
// # Instance Properties
//
//   - [IMIDIUMPCIProfile.EnabledChannelCount]
//   - [IMIDIUMPCIProfile.FirstChannel]
//   - [IMIDIUMPCIProfile.GroupOffset]
//   - [IMIDIUMPCIProfile.IsEnabled]
//   - [IMIDIUMPCIProfile.Name]
//   - [IMIDIUMPCIProfile.ProfileID]
//   - [IMIDIUMPCIProfile.ProfileType]
//   - [IMIDIUMPCIProfile.TotalChannelCount]
//
// # Instance Methods
//
//   - [IMIDIUMPCIProfile.SetProfileStateEnabledChannelCountError]
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile
type IMIDIUMPCIProfile interface {
	objectivec.IObject

	// Topic: Instance Properties

	EnabledChannelCount() MIDIUInteger14
	FirstChannel() MIDIChannelNumber
	GroupOffset() MIDIUMPGroupNumber
	IsEnabled() bool
	Name() string
	ProfileID() MIDICIProfileID
	ProfileType() MIDICIProfileType
	TotalChannelCount() MIDIUInteger14

	// Topic: Instance Methods

	SetProfileStateEnabledChannelCountError(isEnabled bool, enabledChannelCount MIDIUInteger14) (bool, error)
}

// Init initializes the instance.
func (m MIDIUMPCIProfile) Init() MIDIUMPCIProfile {
	rv := objc.Send[MIDIUMPCIProfile](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MIDIUMPCIProfile) Autorelease() MIDIUMPCIProfile {
	rv := objc.Send[MIDIUMPCIProfile](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIUMPCIProfile creates a new MIDIUMPCIProfile instance.
func NewMIDIUMPCIProfile() MIDIUMPCIProfile {
	class := getMIDIUMPCIProfileClass()
	rv := objc.Send[MIDIUMPCIProfile](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/setProfileState(_:enabledChannelCount:)
func (m MIDIUMPCIProfile) SetProfileStateEnabledChannelCountError(isEnabled bool, enabledChannelCount MIDIUInteger14) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("setProfileState:enabledChannelCount:error:"), isEnabled, enabledChannelCount, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setProfileState:enabledChannelCount:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/enabledChannelCount
func (m MIDIUMPCIProfile) EnabledChannelCount() MIDIUInteger14 {
	rv := objc.Send[MIDIUInteger14](m.ID, objc.Sel("enabledChannelCount"))
	return MIDIUInteger14(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/firstChannel
func (m MIDIUMPCIProfile) FirstChannel() MIDIChannelNumber {
	rv := objc.Send[MIDIChannelNumber](m.ID, objc.Sel("firstChannel"))
	return MIDIChannelNumber(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/groupOffset
func (m MIDIUMPCIProfile) GroupOffset() MIDIUMPGroupNumber {
	rv := objc.Send[MIDIUMPGroupNumber](m.ID, objc.Sel("groupOffset"))
	return MIDIUMPGroupNumber(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/isEnabled
func (m MIDIUMPCIProfile) IsEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isEnabled"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/name
func (m MIDIUMPCIProfile) Name() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/profileID
func (m MIDIUMPCIProfile) ProfileID() MIDICIProfileID {
	rv := objc.Send[MIDICIProfileID](m.ID, objc.Sel("profileID"))
	return MIDICIProfileID(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/profileType
func (m MIDIUMPCIProfile) ProfileType() MIDICIProfileType {
	rv := objc.Send[MIDICIProfileType](m.ID, objc.Sel("profileType"))
	return MIDICIProfileType(rv)
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/totalChannelCount
func (m MIDIUMPCIProfile) TotalChannelCount() MIDIUInteger14 {
	rv := objc.Send[MIDIUInteger14](m.ID, objc.Sel("totalChannelCount"))
	return MIDIUInteger14(rv)
}
