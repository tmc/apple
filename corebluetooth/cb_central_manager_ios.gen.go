// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.
//go:build ios
// +build ios

package corebluetooth

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// Register for an event notification when the central manager makes a
// connection matching the given options.
//
// options: A dictionary that specifies options for connection events. See [Peripheral
// Connection Options] for a list of possible options.
//
// # Discussion
//
// When the central manager makes a connection that matches the options, it
// calls the delegate’s [CentralManagerConnectionEventDidOccurForPeripheral]
// method.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/registerForConnectionEvents(options:)
//
// [Peripheral Connection Options]: https://developer.apple.com/documentation/CoreBluetooth/peripheral-connection-options
func (c CBCentralManager) RegisterForConnectionEventsWithOptions(options foundation.INSDictionary) {
	objc.Send[objc.ID](c.ID, objc.Sel("registerForConnectionEventsWithOptions:"), options)
}

// Returns a Boolean that indicates whether the device supports a specific set
// of features.
//
// features: One or more features that you would like to check for support.
//
// See: https://developer.apple.com/documentation/CoreBluetooth/CBCentralManager/supports(_:)
func (_CBCentralManagerClass CBCentralManagerClass) SupportsFeatures(features CBCentralManagerFeature) bool {
	rv := objc.Send[bool](objc.ID(_CBCentralManagerClass.class), objc.Sel("supportsFeatures:"), features)
	return rv
}
