// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// Methods for responding to device events and changes.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate
type ICDeviceDelegate interface {
	objectivec.IObject

	// Tells the delegate when a session is opened on a device.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/device(_:didOpenSessionWithError:)
	DeviceDidOpenSessionWithError(device IICDevice, error_ foundation.NSError)

	// Tells the delegate when a session is closed on a device.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/device(_:didCloseSessionWithError:)
	DeviceDidCloseSessionWithError(device IICDevice, error_ foundation.NSError)

	// Tells the delegate that a device has been removed.
	//
	// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/didRemove(_:)
	DidRemoveDevice(device IICDevice)
}

// ICDeviceDelegateObject wraps an existing Objective-C object that conforms to the ICDeviceDelegate protocol.
type ICDeviceDelegateObject struct {
	objectivec.Object
}

func (o ICDeviceDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// ICDeviceDelegateObjectFromID constructs a [ICDeviceDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func ICDeviceDelegateObjectFromID(id objc.ID) ICDeviceDelegateObject {
	return ICDeviceDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate when a session is opened on a device.
//
// # Discussion
//
// This message completes the process initiated by the message
// “requestOpenSession” sent to the device object.
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/device(_:didOpenSessionWithError:)
func (o ICDeviceDelegateObject) DeviceDidOpenSessionWithError(device IICDevice, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("device:didOpenSessionWithError:"), device, error_)
}

// Tells the delegate when a session is closed on a device.
//
// # Discussion
//
// This message completes the process initiated by the message
// “requestCloseSession” sent to the device object. This message is also
// sent if the device module in control of the device ceases to control the
// device.
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/device(_:didCloseSessionWithError:)
func (o ICDeviceDelegateObject) DeviceDidCloseSessionWithError(device IICDevice, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("device:didCloseSessionWithError:"), device, error_)
}

// Tells the delegate that a device has been removed.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/didRemove(_:)
func (o ICDeviceDelegateObject) DidRemoveDevice(device IICDevice) {
	objc.Send[struct{}](o.ID, objc.Sel("didRemoveDevice:"), device)
}

// Tells the delegate when the device is ready to receive requests.
//
// # Discussion
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/deviceDidBecomeReady(_:)
func (o ICDeviceDelegateObject) DeviceDidBecomeReady(device IICDevice) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceDidBecomeReady:"), device)
}

// Tells the delegate when status information is received from a device.
//
// # Discussion
//
// The ‘status’ dictionary contains two keys, ICStatusNotificationKey and
// ICLocalizedStatusNotificationKey, which are defined above. Status
// information keys are located in their respective ICDevice type class
// header.
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/device(_:didReceiveStatusInformation:)
func (o ICDeviceDelegateObject) DeviceDidReceiveStatusInformation(device IICDevice, status foundation.INSDictionary) {
	objc.Send[struct{}](o.ID, objc.Sel("device:didReceiveStatusInformation:"), device, status)
}

// Tells the delegate when a device encounters an error.
//
// # Discussion
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/device(_:didEncounterError:)
func (o ICDeviceDelegateObject) DeviceDidEncounterError(device IICDevice, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("device:didEncounterError:"), device, error_)
}

// Tells the delegate when the ejection is complete.
//
// # Discussion
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/device(_:didEjectWithError:)
func (o ICDeviceDelegateObject) DeviceDidEjectWithError(device IICDevice, error_ foundation.NSError) {
	objc.Send[struct{}](o.ID, objc.Sel("device:didEjectWithError:"), device, error_)
}

// Tells the delegate when the name of a device changes.
//
// # Discussion
//
// This happens if the device module overrides the default name of the device
// reported by the device’s transport layer, or if the name of the
// filesystem volume mounted by the device is changed by the user.
//
// Execution of the delegate callback occurs on the main thread.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceDelegate/deviceDidChangeName(_:)
func (o ICDeviceDelegateObject) DeviceDidChangeName(device IICDevice) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceDidChangeName:"), device)
}

// ICDeviceDelegateConfig holds optional typed callbacks for [ICDeviceDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/imagecapturecore/icdevicedelegate
type ICDeviceDelegateConfig struct {

	// Responding to Device Events
	// DeviceDidOpenSessionWithError — Tells the delegate when a session is opened on a device.
	DeviceDidOpenSessionWithError func(device ICDevice, error_ foundation.NSError)
	// DeviceDidCloseSessionWithError — Tells the delegate when a session is closed on a device.
	DeviceDidCloseSessionWithError func(device ICDevice, error_ foundation.NSError)
	// DeviceDidBecomeReady — Tells the delegate when the device is ready to receive requests.
	DeviceDidBecomeReady func(device ICDevice)
	// DeviceDidReceiveStatusInformation — Tells the delegate when status information is received from a device.
	DeviceDidReceiveStatusInformation func(device ICDevice, status foundation.INSDictionary)
	// DeviceDidEncounterError — Tells the delegate when a device encounters an error.
	DeviceDidEncounterError func(device ICDevice, error_ foundation.NSError)
	// DeviceDidEjectWithError — Tells the delegate when the ejection is complete.
	DeviceDidEjectWithError func(device ICDevice, error_ foundation.NSError)

	// Responding to Device Changes
	// DeviceDidChangeName — Tells the delegate when the name of a device changes.
	DeviceDidChangeName func(device ICDevice)

	// Other Methods
	// DidRemoveDevice — Tells the delegate that a device has been removed.
	DidRemoveDevice func(device ICDevice)
}

// NewICDeviceDelegate creates an Objective-C object implementing the [ICDeviceDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [ICDeviceDelegateObject] satisfies the [ICDeviceDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/imagecapturecore/icdevicedelegate
func NewICDeviceDelegate(config ICDeviceDelegateConfig) ICDeviceDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoICDeviceDelegate_%d", n)

	var methods []objc.MethodDef

	if config.DeviceDidOpenSessionWithError != nil {
		fn := config.DeviceDidOpenSessionWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("device:didOpenSessionWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICDeviceDelegate", "device:didOpenSessionWithError:")
					}
				}()
				device := ICDeviceFromID(deviceID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(device, error_)
				_delegateDone = true
			},
		})
	}

	if config.DeviceDidCloseSessionWithError != nil {
		fn := config.DeviceDidCloseSessionWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("device:didCloseSessionWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICDeviceDelegate", "device:didCloseSessionWithError:")
					}
				}()
				device := ICDeviceFromID(deviceID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(device, error_)
				_delegateDone = true
			},
		})
	}

	if config.DidRemoveDevice != nil {
		fn := config.DidRemoveDevice
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("didRemoveDevice:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICDeviceDelegate", "didRemoveDevice:")
					}
				}()
				device := ICDeviceFromID(deviceID)
				fn(device)
				_delegateDone = true
			},
		})
	}

	if config.DeviceDidBecomeReady != nil {
		fn := config.DeviceDidBecomeReady
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("deviceDidBecomeReady:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICDeviceDelegate", "deviceDidBecomeReady:")
					}
				}()
				device := ICDeviceFromID(deviceID)
				fn(device)
				_delegateDone = true
			},
		})
	}

	if config.DeviceDidReceiveStatusInformation != nil {
		fn := config.DeviceDidReceiveStatusInformation
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("device:didReceiveStatusInformation:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, statusID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICDeviceDelegate", "device:didReceiveStatusInformation:")
					}
				}()
				device := ICDeviceFromID(deviceID)
				status := foundation.NSDictionaryFromID(statusID)
				fn(device, status)
				_delegateDone = true
			},
		})
	}

	if config.DeviceDidEncounterError != nil {
		fn := config.DeviceDidEncounterError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("device:didEncounterError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICDeviceDelegate", "device:didEncounterError:")
					}
				}()
				device := ICDeviceFromID(deviceID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(device, error_)
				_delegateDone = true
			},
		})
	}

	if config.DeviceDidEjectWithError != nil {
		fn := config.DeviceDidEjectWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("device:didEjectWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, error_ID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICDeviceDelegate", "device:didEjectWithError:")
					}
				}()
				device := ICDeviceFromID(deviceID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(device, error_)
				_delegateDone = true
			},
		})
	}

	if config.DeviceDidChangeName != nil {
		fn := config.DeviceDidChangeName
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("deviceDidChangeName:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("ICDeviceDelegate", "deviceDidChangeName:")
					}
				}()
				device := ICDeviceFromID(deviceID)
				fn(device)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("ICDeviceDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewICDeviceDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return ICDeviceDelegateObjectFromID(instance)
}
