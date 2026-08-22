// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A set of optional methods for receiving information about status changes for a connected Bluetooth hands-free phone or headset.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGatewayDelegate
type IOBluetoothHandsFreeAudioGatewayDelegate interface {
	objectivec.IObject
}

// IOBluetoothHandsFreeAudioGatewayDelegateObject wraps an existing Objective-C object that conforms to the IOBluetoothHandsFreeAudioGatewayDelegate protocol.
type IOBluetoothHandsFreeAudioGatewayDelegateObject struct {
	objectivec.Object
}

func (o IOBluetoothHandsFreeAudioGatewayDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// IOBluetoothHandsFreeAudioGatewayDelegateObjectFromID constructs a [IOBluetoothHandsFreeAudioGatewayDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func IOBluetoothHandsFreeAudioGatewayDelegateObjectFromID(id objc.ID) IOBluetoothHandsFreeAudioGatewayDelegateObject {
	return IOBluetoothHandsFreeAudioGatewayDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate the connected Bluetooth hands-free phone or headset is
// sending a hang-up signal.
//
// device: The remote hands-free Bluetooth device that’s sending a hang-up signal.
//
// hangup: A number that indicates whether the device is sending a hang-up signal.
// This value is always set to 1.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGatewayDelegate/handsFree(_:hangup:)
func (o IOBluetoothHandsFreeAudioGatewayDelegateObject) HandsFreeHangup(device IIOBluetoothHandsFreeAudioGateway, hangup foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:hangup:"), device, hangup)
}

// Tells the delegate the connected Bluetooth hands-free phone or headset is
// redialing the last phone number.
//
// device: The audio gateway for the remote hands-free Bluetooth device.
//
// redial: A number that indicates whether the device is attempting to redial. This
// value is always set to 1.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGatewayDelegate/handsFree(_:redial:)
func (o IOBluetoothHandsFreeAudioGatewayDelegateObject) HandsFreeRedial(device IIOBluetoothHandsFreeAudioGateway, redial foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:redial:"), device, redial)
}

// IOBluetoothHandsFreeAudioGatewayDelegateConfig holds optional typed callbacks for [IOBluetoothHandsFreeAudioGatewayDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfreeaudiogatewaydelegate
type IOBluetoothHandsFreeAudioGatewayDelegateConfig struct {

	// Receiving Status Change Information
	// HandsFreeHangup — Tells the delegate the connected Bluetooth hands-free phone or headset is sending a hang-up signal.
	HandsFreeHangup func(device IOBluetoothHandsFreeAudioGateway, hangup foundation.NSNumber)
	// HandsFreeRedial — Tells the delegate the connected Bluetooth hands-free phone or headset is redialing the last phone number.
	HandsFreeRedial func(device IOBluetoothHandsFreeAudioGateway, redial foundation.NSNumber)
}

// NewIOBluetoothHandsFreeAudioGatewayDelegate creates an Objective-C object implementing the [IOBluetoothHandsFreeAudioGatewayDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [IOBluetoothHandsFreeAudioGatewayDelegateObject] satisfies the [IOBluetoothHandsFreeAudioGatewayDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfreeaudiogatewaydelegate
func NewIOBluetoothHandsFreeAudioGatewayDelegate(config IOBluetoothHandsFreeAudioGatewayDelegateConfig) IOBluetoothHandsFreeAudioGatewayDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoIOBluetoothHandsFreeAudioGatewayDelegate_%d", n)

	var methods []objc.MethodDef

	if config.HandsFreeHangup != nil {
		fn := config.HandsFreeHangup
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("handsFree:hangup:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, hangupID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothHandsFreeAudioGatewayDelegate", "handsFree:hangup:")
					}
				}()
				device := IOBluetoothHandsFreeAudioGatewayFromID(deviceID)
				hangup := foundation.NSNumberFromID(hangupID)
				fn(device, hangup)
				_delegateDone = true
			},
		})
	}

	if config.HandsFreeRedial != nil {
		fn := config.HandsFreeRedial
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("handsFree:redial:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, redialID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothHandsFreeAudioGatewayDelegate", "handsFree:redial:")
					}
				}()
				device := IOBluetoothHandsFreeAudioGatewayFromID(deviceID)
				redial := foundation.NSNumberFromID(redialID)
				fn(device, redial)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("IOBluetoothHandsFreeAudioGatewayDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewIOBluetoothHandsFreeAudioGatewayDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return IOBluetoothHandsFreeAudioGatewayDelegateObjectFromID(instance)
}
