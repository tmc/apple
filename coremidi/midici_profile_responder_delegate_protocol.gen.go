// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A protocol that defines the methods to respond to MIDI-CI responder life-cycle events.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileResponderDelegate
type MIDICIProfileResponderDelegate interface {
	objectivec.IObject

	// Enables a MIDI-CI initiator to create a session or reject the connection attempt.
	//
	// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileResponderDelegate/connectInitiator(_:with:)
	ConnectInitiatorWithDeviceInfo(initiatorMUID MIDICIInitiatiorMUID, deviceInfo IMIDICIDeviceInfo) bool
}

// MIDICIProfileResponderDelegateObject wraps an existing Objective-C object that conforms to the MIDICIProfileResponderDelegate protocol.
type MIDICIProfileResponderDelegateObject struct {
	objectivec.Object
}

func (o MIDICIProfileResponderDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// MIDICIProfileResponderDelegateObjectFromID constructs a [MIDICIProfileResponderDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MIDICIProfileResponderDelegateObjectFromID(id objc.ID) MIDICIProfileResponderDelegateObject {
	return MIDICIProfileResponderDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Enables a MIDI-CI initiator to create a session or reject the connection
// attempt.
//
// initiatorMUID: The ID of the MIDI-CI initiator.
//
// deviceInfo: The information that describes a device.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileResponderDelegate/connectInitiator(_:with:)
func (o MIDICIProfileResponderDelegateObject) ConnectInitiatorWithDeviceInfo(initiatorMUID MIDICIInitiatiorMUID, deviceInfo IMIDICIDeviceInfo) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("connectInitiator:withDeviceInfo:"), initiatorMUID, deviceInfo)
	return rv
}

// Processes MIDI data for a profile and channel.
//
// aProfile: The MIDI-CI profile.
//
// channel: The MIDI channel.
//
// inData: The data to process.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileResponderDelegate/handleData(for:onChannel:data:)
func (o MIDICIProfileResponderDelegateObject) HandleDataForProfileOnChannelData(aProfile IMIDICIProfile, channel MIDIChannelNumber, inData foundation.NSData) {
	objc.Send[struct{}](o.ID, objc.Sel("handleDataForProfile:onChannel:data:"), aProfile, channel, inData)
}

// Provides an opportunity to perform an action before the system sets the
// profile.
//
// aProfile: The profile the system uses to configure the device.
//
// channel: The MIDI channel assignment.
//
// shouldEnable: A Booean value that indicates whether the system should enable the profile.
//
// # Return Value
//
// A Boolean value that indicates whether the system enabled the profile.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileResponderDelegate/willSetProfile(_:onChannel:enabled:)
func (o MIDICIProfileResponderDelegateObject) WillSetProfileOnChannelEnabled(aProfile IMIDICIProfile, channel MIDIChannelNumber, shouldEnable bool) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("willSetProfile:onChannel:enabled:"), aProfile, channel, shouldEnable)
	return rv
}

// MIDICIProfileResponderDelegateConfig holds optional typed callbacks for [MIDICIProfileResponderDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/coremidi/midiciprofileresponderdelegate
type MIDICIProfileResponderDelegateConfig struct {

	// Protocol Methods
	// WillSetProfileOnChannelEnabled — Provides an opportunity to perform an action before the system sets the profile.
	WillSetProfileOnChannelEnabled func(aProfile MIDICIProfile, channel MIDIChannelNumber, shouldEnable bool) bool
	// InitiatorDisconnected — Provides an opportunity to perform an action after the system disconnects the initiator.
	InitiatorDisconnected func(initiatorMUID MIDICIInitiatiorMUID)

	// Other Methods
	// ConnectInitiatorWithDeviceInfo — Enables a MIDI-CI initiator to create a session or reject the connection attempt.
	ConnectInitiatorWithDeviceInfo func(initiatorMUID MIDICIInitiatiorMUID, deviceInfo MIDICIDeviceInfo) bool
	// HandleDataForProfileOnChannelData — Processes MIDI data for a profile and channel.
	HandleDataForProfileOnChannelData func(aProfile MIDICIProfile, channel MIDIChannelNumber, inData foundation.NSData)
}

// NewMIDICIProfileResponderDelegate creates an Objective-C object implementing the [MIDICIProfileResponderDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [MIDICIProfileResponderDelegateObject] satisfies the [MIDICIProfileResponderDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/coremidi/midiciprofileresponderdelegate
func NewMIDICIProfileResponderDelegate(config MIDICIProfileResponderDelegateConfig) MIDICIProfileResponderDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoMIDICIProfileResponderDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ConnectInitiatorWithDeviceInfo != nil {
		fn := config.ConnectInitiatorWithDeviceInfo
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("connectInitiator:withDeviceInfo:"),
			Fn: func(self objc.ID, _cmd objc.SEL, initiatorMUID MIDICIInitiatiorMUID, deviceInfoID objc.ID) bool {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("MIDICIProfileResponderDelegate", "connectInitiator:withDeviceInfo:")
					}
				}()
				deviceInfo := MIDICIDeviceInfoFromID(deviceInfoID)
				_delegateResult := fn(initiatorMUID, deviceInfo)
				_delegateDone = true
				return _delegateResult
			},
		})
	}

	if config.InitiatorDisconnected != nil {
		fn := config.InitiatorDisconnected
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("initiatorDisconnected:"),
			Fn: func(self objc.ID, _cmd objc.SEL, initiatorMUID MIDICIInitiatiorMUID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("MIDICIProfileResponderDelegate", "initiatorDisconnected:")
					}
				}()
				fn(initiatorMUID)
				_delegateDone = true
			},
		})
	}

	if config.HandleDataForProfileOnChannelData != nil {
		fn := config.HandleDataForProfileOnChannelData
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("handleDataForProfile:onChannel:data:"),
			Fn: func(self objc.ID, _cmd objc.SEL, aProfileID objc.ID, channel MIDIChannelNumber, inDataID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("MIDICIProfileResponderDelegate", "handleDataForProfile:onChannel:data:")
					}
				}()
				aProfile := MIDICIProfileFromID(aProfileID)
				inData := foundation.NSDataFromID(inDataID)
				fn(aProfile, channel, inData)
				_delegateDone = true
			},
		})
	}

	if config.WillSetProfileOnChannelEnabled != nil {
		fn := config.WillSetProfileOnChannelEnabled
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("willSetProfile:onChannel:enabled:"),
			Fn: func(self objc.ID, _cmd objc.SEL, aProfileID objc.ID, channel MIDIChannelNumber, shouldEnable bool) bool {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("MIDICIProfileResponderDelegate", "willSetProfile:onChannel:enabled:")
					}
				}()
				aProfile := MIDICIProfileFromID(aProfileID)
				_delegateResult := fn(aProfile, channel, shouldEnable)
				_delegateDone = true
				return _delegateResult
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("MIDICIProfileResponderDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewMIDICIProfileResponderDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return MIDICIProfileResponderDelegateObjectFromID(instance)
}
