// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// IOBluetoothDevicePairDelegate protocol.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePairDelegate
type IOBluetoothDevicePairDelegate interface {
	objectivec.IObject
}

// IOBluetoothDevicePairDelegateObject wraps an existing Objective-C object that conforms to the IOBluetoothDevicePairDelegate protocol.
type IOBluetoothDevicePairDelegateObject struct {
	objectivec.Object
}

func (o IOBluetoothDevicePairDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// IOBluetoothDevicePairDelegateObjectFromID constructs a [IOBluetoothDevicePairDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func IOBluetoothDevicePairDelegateObjectFromID(id objc.ID) IOBluetoothDevicePairDelegateObject {
	return IOBluetoothDevicePairDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePairDelegate/devicePairingConnected(_:)
func (o IOBluetoothDevicePairDelegateObject) DevicePairingConnected(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("devicePairingConnected:"), sender)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePairDelegate/devicePairingConnecting(_:)
func (o IOBluetoothDevicePairDelegateObject) DevicePairingConnecting(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("devicePairingConnecting:"), sender)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePairDelegate/devicePairingFinished(_:error:)
func (o IOBluetoothDevicePairDelegateObject) DevicePairingFinishedError(sender objectivec.IObject, error_ kernel.IOReturn) {
	objc.Send[struct{}](o.ID, objc.Sel("devicePairingFinished:error:"), sender, error_)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePairDelegate/devicePairingPINCodeRequest(_:)
func (o IOBluetoothDevicePairDelegateObject) DevicePairingPINCodeRequest(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("devicePairingPINCodeRequest:"), sender)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePairDelegate/devicePairingStarted(_:)
func (o IOBluetoothDevicePairDelegateObject) DevicePairingStarted(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("devicePairingStarted:"), sender)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePairDelegate/devicePairingUserConfirmationRequest(_:numericValue:)
func (o IOBluetoothDevicePairDelegateObject) DevicePairingUserConfirmationRequestNumericValue(sender objectivec.IObject, numericValue BluetoothNumericValue) {
	objc.Send[struct{}](o.ID, objc.Sel("devicePairingUserConfirmationRequest:numericValue:"), sender, numericValue)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePairDelegate/devicePairingUserPasskeyNotification(_:passkey:)
func (o IOBluetoothDevicePairDelegateObject) DevicePairingUserPasskeyNotificationPasskey(sender objectivec.IObject, passkey BluetoothPasskey) {
	objc.Send[struct{}](o.ID, objc.Sel("devicePairingUserPasskeyNotification:passkey:"), sender, passkey)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePairDelegate/deviceSimplePairingComplete(_:status:)
func (o IOBluetoothDevicePairDelegateObject) DeviceSimplePairingCompleteStatus(sender objectivec.IObject, status BluetoothHCIEventStatus) {
	objc.Send[struct{}](o.ID, objc.Sel("deviceSimplePairingComplete:status:"), sender, status)
}
