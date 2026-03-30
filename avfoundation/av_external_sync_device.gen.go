// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVExternalSyncDevice] class.
var (
	_AVExternalSyncDeviceClass     AVExternalSyncDeviceClass
	_AVExternalSyncDeviceClassOnce sync.Once
)

func getAVExternalSyncDeviceClass() AVExternalSyncDeviceClass {
	_AVExternalSyncDeviceClassOnce.Do(func() {
		_AVExternalSyncDeviceClass = AVExternalSyncDeviceClass{class: objc.GetClass("AVExternalSyncDevice")}
	})
	return _AVExternalSyncDeviceClass
}

// GetAVExternalSyncDeviceClass returns the class object for AVExternalSyncDevice.
func GetAVExternalSyncDeviceClass() AVExternalSyncDeviceClass {
	return getAVExternalSyncDeviceClass()
}

type AVExternalSyncDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVExternalSyncDeviceClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVExternalSyncDeviceClass) Alloc() AVExternalSyncDevice {
	rv := objc.Send[AVExternalSyncDevice](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An external sync device connected to a host device that can be used to
// drive the timing of an internal component, such as a camera sensor.
//
// # Overview
//
// Each instance of [AVExternalSyncDevice] corresponds to a physical external
// device that can drive an internal component, like a camera readout. You
// cannot create instances of [AVExternalSyncDevice]. Instead, you obtain an
// array of all currently available external sync devices using
// [AVExternalSyncDeviceDiscoverySession].
//
// # Inspecting a device
//
//   - [AVExternalSyncDevice.Clock]: A clock representing the source of time from the external sync device.
//   - [AVExternalSyncDevice.ProductID]: The USB product identifier associated with the external sync device.
//   - [AVExternalSyncDevice.SignalCompensationDelay]: Delay to wait before starting the frame capture.
//   - [AVExternalSyncDevice.SetSignalCompensationDelay]
//   - [AVExternalSyncDevice.Status]: The status of the externally connected device.
//   - [AVExternalSyncDevice.Uuid]: A unique identifier for an external sync device.
//   - [AVExternalSyncDevice.VendorID]: The USB vendor identifier associated with the external sync device.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice
type AVExternalSyncDevice struct {
	objectivec.Object
}

// AVExternalSyncDeviceFromID constructs a [AVExternalSyncDevice] from an objc.ID.
//
// An external sync device connected to a host device that can be used to
// drive the timing of an internal component, such as a camera sensor.
func AVExternalSyncDeviceFromID(id objc.ID) AVExternalSyncDevice {
	return AVExternalSyncDevice{objectivec.Object{ID: id}}
}

// NOTE: AVExternalSyncDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVExternalSyncDevice] class.
//
// # Inspecting a device
//
//   - [IAVExternalSyncDevice.Clock]: A clock representing the source of time from the external sync device.
//   - [IAVExternalSyncDevice.ProductID]: The USB product identifier associated with the external sync device.
//   - [IAVExternalSyncDevice.SignalCompensationDelay]: Delay to wait before starting the frame capture.
//   - [IAVExternalSyncDevice.SetSignalCompensationDelay]
//   - [IAVExternalSyncDevice.Status]: The status of the externally connected device.
//   - [IAVExternalSyncDevice.Uuid]: A unique identifier for an external sync device.
//   - [IAVExternalSyncDevice.VendorID]: The USB vendor identifier associated with the external sync device.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice
type IAVExternalSyncDevice interface {
	objectivec.IObject

	// Topic: Inspecting a device

	// A clock representing the source of time from the external sync device.
	Clock() uintptr
	// The USB product identifier associated with the external sync device.
	ProductID() uint32
	// Delay to wait before starting the frame capture.
	SignalCompensationDelay() coremedia.CMTime
	SetSignalCompensationDelay(value coremedia.CMTime)
	// The status of the externally connected device.
	Status() AVExternalSyncDeviceStatus
	// A unique identifier for an external sync device.
	Uuid() foundation.NSUUID
	// The USB vendor identifier associated with the external sync device.
	VendorID() uint32
}

// Init initializes the instance.
func (e AVExternalSyncDevice) Init() AVExternalSyncDevice {
	rv := objc.Send[AVExternalSyncDevice](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e AVExternalSyncDevice) Autorelease() AVExternalSyncDevice {
	rv := objc.Send[AVExternalSyncDevice](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVExternalSyncDevice creates a new AVExternalSyncDevice instance.
func NewAVExternalSyncDevice() AVExternalSyncDevice {
	class := getAVExternalSyncDeviceClass()
	rv := objc.Send[AVExternalSyncDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A clock representing the source of time from the external sync device.
//
// # Discussion
//
// This property returns [NULL] until the [Status] reaches
// [AVExternalSyncDeviceStatusActiveSync].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/clock
func (e AVExternalSyncDevice) Clock() uintptr {
	rv := objc.Send[uintptr](e.ID, objc.Sel("clock"))
	return rv
}

// The USB product identifier associated with the external sync device.
//
// # Discussion
//
// This [UInt32] value comes from the hardware vendor, and returns 0 if not
// available. Use this value in conjunction with the [VendorID] to determine a
// specific product.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/productID
func (e AVExternalSyncDevice) ProductID() uint32 {
	rv := objc.Send[uint32](e.ID, objc.Sel("productID"))
	return rv
}

// Delay to wait before starting the frame capture.
//
// # Discussion
//
// An external sync is generally used to configure multiple devices in the
// real world. A display and a camera may receive a signal at the same time,
// but that does not mean the refresh of the display and camera are aligned in
// a way that does not cause tearing in the recording. The signal compensation
// delay can be used to offset the readout of a camera on an intra-frame
// scale.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/signalCompensationDelay
func (e AVExternalSyncDevice) SignalCompensationDelay() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](e.ID, objc.Sel("signalCompensationDelay"))
	return coremedia.CMTime(rv)
}
func (e AVExternalSyncDevice) SetSignalCompensationDelay(value coremedia.CMTime) {
	objc.Send[struct{}](e.ID, objc.Sel("setSignalCompensationDelay:"), value)
}

// The status of the externally connected device.
//
// # Discussion
//
// Use this property to query the current connection status of the external
// sync device. This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/status
func (e AVExternalSyncDevice) Status() AVExternalSyncDeviceStatus {
	rv := objc.Send[AVExternalSyncDeviceStatus](e.ID, objc.Sel("status"))
	return AVExternalSyncDeviceStatus(rv)
}

// A unique identifier for an external sync device.
//
// # Discussion
//
// Use this property to select a specific external sync device.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/uuid
func (e AVExternalSyncDevice) Uuid() foundation.NSUUID {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("uuid"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}

// The USB vendor identifier associated with the external sync device.
//
// # Discussion
//
// This [UInt32] value is provided by the hardware vendor, and returns 0 if
// not available.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/vendorID
func (e AVExternalSyncDevice) VendorID() uint32 {
	rv := objc.Send[uint32](e.ID, objc.Sel("vendorID"))
	return rv
}
