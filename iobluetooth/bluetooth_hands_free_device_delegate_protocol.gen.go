// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A set of optional methods for receiving status change updates and information about a connected Bluetooth hands-free phone or headset.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDeviceDelegate
type IOBluetoothHandsFreeDeviceDelegate interface {
	objectivec.IObject
	IOBluetoothHandsFreeDelegate
}

// IOBluetoothHandsFreeDeviceDelegateObject wraps an existing Objective-C object that conforms to the IOBluetoothHandsFreeDeviceDelegate protocol.
type IOBluetoothHandsFreeDeviceDelegateObject struct {
	objectivec.Object
}

func (o IOBluetoothHandsFreeDeviceDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// IOBluetoothHandsFreeDeviceDelegateObjectFromID constructs a [IOBluetoothHandsFreeDeviceDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func IOBluetoothHandsFreeDeviceDelegateObjectFromID(id objc.ID) IOBluetoothHandsFreeDeviceDelegateObject {
	return IOBluetoothHandsFreeDeviceDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate the call setup indicator of the connected Bluetooth
// hands-free phone or headset has changed.
//
// device: The connected Bluetooth hands-free phone or headset.
//
// callSetupMode: The new value of the call setup indicator. For possible values, see
// [IOBluetoothHandsFreeIndicatorCallSetup].
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDeviceDelegate/handsFree(_:callSetupMode:)
//
// [IOBluetoothHandsFreeIndicatorCallSetup]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorCallSetup
func (o IOBluetoothHandsFreeDeviceDelegateObject) HandsFreeCallSetupMode(device IIOBluetoothHandsFreeDevice, callSetupMode foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:callSetupMode:"), device, callSetupMode)
}

// Tells the delegate the active call indicator of the connected Bluetooth
// hands-free phone or headset has changed.
//
// device: The connected Bluetooth hands-free phone or headset.
//
// isCallActive: The new value of the active call indicator. For possible values, see
// [IOBluetoothHandsFreeIndicatorCall].
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDeviceDelegate/handsFree(_:isCallActive:)
//
// [IOBluetoothHandsFreeIndicatorCall]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorCall
func (o IOBluetoothHandsFreeDeviceDelegateObject) HandsFreeIsCallActive(device IIOBluetoothHandsFreeDevice, isCallActive foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:isCallActive:"), device, isCallActive)
}

// Tells the delegate the service level indicator of the connected Bluetooth
// hands-free phone or headset has changed.
//
// device: The connected Bluetooth hands-free phone or headset.
//
// isServiceAvailable: The new service level. For possible values, see
// [IOBluetoothHandsFreeIndicatorService].
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDeviceDelegate/handsFree(_:isServiceAvailable:)
//
// [IOBluetoothHandsFreeIndicatorService]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorService
func (o IOBluetoothHandsFreeDeviceDelegateObject) HandsFreeIsServiceAvailable(device IIOBluetoothHandsFreeDevice, isServiceAvailable foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:isServiceAvailable:"), device, isServiceAvailable)
}

// Tells the delegate the call setup signal strength indicator of the
// connected Bluetooth hands-free phone or headset has changed.
//
// device: The connected Bluetooth hands-free phone or headset.
//
// signalStrength: The new value of the signal strength indicator. For possible values, see
// [IOBluetoothHandsFreeIndicatorSignal].
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDeviceDelegate/handsFree(_:signalStrength:)
//
// [IOBluetoothHandsFreeIndicatorSignal]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorSignal
func (o IOBluetoothHandsFreeDeviceDelegateObject) HandsFreeSignalStrength(device IIOBluetoothHandsFreeDevice, signalStrength foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:signalStrength:"), device, signalStrength)
}

// Tells the delegate the call held indicator of the connected Bluetooth
// hands-free phone or headset has changed.
//
// device: The connected Bluetooth hands-free phone or headset.
//
// callHoldState: The new value of the call hold indicator. For possible values, see
// [IOBluetoothHandsFreeIndicatorCallHeld].
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDeviceDelegate/handsFree(_:callHoldState:)
//
// [IOBluetoothHandsFreeIndicatorCallHeld]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorCallHeld
func (o IOBluetoothHandsFreeDeviceDelegateObject) HandsFreeCallHoldState(device IIOBluetoothHandsFreeDevice, callHoldState foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:callHoldState:"), device, callHoldState)
}

// Tells the delegate the roaming indicator of the connected Bluetooth
// hands-free phone or headset has changed.
//
// device: The connected Bluetooth hands-free phone or headset.
//
// isRoaming: The new value of the roaming indicator. For possible values, see
// [IOBluetoothHandsFreeIndicatorRoam].
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDeviceDelegate/handsFree(_:isRoaming:)
//
// [IOBluetoothHandsFreeIndicatorRoam]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorRoam
func (o IOBluetoothHandsFreeDeviceDelegateObject) HandsFreeIsRoaming(device IIOBluetoothHandsFreeDevice, isRoaming foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:isRoaming:"), device, isRoaming)
}

// Tells the delegate the battery level indicator of the connected Bluetooth
// hands-free phone or headset has changed.
//
// device: The connected Bluetooth hands-free phone or headset.
//
// batteryCharge: The new value of the battery level indicator. For possible values, see
// [IOBluetoothHandsFreeIndicatorBattChg].
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDeviceDelegate/handsFree(_:batteryCharge:)
//
// [IOBluetoothHandsFreeIndicatorBattChg]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorBattChg
func (o IOBluetoothHandsFreeDeviceDelegateObject) HandsFreeBatteryCharge(device IIOBluetoothHandsFreeDevice, batteryCharge foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:batteryCharge:"), device, batteryCharge)
}

// Tells the delegate there’s an incoming call on the connected Bluetooth
// hands-free phone or headset.
//
// device: The connected Bluetooth hands-free phone or headset.
//
// number: The phone number of the caller.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDeviceDelegate/handsFree(_:incomingCallFrom:)
func (o IOBluetoothHandsFreeDeviceDelegateObject) HandsFreeIncomingCallFrom(device IIOBluetoothHandsFreeDevice, number string) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:incomingCallFrom:"), device, objc.String(number))
}

// Sends the delegate information about the current call.
//
// device: The connected Bluetooth hands-free phone or headset.
//
// currentCall: A dictionary with the incoming SMS message. For dictionary keys, see
// [Current Call Information Constants].
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDeviceDelegate/handsFree(_:currentCall:)
//
// [Current Call Information Constants]: https://developer.apple.com/documentation/IOBluetooth/current-call-information-constants
func (o IOBluetoothHandsFreeDeviceDelegateObject) HandsFreeCurrentCall(device IIOBluetoothHandsFreeDevice, currentCall foundation.INSDictionary) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:currentCall:"), device, currentCall)
}

// Tells the delegate there’s an incoming text message.
//
// device: The connected Bluetooth hands-free phone or headset.
//
// sms: A dictionary containing the incoming SMS message. For dictionary keys, see
// [SMS Dictionary Key Constants].
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDeviceDelegate/handsFree(_:incomingSMS:)
//
// [SMS Dictionary Key Constants]: https://developer.apple.com/documentation/IOBluetooth/sms-dictionary-key-constants
func (o IOBluetoothHandsFreeDeviceDelegateObject) HandsFreeIncomingSMS(device IIOBluetoothHandsFreeDevice, sms foundation.INSDictionary) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:incomingSMS:"), device, sms)
}

// Tells the delegate the subscriber number of a call.
//
// device: The connected Bluetooth hands-free phone or headset.
//
// subscriberNumber: The subscriber number.
//
// # Discussion
//
// If multiple subscriber numbers are on the gateway, this function is called
// once for each subscriber number.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDeviceDelegate/handsFree(_:subscriberNumber:)
func (o IOBluetoothHandsFreeDeviceDelegateObject) HandsFreeSubscriberNumber(device IIOBluetoothHandsFreeDevice, subscriberNumber string) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:subscriberNumber:"), device, objc.String(subscriberNumber))
}

// Tells the delegate the phone is ringing.
//
// device: The connected Bluetooth hands-free phone or headset.
//
// ringAttempt: The number of ring alerts received for the call.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDeviceDelegate/handsFree(_:ringAttempt:)
func (o IOBluetoothHandsFreeDeviceDelegateObject) HandsFreeRingAttempt(device IIOBluetoothHandsFreeDevice, ringAttempt foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:ringAttempt:"), device, ringAttempt)
}

// Tells the delegate the phone sent an unknown code.
//
// device: The connected Bluetooth hands-free phone or headset.
//
// resultCode: A string containing the result code. The `“/r/n”` strings are stripped
// from the beginning and end.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDeviceDelegate/handsFree(_:unhandledResultCode:)
func (o IOBluetoothHandsFreeDeviceDelegateObject) HandsFreeUnhandledResultCode(device IIOBluetoothHandsFreeDevice, resultCode string) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:unhandledResultCode:"), device, objc.String(resultCode))
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDelegate/handsFree(_:connected:)
func (o IOBluetoothHandsFreeDeviceDelegateObject) HandsFreeConnected(device IIOBluetoothHandsFree, status foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:connected:"), device, status)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDelegate/handsFree(_:disconnected:)
func (o IOBluetoothHandsFreeDeviceDelegateObject) HandsFreeDisconnected(device IIOBluetoothHandsFree, status foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:disconnected:"), device, status)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDelegate/handsFree(_:scoConnectionClosed:)
func (o IOBluetoothHandsFreeDeviceDelegateObject) HandsFreeScoConnectionClosed(device IIOBluetoothHandsFree, status foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:scoConnectionClosed:"), device, status)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDelegate/handsFree(_:scoConnectionOpened:)
func (o IOBluetoothHandsFreeDeviceDelegateObject) HandsFreeScoConnectionOpened(device IIOBluetoothHandsFree, status foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:scoConnectionOpened:"), device, status)
}

// IOBluetoothHandsFreeDeviceDelegateConfig holds optional typed callbacks for [IOBluetoothHandsFreeDeviceDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfreedevicedelegate
type IOBluetoothHandsFreeDeviceDelegateConfig struct {

	// Receiving Status Indicator Changes
	// HandsFreeCallSetupMode — Tells the delegate the call setup indicator of the connected Bluetooth hands-free phone or headset has changed.
	HandsFreeCallSetupMode func(device IOBluetoothHandsFreeDevice, callSetupMode foundation.NSNumber)
	// HandsFreeIsCallActive — Tells the delegate the active call indicator of the connected Bluetooth hands-free phone or headset has changed.
	HandsFreeIsCallActive func(device IOBluetoothHandsFreeDevice, isCallActive foundation.NSNumber)
	// HandsFreeIsServiceAvailable — Tells the delegate the service level indicator of the connected Bluetooth hands-free phone or headset has changed.
	HandsFreeIsServiceAvailable func(device IOBluetoothHandsFreeDevice, isServiceAvailable foundation.NSNumber)
	// HandsFreeSignalStrength — Tells the delegate the call setup signal strength indicator of the connected Bluetooth hands-free phone or headset has changed.
	HandsFreeSignalStrength func(device IOBluetoothHandsFreeDevice, signalStrength foundation.NSNumber)
	// HandsFreeCallHoldState — Tells the delegate the call held indicator of the connected Bluetooth hands-free phone or headset has changed.
	HandsFreeCallHoldState func(device IOBluetoothHandsFreeDevice, callHoldState foundation.NSNumber)
	// HandsFreeIsRoaming — Tells the delegate the roaming indicator of the connected Bluetooth hands-free phone or headset has changed.
	HandsFreeIsRoaming func(device IOBluetoothHandsFreeDevice, isRoaming foundation.NSNumber)
	// HandsFreeBatteryCharge — Tells the delegate the battery level indicator of the connected Bluetooth hands-free phone or headset has changed.
	HandsFreeBatteryCharge func(device IOBluetoothHandsFreeDevice, batteryCharge foundation.NSNumber)

	// Receiving Call Status
	// HandsFreeCurrentCall — Sends the delegate information about the current call.
	HandsFreeCurrentCall func(device IOBluetoothHandsFreeDevice, currentCall foundation.INSDictionary)

	// Receiving SMS Information
	// HandsFreeIncomingSMS — Tells the delegate there’s an incoming text message.
	HandsFreeIncomingSMS func(device IOBluetoothHandsFreeDevice, sms foundation.INSDictionary)

	// Receiving Other Information
	// HandsFreeRingAttempt — Tells the delegate the phone is ringing.
	HandsFreeRingAttempt func(device IOBluetoothHandsFreeDevice, ringAttempt foundation.NSNumber)
}

// NewIOBluetoothHandsFreeDeviceDelegate creates an Objective-C object implementing the [IOBluetoothHandsFreeDeviceDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [IOBluetoothHandsFreeDeviceDelegateObject] satisfies the [IOBluetoothHandsFreeDeviceDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfreedevicedelegate
func NewIOBluetoothHandsFreeDeviceDelegate(config IOBluetoothHandsFreeDeviceDelegateConfig) IOBluetoothHandsFreeDeviceDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoIOBluetoothHandsFreeDeviceDelegate_%d", n)

	var methods []objc.MethodDef

	if config.HandsFreeCallSetupMode != nil {
		fn := config.HandsFreeCallSetupMode
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("handsFree:callSetupMode:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, callSetupModeID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothHandsFreeDeviceDelegate", "handsFree:callSetupMode:")
					}
				}()
				device := IOBluetoothHandsFreeDeviceFromID(deviceID)
				callSetupMode := foundation.NSNumberFromID(callSetupModeID)
				fn(device, callSetupMode)
				_delegateDone = true
			},
		})
	}

	if config.HandsFreeIsCallActive != nil {
		fn := config.HandsFreeIsCallActive
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("handsFree:isCallActive:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, isCallActiveID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothHandsFreeDeviceDelegate", "handsFree:isCallActive:")
					}
				}()
				device := IOBluetoothHandsFreeDeviceFromID(deviceID)
				isCallActive := foundation.NSNumberFromID(isCallActiveID)
				fn(device, isCallActive)
				_delegateDone = true
			},
		})
	}

	if config.HandsFreeIsServiceAvailable != nil {
		fn := config.HandsFreeIsServiceAvailable
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("handsFree:isServiceAvailable:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, isServiceAvailableID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothHandsFreeDeviceDelegate", "handsFree:isServiceAvailable:")
					}
				}()
				device := IOBluetoothHandsFreeDeviceFromID(deviceID)
				isServiceAvailable := foundation.NSNumberFromID(isServiceAvailableID)
				fn(device, isServiceAvailable)
				_delegateDone = true
			},
		})
	}

	if config.HandsFreeSignalStrength != nil {
		fn := config.HandsFreeSignalStrength
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("handsFree:signalStrength:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, signalStrengthID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothHandsFreeDeviceDelegate", "handsFree:signalStrength:")
					}
				}()
				device := IOBluetoothHandsFreeDeviceFromID(deviceID)
				signalStrength := foundation.NSNumberFromID(signalStrengthID)
				fn(device, signalStrength)
				_delegateDone = true
			},
		})
	}

	if config.HandsFreeCallHoldState != nil {
		fn := config.HandsFreeCallHoldState
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("handsFree:callHoldState:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, callHoldStateID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothHandsFreeDeviceDelegate", "handsFree:callHoldState:")
					}
				}()
				device := IOBluetoothHandsFreeDeviceFromID(deviceID)
				callHoldState := foundation.NSNumberFromID(callHoldStateID)
				fn(device, callHoldState)
				_delegateDone = true
			},
		})
	}

	if config.HandsFreeIsRoaming != nil {
		fn := config.HandsFreeIsRoaming
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("handsFree:isRoaming:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, isRoamingID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothHandsFreeDeviceDelegate", "handsFree:isRoaming:")
					}
				}()
				device := IOBluetoothHandsFreeDeviceFromID(deviceID)
				isRoaming := foundation.NSNumberFromID(isRoamingID)
				fn(device, isRoaming)
				_delegateDone = true
			},
		})
	}

	if config.HandsFreeBatteryCharge != nil {
		fn := config.HandsFreeBatteryCharge
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("handsFree:batteryCharge:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, batteryChargeID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothHandsFreeDeviceDelegate", "handsFree:batteryCharge:")
					}
				}()
				device := IOBluetoothHandsFreeDeviceFromID(deviceID)
				batteryCharge := foundation.NSNumberFromID(batteryChargeID)
				fn(device, batteryCharge)
				_delegateDone = true
			},
		})
	}

	if config.HandsFreeCurrentCall != nil {
		fn := config.HandsFreeCurrentCall
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("handsFree:currentCall:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, currentCallID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothHandsFreeDeviceDelegate", "handsFree:currentCall:")
					}
				}()
				device := IOBluetoothHandsFreeDeviceFromID(deviceID)
				currentCall := foundation.NSDictionaryFromID(currentCallID)
				fn(device, currentCall)
				_delegateDone = true
			},
		})
	}

	if config.HandsFreeIncomingSMS != nil {
		fn := config.HandsFreeIncomingSMS
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("handsFree:incomingSMS:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, smsID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothHandsFreeDeviceDelegate", "handsFree:incomingSMS:")
					}
				}()
				device := IOBluetoothHandsFreeDeviceFromID(deviceID)
				sms := foundation.NSDictionaryFromID(smsID)
				fn(device, sms)
				_delegateDone = true
			},
		})
	}

	if config.HandsFreeRingAttempt != nil {
		fn := config.HandsFreeRingAttempt
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("handsFree:ringAttempt:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, ringAttemptID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothHandsFreeDeviceDelegate", "handsFree:ringAttempt:")
					}
				}()
				device := IOBluetoothHandsFreeDeviceFromID(deviceID)
				ringAttempt := foundation.NSNumberFromID(ringAttemptID)
				fn(device, ringAttempt)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("IOBluetoothHandsFreeDeviceDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewIOBluetoothHandsFreeDeviceDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return IOBluetoothHandsFreeDeviceDelegateObjectFromID(instance)
}
