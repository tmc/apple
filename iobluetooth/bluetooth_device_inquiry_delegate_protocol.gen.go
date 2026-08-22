// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"fmt"

	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// This category on NSObject describes the delegate methods for the IOBluetoothDeviceInquiry object. All methods are optional, but it is highly recommended you implement them all. Do NOT invoke remote name requests on found IOBluetoothDevice objects unless the inquiry object has been stopped. Doing so may deadlock your process.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiryDelegate
type IOBluetoothDeviceInquiryDelegate interface {
	objectivec.IObject
}

// IOBluetoothDeviceInquiryDelegateObject wraps an existing Objective-C object that conforms to the IOBluetoothDeviceInquiryDelegate protocol.
type IOBluetoothDeviceInquiryDelegateObject struct {
	objectivec.Object
}

func (o IOBluetoothDeviceInquiryDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// IOBluetoothDeviceInquiryDelegateObjectFromID constructs a [IOBluetoothDeviceInquiryDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func IOBluetoothDeviceInquiryDelegateObjectFromID(id objc.ID) IOBluetoothDeviceInquiryDelegateObject {
	return IOBluetoothDeviceInquiryDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// sender: Inquiry object that sent this delegate message.
//
// error: Error code. kIOReturnSuccess if the inquiry completed without incident.
//
// aborted: TRUE if user called -stop on the inquiry.
//
// # Discussion
//
// When the inquiry is completely stopped, this delegate method will be
// invoked. It will supply an error code value, kIOReturnSuccess if the
// inquiry stopped without problem, otherwise a non-kIOReturnSuccess error
// code will be supplied.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiryDelegate/deviceInquiryComplete(_:error:aborted:)
func (o IOBluetoothDeviceInquiryDelegateObject) DeviceInquiryCompleteErrorAborted(sender IIOBluetoothDeviceInquiry, error_ kernel.IOReturn, aborted bool) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceInquiryComplete:error:aborted:"), sender, error_, aborted)
}

// sender: Inquiry object that sent this delegate message.
//
// device: IOBluetoothDevice that was found.
//
// # Discussion
//
// A new device has been found. You do not need to retain the device - it will
// be held in the internal storage of the inquiry, and can be accessed later
// using -foundDevices.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiryDelegate/deviceInquiryDeviceFound(_:device:)
func (o IOBluetoothDeviceInquiryDelegateObject) DeviceInquiryDeviceFoundDevice(sender IIOBluetoothDeviceInquiry, device IIOBluetoothDevice) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceInquiryDeviceFound:device:"), sender, device)
}

// sender: Inquiry object that sent this delegate message.
//
// device: IOBluetoothDevice that was updated.
//
// devicesRemaining: Number of devices remaining to update.
//
// # Discussion
//
// A device name has been retrieved. Also indicates how many devices are left
// to be updated.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiryDelegate/deviceInquiryDeviceNameUpdated(_:device:devicesRemaining:)
func (o IOBluetoothDeviceInquiryDelegateObject) DeviceInquiryDeviceNameUpdatedDeviceDevicesRemaining(sender IIOBluetoothDeviceInquiry, device IIOBluetoothDevice, devicesRemaining uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceInquiryDeviceNameUpdated:device:devicesRemaining:"), sender, device, devicesRemaining)
}

// sender: Inquiry object that sent this delegate message.
//
// # Discussion
//
// This message will be delivered when the inquiry actually starts. Since the
// inquiry could be throttled, this message may not be received immediately
// after called -start.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiryDelegate/deviceInquiryStarted(_:)
func (o IOBluetoothDeviceInquiryDelegateObject) DeviceInquiryStarted(sender IIOBluetoothDeviceInquiry) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceInquiryStarted:"), sender)
}

// sender: Inquiry object that sent this delegate message.
//
// devicesRemaining: Number of devices remaining to update.
//
// # Discussion
//
// The inquiry has begun updating device names that were found during the
// search.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiryDelegate/deviceInquiryUpdatingDeviceNamesStarted(_:devicesRemaining:)
func (o IOBluetoothDeviceInquiryDelegateObject) DeviceInquiryUpdatingDeviceNamesStartedDevicesRemaining(sender IIOBluetoothDeviceInquiry, devicesRemaining uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceInquiryUpdatingDeviceNamesStarted:devicesRemaining:"), sender, devicesRemaining)
}

// IOBluetoothDeviceInquiryDelegateConfig holds optional typed callbacks for [IOBluetoothDeviceInquiryDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/iobluetooth/iobluetoothdeviceinquirydelegate
type IOBluetoothDeviceInquiryDelegateConfig struct {

	// Instance Methods
	DeviceInquiryCompleteErrorAborted                       func(sender IOBluetoothDeviceInquiry, error_ kernel.IOReturn, aborted bool)
	DeviceInquiryDeviceFoundDevice                          func(sender IOBluetoothDeviceInquiry, device IOBluetoothDevice)
	DeviceInquiryDeviceNameUpdatedDeviceDevicesRemaining    func(sender IOBluetoothDeviceInquiry, device IOBluetoothDevice, devicesRemaining uint32)
	DeviceInquiryStarted                                    func(sender IOBluetoothDeviceInquiry)
	DeviceInquiryUpdatingDeviceNamesStartedDevicesRemaining func(sender IOBluetoothDeviceInquiry, devicesRemaining uint32)
}

// NewIOBluetoothDeviceInquiryDelegate creates an Objective-C object implementing the [IOBluetoothDeviceInquiryDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [IOBluetoothDeviceInquiryDelegateObject] satisfies the [IOBluetoothDeviceInquiryDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/iobluetooth/iobluetoothdeviceinquirydelegate
func NewIOBluetoothDeviceInquiryDelegate(config IOBluetoothDeviceInquiryDelegateConfig) IOBluetoothDeviceInquiryDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoIOBluetoothDeviceInquiryDelegate_%d", n)

	var methods []objc.MethodDef

	if config.DeviceInquiryCompleteErrorAborted != nil {
		fn := config.DeviceInquiryCompleteErrorAborted
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("deviceInquiryComplete:error:aborted:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID, error_ kernel.IOReturn, aborted bool) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothDeviceInquiryDelegate", "deviceInquiryComplete:error:aborted:")
					}
				}()
				sender := IOBluetoothDeviceInquiryFromID(senderID)
				fn(sender, error_, aborted)
				_delegateDone = true
			},
		})
	}

	if config.DeviceInquiryDeviceFoundDevice != nil {
		fn := config.DeviceInquiryDeviceFoundDevice
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("deviceInquiryDeviceFound:device:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID, deviceID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothDeviceInquiryDelegate", "deviceInquiryDeviceFound:device:")
					}
				}()
				sender := IOBluetoothDeviceInquiryFromID(senderID)
				device := IOBluetoothDeviceFromID(deviceID)
				fn(sender, device)
				_delegateDone = true
			},
		})
	}

	if config.DeviceInquiryDeviceNameUpdatedDeviceDevicesRemaining != nil {
		fn := config.DeviceInquiryDeviceNameUpdatedDeviceDevicesRemaining
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("deviceInquiryDeviceNameUpdated:device:devicesRemaining:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID, deviceID objc.ID, devicesRemaining uint32) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothDeviceInquiryDelegate", "deviceInquiryDeviceNameUpdated:device:devicesRemaining:")
					}
				}()
				sender := IOBluetoothDeviceInquiryFromID(senderID)
				device := IOBluetoothDeviceFromID(deviceID)
				fn(sender, device, devicesRemaining)
				_delegateDone = true
			},
		})
	}

	if config.DeviceInquiryStarted != nil {
		fn := config.DeviceInquiryStarted
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("deviceInquiryStarted:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothDeviceInquiryDelegate", "deviceInquiryStarted:")
					}
				}()
				sender := IOBluetoothDeviceInquiryFromID(senderID)
				fn(sender)
				_delegateDone = true
			},
		})
	}

	if config.DeviceInquiryUpdatingDeviceNamesStartedDevicesRemaining != nil {
		fn := config.DeviceInquiryUpdatingDeviceNamesStartedDevicesRemaining
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("deviceInquiryUpdatingDeviceNamesStarted:devicesRemaining:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID, devicesRemaining uint32) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("IOBluetoothDeviceInquiryDelegate", "deviceInquiryUpdatingDeviceNamesStarted:devicesRemaining:")
					}
				}()
				sender := IOBluetoothDeviceInquiryFromID(senderID)
				fn(sender, devicesRemaining)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("IOBluetoothDeviceInquiryDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewIOBluetoothDeviceInquiryDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return IOBluetoothDeviceInquiryDelegateObjectFromID(instance)
}
