// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MIDICIResponder] class.
var (
	_MIDICIResponderClass     MIDICIResponderClass
	_MIDICIResponderClassOnce sync.Once
)

func getMIDICIResponderClass() MIDICIResponderClass {
	_MIDICIResponderClassOnce.Do(func() {
		_MIDICIResponderClass = MIDICIResponderClass{class: objc.GetClass("MIDICIResponder")}
	})
	return _MIDICIResponderClass
}

// GetMIDICIResponderClass returns the class object for MIDICIResponder.
func GetMIDICIResponderClass() MIDICIResponderClass {
	return getMIDICIResponderClass()
}

type MIDICIResponderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MIDICIResponderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MIDICIResponderClass) Alloc() MIDICIResponder {
	rv := objc.Send[MIDICIResponder](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that responds to MIDI-CI inquiries from an initiator on behalf of
// a MIDI client, and handles profile and property exchange operations.
//
// # Setting a Responder Delegate
//
//   - [MIDICIResponder.ProfileDelegate]: The profile delegate.
//
// # Inspecting a Responder
//
//   - [MIDICIResponder.Initiators]: An array of initiators.
//   - [MIDICIResponder.DeviceInfo]: The MIDI-CI device’s information.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIResponder
type MIDICIResponder struct {
	objectivec.Object
}

// MIDICIResponderFromID constructs a [MIDICIResponder] from an objc.ID.
//
// An object that responds to MIDI-CI inquiries from an initiator on behalf of
// a MIDI client, and handles profile and property exchange operations.
func MIDICIResponderFromID(id objc.ID) MIDICIResponder {
	return MIDICIResponder{objectivec.Object{ID: id}}
}

// NOTE: MIDICIResponder adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MIDICIResponder] class.
//
// # Setting a Responder Delegate
//
//   - [IMIDICIResponder.ProfileDelegate]: The profile delegate.
//
// # Inspecting a Responder
//
//   - [IMIDICIResponder.Initiators]: An array of initiators.
//   - [IMIDICIResponder.DeviceInfo]: The MIDI-CI device’s information.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIResponder
type IMIDICIResponder interface {
	objectivec.IObject

	// Topic: Setting a Responder Delegate

	// The profile delegate.
	ProfileDelegate() MIDICIProfileResponderDelegate

	// Topic: Inspecting a Responder

	// An array of initiators.
	Initiators() []foundation.NSNumber
	// The MIDI-CI device’s information.
	DeviceInfo() IMIDICIDeviceInfo
}

// Init initializes the instance.
func (m MIDICIResponder) Init() MIDICIResponder {
	rv := objc.Send[MIDICIResponder](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MIDICIResponder) Autorelease() MIDICIResponder {
	rv := objc.Send[MIDICIResponder](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIResponder creates a new MIDICIResponder instance.
func NewMIDICIResponder() MIDICIResponder {
	class := getMIDICIResponderClass()
	rv := objc.Send[MIDICIResponder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new responder.
//
// deviceInfo: The MIDI-CI device information.
//
// delegate: The responder’s delegate object.
//
// profileList: The list of profile state objects.
//
// propertiesSupported: A Boolean value that indicates whether the responder supports properties.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIResponder/init(deviceInfo:profileDelegate:profileStates:supportProperties:)
func NewMIDICIResponderWithDeviceInfoProfileDelegateProfileStatesSupportProperties(deviceInfo IMIDICIDeviceInfo, delegate MIDICIProfileResponderDelegate, profileList MIDICIProfileStateList, propertiesSupported bool) MIDICIResponder {
	instance := getMIDICIResponderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDeviceInfo:profileDelegate:profileStates:supportProperties:"), deviceInfo, delegate, profileList, propertiesSupported)
	return MIDICIResponderFromID(rv)
}

// The profile delegate.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIResponder/profileDelegate
func (m MIDICIResponder) ProfileDelegate() MIDICIProfileResponderDelegate {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("profileDelegate"))
	return MIDICIProfileResponderDelegateObjectFromID(rv)
}

// An array of initiators.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIResponder/initiators
func (m MIDICIResponder) Initiators() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("initiators"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// The MIDI-CI device’s information.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIResponder/deviceInfo
func (m MIDICIResponder) DeviceInfo() IMIDICIDeviceInfo {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("deviceInfo"))
	return MIDICIDeviceInfoFromID(objc.ID(rv))
}
