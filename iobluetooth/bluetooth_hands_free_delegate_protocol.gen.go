// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// IOBluetoothHandsFreeDelegate protocol.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDelegate
type IOBluetoothHandsFreeDelegate interface {
	objectivec.IObject
}

// IOBluetoothHandsFreeDelegateObject wraps an existing Objective-C object that conforms to the IOBluetoothHandsFreeDelegate protocol.
type IOBluetoothHandsFreeDelegateObject struct {
	objectivec.Object
}

func (o IOBluetoothHandsFreeDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// IOBluetoothHandsFreeDelegateObjectFromID constructs a [IOBluetoothHandsFreeDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func IOBluetoothHandsFreeDelegateObjectFromID(id objc.ID) IOBluetoothHandsFreeDelegateObject {
	return IOBluetoothHandsFreeDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDelegate/handsFree(_:connected:)
func (o IOBluetoothHandsFreeDelegateObject) HandsFreeConnected(device IIOBluetoothHandsFree, status foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:connected:"), device, status)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDelegate/handsFree(_:disconnected:)
func (o IOBluetoothHandsFreeDelegateObject) HandsFreeDisconnected(device IIOBluetoothHandsFree, status foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:disconnected:"), device, status)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDelegate/handsFree(_:scoConnectionClosed:)
func (o IOBluetoothHandsFreeDelegateObject) HandsFreeScoConnectionClosed(device IIOBluetoothHandsFree, status foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:scoConnectionClosed:"), device, status)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDelegate/handsFree(_:scoConnectionOpened:)
func (o IOBluetoothHandsFreeDelegateObject) HandsFreeScoConnectionOpened(device IIOBluetoothHandsFree, status foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("handsFree:scoConnectionOpened:"), device, status)
}

// IOBluetoothHandsFreeDelegateConfig holds optional typed callbacks for [IOBluetoothHandsFreeDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfreedelegate
type IOBluetoothHandsFreeDelegateConfig struct {

	// Instance Methods
	HandsFreeConnected           func(device IOBluetoothHandsFree, status foundation.NSNumber)
	HandsFreeDisconnected        func(device IOBluetoothHandsFree, status foundation.NSNumber)
	HandsFreeScoConnectionClosed func(device IOBluetoothHandsFree, status foundation.NSNumber)
	HandsFreeScoConnectionOpened func(device IOBluetoothHandsFree, status foundation.NSNumber)
}

// NewIOBluetoothHandsFreeDelegate creates an Objective-C object implementing the [IOBluetoothHandsFreeDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [IOBluetoothHandsFreeDelegateObject] satisfies the [IOBluetoothHandsFreeDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/iobluetooth/iobluetoothhandsfreedelegate
func NewIOBluetoothHandsFreeDelegate(config IOBluetoothHandsFreeDelegateConfig) IOBluetoothHandsFreeDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoIOBluetoothHandsFreeDelegate_%d", n)

	var methods []objc.MethodDef

	if config.HandsFreeConnected != nil {
		fn := config.HandsFreeConnected
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("handsFree:connected:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, statusID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothHandsFreeDelegate", "handsFree:connected:")
					}
				}()
				device := IOBluetoothHandsFreeFromID(deviceID)
				status := foundation.NSNumberFromID(statusID)
				fn(device, status)
				_delegateDone = true
			},
		})
	}

	if config.HandsFreeDisconnected != nil {
		fn := config.HandsFreeDisconnected
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("handsFree:disconnected:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, statusID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothHandsFreeDelegate", "handsFree:disconnected:")
					}
				}()
				device := IOBluetoothHandsFreeFromID(deviceID)
				status := foundation.NSNumberFromID(statusID)
				fn(device, status)
				_delegateDone = true
			},
		})
	}

	if config.HandsFreeScoConnectionClosed != nil {
		fn := config.HandsFreeScoConnectionClosed
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("handsFree:scoConnectionClosed:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, statusID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothHandsFreeDelegate", "handsFree:scoConnectionClosed:")
					}
				}()
				device := IOBluetoothHandsFreeFromID(deviceID)
				status := foundation.NSNumberFromID(statusID)
				fn(device, status)
				_delegateDone = true
			},
		})
	}

	if config.HandsFreeScoConnectionOpened != nil {
		fn := config.HandsFreeScoConnectionOpened
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("handsFree:scoConnectionOpened:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, statusID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothHandsFreeDelegate", "handsFree:scoConnectionOpened:")
					}
				}()
				device := IOBluetoothHandsFreeFromID(deviceID)
				status := foundation.NSNumberFromID(statusID)
				fn(device, status)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("IOBluetoothHandsFreeDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewIOBluetoothHandsFreeDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return IOBluetoothHandsFreeDelegateObjectFromID(instance)
}
