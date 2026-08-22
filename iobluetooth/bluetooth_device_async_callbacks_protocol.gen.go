// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// IOBluetoothDeviceAsyncCallbacks protocol.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceAsyncCallbacks
type IOBluetoothDeviceAsyncCallbacks interface {
	objectivec.IObject

	// ConnectionCompleteStatus protocol.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceAsyncCallbacks/connectionComplete(_:status:)
	ConnectionCompleteStatus(device IIOBluetoothDevice, status kernel.IOReturn)

	// RemoteNameRequestCompleteStatus protocol.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceAsyncCallbacks/remoteNameRequestComplete(_:status:)
	RemoteNameRequestCompleteStatus(device IIOBluetoothDevice, status kernel.IOReturn)

	// SdpQueryCompleteStatus protocol.
	//
	// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceAsyncCallbacks/sdpQueryComplete(_:status:)
	SdpQueryCompleteStatus(device IIOBluetoothDevice, status kernel.IOReturn)
}

// IOBluetoothDeviceAsyncCallbacksObject wraps an existing Objective-C object that conforms to the IOBluetoothDeviceAsyncCallbacks protocol.
type IOBluetoothDeviceAsyncCallbacksObject struct {
	objectivec.Object
}

func (o IOBluetoothDeviceAsyncCallbacksObject) BaseObject() objectivec.Object {
	return o.Object
}

// IOBluetoothDeviceAsyncCallbacksObjectFromID constructs a [IOBluetoothDeviceAsyncCallbacksObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func IOBluetoothDeviceAsyncCallbacksObjectFromID(id objc.ID) IOBluetoothDeviceAsyncCallbacksObject {
	return IOBluetoothDeviceAsyncCallbacksObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceAsyncCallbacks/connectionComplete(_:status:)
func (o IOBluetoothDeviceAsyncCallbacksObject) ConnectionCompleteStatus(device IIOBluetoothDevice, status kernel.IOReturn) {
	objc.Send[struct{}](o.ID, objc.Sel("connectionComplete:status:"), device, status)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceAsyncCallbacks/remoteNameRequestComplete(_:status:)
func (o IOBluetoothDeviceAsyncCallbacksObject) RemoteNameRequestCompleteStatus(device IIOBluetoothDevice, status kernel.IOReturn) {
	objc.Send[struct{}](o.ID, objc.Sel("remoteNameRequestComplete:status:"), device, status)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceAsyncCallbacks/sdpQueryComplete(_:status:)
func (o IOBluetoothDeviceAsyncCallbacksObject) SdpQueryCompleteStatus(device IIOBluetoothDevice, status kernel.IOReturn) {
	objc.Send[struct{}](o.ID, objc.Sel("sdpQueryComplete:status:"), device, status)
}
