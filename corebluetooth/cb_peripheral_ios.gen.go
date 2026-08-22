// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.
//go:build ios
// +build ios

package corebluetooth

import (
	"github.com/tmc/apple/objc"
)

// # Discussion
//
// Cancels the active channel sounding session, if it exists.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/cancelChannelSoundingSession()
func (c CBPeripheral) CancelChannelSoundingSession() {
	objc.Send[objc.ID](c.ID, objc.Sel("cancelChannelSoundingSession"))
}

// configuration: An object specifying the channel sounding session configuration.
//
// # Discussion
//
// Initiate a channel sounding session.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/startChannelSoundingSession(_:)
func (c CBPeripheral) StartChannelSoundingSession(configuration *uintptr) {
	objc.Send[objc.ID](c.ID, objc.Sel("startChannelSoundingSession:"), configuration)
}

// A Boolean value that indicates if the remote device has authorization to
// receive data over ANCS protocol.
//
// # Discussion
//
// If this value is false, a user authorization sets this value to true, which
// results in a call to the delegate’s
// [CentralManagerDidUpdateANCSAuthorizationForPeripheral] method.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBPeripheral/ancsAuthorized
func (c CBPeripheral) AncsAuthorized() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("ancsAuthorized"))
	return rv
}
