// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// Defines an interface for delegates of [AVCaptureDeviceInput](<doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureDeviceInput>) to respond to events that occur when connecting, calibrating, and disconnecting external sync devices.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDeviceDelegate
type AVExternalSyncDeviceDelegate interface {
	objectivec.IObject
}

// AVExternalSyncDeviceDelegateObject wraps an existing Objective-C object that conforms to the AVExternalSyncDeviceDelegate protocol.
type AVExternalSyncDeviceDelegateObject struct {
	objectivec.Object
}
func (o AVExternalSyncDeviceDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVExternalSyncDeviceDelegateObjectFromID constructs a [AVExternalSyncDeviceDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVExternalSyncDeviceDelegateObjectFromID(id objc.ID) AVExternalSyncDeviceDelegateObject {
	return AVExternalSyncDeviceDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDeviceDelegate/externalSyncDevice(_:failedWithError:)
func (o AVExternalSyncDeviceDelegateObject) ExternalSyncDeviceFailedWithError(device IAVExternalSyncDevice, error_ foundation.INSError) {
	objc.Send[struct{}](o.ID, objc.Sel("externalSyncDevice:failedWithError:"), device, error_)
	}
// Informs your delegate when the external sync device status has changed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDeviceDelegate/externalSyncDeviceStatusDidChange(_:)
func (o AVExternalSyncDeviceDelegateObject) ExternalSyncDeviceStatusDidChange(device IAVExternalSyncDevice) {
	objc.Send[struct{}](o.ID, objc.Sel("externalSyncDeviceStatusDidChange:"), device)
	}

// AVExternalSyncDeviceDelegateConfig holds optional typed callbacks for [AVExternalSyncDeviceDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avexternalsyncdevicedelegate
type AVExternalSyncDeviceDelegateConfig struct {

	// Responding to device events
	ExternalSyncDeviceFailedWithError func(device AVExternalSyncDevice, error_ foundation.NSError)
	// ExternalSyncDeviceStatusDidChange — Informs your delegate when the external sync device status has changed.
	ExternalSyncDeviceStatusDidChange func(device AVExternalSyncDevice)
}

// NewAVExternalSyncDeviceDelegate creates an Objective-C object implementing the [AVExternalSyncDeviceDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AVExternalSyncDeviceDelegateObject] satisfies the [AVExternalSyncDeviceDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avexternalsyncdevicedelegate
func NewAVExternalSyncDeviceDelegate(config AVExternalSyncDeviceDelegateConfig) AVExternalSyncDeviceDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAVExternalSyncDeviceDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ExternalSyncDeviceFailedWithError != nil {
		fn := config.ExternalSyncDeviceFailedWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("externalSyncDevice:failedWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID, error_ID objc.ID) {
				device := AVExternalSyncDeviceFromID(deviceID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(device, error_)
			},
		})
	}

	if config.ExternalSyncDeviceStatusDidChange != nil {
		fn := config.ExternalSyncDeviceStatusDidChange
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("externalSyncDeviceStatusDidChange:"),
			Fn: func(self objc.ID, _cmd objc.SEL, deviceID objc.ID) {
				device := AVExternalSyncDeviceFromID(deviceID)
				fn(device)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("AVExternalSyncDeviceDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAVExternalSyncDeviceDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AVExternalSyncDeviceDelegateObjectFromID(instance)
}

