// Code generated from Apple documentation for DeviceCheck. DO NOT EDIT.

package devicecheck

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DCDevice] class.
var (
	_DCDeviceClass     DCDeviceClass
	_DCDeviceClassOnce sync.Once
)

func getDCDeviceClass() DCDeviceClass {
	_DCDeviceClassOnce.Do(func() {
		_DCDeviceClass = DCDeviceClass{class: objc.GetClass("DCDevice")}
	})
	return _DCDeviceClass
}

// GetDCDeviceClass returns the class object for DCDevice.
func GetDCDeviceClass() DCDeviceClass {
	return getDCDeviceClass()
}

type DCDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DCDeviceClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DCDeviceClass) Alloc() DCDevice {
	rv := objc.Send[DCDevice](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a device that provides a unique, authenticated token.
//
// # Overview
//
// Use the shared instance of the [DCDevice] class to generate a token that
// identifies a device. Call the [DCDevice.GenerateTokenWithCompletionHandler]
// method to get the token, and then send it to your server:
//
// On your server, combine the token with an authentication key that you
// obtain from Apple, and use the result to request access to two per-device
// binary digits (bits). After authenticating the device, Apple passes the
// current values of the bits, along with the date they were last modified, to
// your server. Your server applies its business logic to this information and
// communicates the results to your app. For more information about
// server-side procedures, see [Accessing and modifying per-device data].
//
// Apple records the bits for you, and reports the bits back to you, but
// you’re responsible for keeping track of what the bits mean. You’re also
// responsible for determining when to reset the bits for a given device; for
// example, when a user sells the device to someone else.
//
// # Determining API support
//
//   - [DCDevice.IsSupported]: A Boolean value that indicates whether the device supports the DeviceCheck API.
//
// # Getting a device token
//
//   - [DCDevice.GenerateTokenWithCompletionHandler]: Generates a token that identifies the current device.
//
// See: https://developer.apple.com/documentation/DeviceCheck/DCDevice
//
// [Accessing and modifying per-device data]: https://developer.apple.com/documentation/DeviceCheck/accessing-and-modifying-per-device-data
type DCDevice struct {
	objectivec.Object
}

// DCDeviceFromID constructs a [DCDevice] from an objc.ID.
//
// A representation of a device that provides a unique, authenticated token.
func DCDeviceFromID(id objc.ID) DCDevice {
	return DCDevice{objectivec.Object{ID: id}}
}

// NOTE: DCDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [DCDevice] class.
//
// # Determining API support
//
//   - [IDCDevice.IsSupported]: A Boolean value that indicates whether the device supports the DeviceCheck API.
//
// # Getting a device token
//
//   - [IDCDevice.GenerateTokenWithCompletionHandler]: Generates a token that identifies the current device.
//
// See: https://developer.apple.com/documentation/DeviceCheck/DCDevice
type IDCDevice interface {
	objectivec.IObject

	// Topic: Determining API support

	// A Boolean value that indicates whether the device supports the DeviceCheck API.
	IsSupported() bool

	// Topic: Getting a device token

	// Generates a token that identifies the current device.
	GenerateTokenWithCompletionHandler(completion DataErrorHandler)
}

// Init initializes the instance.
func (d DCDevice) Init() DCDevice {
	rv := objc.Send[DCDevice](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DCDevice) Autorelease() DCDevice {
	rv := objc.Send[DCDevice](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDCDevice creates a new DCDevice instance.
func NewDCDevice() DCDevice {
	class := getDCDeviceClass()
	rv := objc.Send[DCDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Generates a token that identifies the current device.
//
// completion: A completion block that includes the following parameters:
//
// - `token`: An ephemeral token that identifies the current device. -
// `error`: The error that occurred, if any.
//
// # Discussion
//
// Your server uses the generated token in its requests to get or set the
// persistent bits for the current device. You should treat the token you
// receive in the completion block as single-use. Although the token remains
// valid long enough for your server to retry a specific request if necessary,
// you should not use a token multiple times. Instead, use this method to
// generate a new token.
//
// See: https://developer.apple.com/documentation/DeviceCheck/DCDevice/generateToken(completionHandler:)
func (d DCDevice) GenerateTokenWithCompletionHandler(completion DataErrorHandler) {
	_block0, _ := NewDataErrorBlock(completion)
	objc.Send[objc.ID](d.ID, objc.Sel("generateTokenWithCompletionHandler:"), _block0)
}

// A Boolean value that indicates whether the device supports the DeviceCheck
// API.
//
// See: https://developer.apple.com/documentation/DeviceCheck/DCDevice/isSupported
func (d DCDevice) IsSupported() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isSupported"))
	return rv
}

// A representation of the device for which you want to query the two bits of
// data.
//
// See: https://developer.apple.com/documentation/DeviceCheck/DCDevice/current
func (_DCDeviceClass DCDeviceClass) CurrentDevice() DCDevice {
	rv := objc.Send[objc.ID](objc.ID(_DCDeviceClass.class), objc.Sel("currentDevice"))
	return DCDeviceFromID(objc.ID(rv))
}

// GenerateToken is a synchronous wrapper around [DCDevice.GenerateTokenWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d DCDevice) GenerateToken(ctx context.Context) (*foundation.NSData, error) {
	type result struct {
		val *foundation.NSData
		err error
	}
	done := make(chan result, 1)
	d.GenerateTokenWithCompletionHandler(func(val *foundation.NSData, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
