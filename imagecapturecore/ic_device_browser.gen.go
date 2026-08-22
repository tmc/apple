// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ICDeviceBrowser] class.
var (
	_ICDeviceBrowserClass     ICDeviceBrowserClass
	_ICDeviceBrowserClassOnce sync.Once
)

func getICDeviceBrowserClass() ICDeviceBrowserClass {
	_ICDeviceBrowserClassOnce.Do(func() {
		_ICDeviceBrowserClass = ICDeviceBrowserClass{class: objc.GetClass("ICDeviceBrowser")}
	})
	return _ICDeviceBrowserClass
}

// GetICDeviceBrowserClass returns the class object for ICDeviceBrowser.
func GetICDeviceBrowserClass() ICDeviceBrowserClass {
	return getICDeviceBrowserClass()
}

type ICDeviceBrowserClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic ICDeviceBrowserClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic ICDeviceBrowserClass) Alloc() ICDeviceBrowser {
	rv := objc.Send[ICDeviceBrowser](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An object for finding digital cameras and scanners.
//
// # Managing Device Browsing
//
//   - [ICDeviceBrowser.Delegate]: The object that acts as the delegate of the device browser.
//   - [ICDeviceBrowser.SetDelegate]
//
// # Browsing Devices
//
//   - [ICDeviceBrowser.IsBrowsing]: A Boolean value indicating whether the device browser is browsing for devices.
//   - [ICDeviceBrowser.Devices]: All devices found by the browser.
//   - [ICDeviceBrowser.BrowsedDeviceTypeMask]: A mask whose set bits indicate the type of devices being browsed after the delegate receives the start message.
//   - [ICDeviceBrowser.SetBrowsedDeviceTypeMask]
//   - [ICDeviceBrowser.Start]: Tells the delegate to start looking for devices.
//   - [ICDeviceBrowser.Stop]: Tells the delegate to stop looking for devices.
//
// # Setting a Preferred Device
//
//   - [ICDeviceBrowser.PreferredDevice]: Returns a device object that the client application should select when it launches.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser
type ICDeviceBrowser struct {
	objectivec.Object
}

// ICDeviceBrowserFromID constructs a [ICDeviceBrowser] from an objc.ID.
//
// An object for finding digital cameras and scanners.
func ICDeviceBrowserFromID(id objc.ID) ICDeviceBrowser {
	return ICDeviceBrowser{objectivec.Object{ID: id}}
}

// NOTE: ICDeviceBrowser adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ICDeviceBrowser] class.
//
// # Managing Device Browsing
//
//   - [IICDeviceBrowser.Delegate]: The object that acts as the delegate of the device browser.
//   - [IICDeviceBrowser.SetDelegate]
//
// # Browsing Devices
//
//   - [IICDeviceBrowser.IsBrowsing]: A Boolean value indicating whether the device browser is browsing for devices.
//   - [IICDeviceBrowser.Devices]: All devices found by the browser.
//   - [IICDeviceBrowser.BrowsedDeviceTypeMask]: A mask whose set bits indicate the type of devices being browsed after the delegate receives the start message.
//   - [IICDeviceBrowser.SetBrowsedDeviceTypeMask]
//   - [IICDeviceBrowser.Start]: Tells the delegate to start looking for devices.
//   - [IICDeviceBrowser.Stop]: Tells the delegate to stop looking for devices.
//
// # Setting a Preferred Device
//
//   - [IICDeviceBrowser.PreferredDevice]: Returns a device object that the client application should select when it launches.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser
type IICDeviceBrowser interface {
	objectivec.IObject

	// Topic: Managing Device Browsing

	// The object that acts as the delegate of the device browser.
	Delegate() ICDeviceBrowserDelegate
	SetDelegate(value ICDeviceBrowserDelegate)

	// Topic: Browsing Devices

	// A Boolean value indicating whether the device browser is browsing for devices.
	IsBrowsing() bool
	// All devices found by the browser.
	Devices() []ICDevice
	// A mask whose set bits indicate the type of devices being browsed after the delegate receives the start message.
	BrowsedDeviceTypeMask() ICDeviceTypeMask
	SetBrowsedDeviceTypeMask(value ICDeviceTypeMask)
	// Tells the delegate to start looking for devices.
	Start()
	// Tells the delegate to stop looking for devices.
	Stop()

	// Topic: Setting a Preferred Device

	// Returns a device object that the client application should select when it launches.
	PreferredDevice() IICDevice
}

// Init initializes the instance.
func (d ICDeviceBrowser) Init() ICDeviceBrowser {
	rv := objc.Send[ICDeviceBrowser](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d ICDeviceBrowser) Autorelease() ICDeviceBrowser {
	rv := objc.Send[ICDeviceBrowser](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewICDeviceBrowser creates a new ICDeviceBrowser instance.
func NewICDeviceBrowser() ICDeviceBrowser {
	class := getICDeviceBrowserClass()
	rv := objc.Send[ICDeviceBrowser](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Tells the delegate to start looking for devices.
//
// # Discussion
//
// Set the [ICDeviceBrowser.Delegate] before calling this method; otherwise,
// the method call is ignored.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/start()
func (d ICDeviceBrowser) Start() {
	objc.Send[objc.ID](d.ID, objc.Sel("start"))
}

// Tells the delegate to stop looking for devices.
//
// # Discussion
//
// Calling this method frees all device instances that are not in use.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/stop()
func (d ICDeviceBrowser) Stop() {
	objc.Send[objc.ID](d.ID, objc.Sel("stop"))
}

// The object that acts as the delegate of the device browser.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/delegate
func (d ICDeviceBrowser) Delegate() ICDeviceBrowserDelegate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("delegate"))
	return ICDeviceBrowserDelegateObjectFromID(rv)
}
func (d ICDeviceBrowser) SetDelegate(value ICDeviceBrowserDelegate) {
	objc.Send[struct{}](d.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value indicating whether the device browser is browsing for
// devices.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/isBrowsing
func (d ICDeviceBrowser) IsBrowsing() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isBrowsing"))
	return rv
}

// All devices found by the browser.
//
// # Discussion
//
// This array is empty before the first invocation of the delegate method
// [DeviceBrowserDidAddDeviceMoreComing]. The value of this property changes
// as devices appear and disappear.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/devices
func (d ICDeviceBrowser) Devices() []ICDevice {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("devices"))
	return objc.ConvertSlice(rv, func(id objc.ID) ICDevice {
		return ICDeviceFromID(id)
	})
}

// A mask whose set bits indicate the type of devices being browsed after the
// delegate receives the start message.
//
// # Discussion
//
// Construct this property by performing bitwise OR on values of
// [ICDeviceTypeMask] with values of [ICDeviceLocationTypeMask]. You can
// change this property while the browser is looking for devices.
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/browsedDeviceTypeMask
//
// [ICDeviceLocationTypeMask]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceLocationTypeMask
// [ICDeviceTypeMask]: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceTypeMask
func (d ICDeviceBrowser) BrowsedDeviceTypeMask() ICDeviceTypeMask {
	rv := objc.Send[ICDeviceTypeMask](d.ID, objc.Sel("browsedDeviceTypeMask"))
	return ICDeviceTypeMask(rv)
}
func (d ICDeviceBrowser) SetBrowsedDeviceTypeMask(value ICDeviceTypeMask) {
	objc.Send[struct{}](d.ID, objc.Sel("setBrowsedDeviceTypeMask:"), value)
}

// Returns a device object that the client application should select when it
// launches.
//
// # Discussion
//
// If the client application that calls this method is the autolaunch
// application associated with a device, and that device is the last one
// attached (through USB, FireWire, or network), then that device is the
// preferred device.
//
// Call this method in the implementation of
// [DeviceBrowserDidAddDeviceMoreComing] if the value of `moreComing` is
// `false`; or in the implementation of
// [DeviceBrowserDidEnumerateLocalDevices].
//
// See: https://developer.apple.com/documentation/ImageCaptureCore/ICDeviceBrowser/preferredDevice
func (d ICDeviceBrowser) PreferredDevice() IICDevice {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("preferredDevice"))
	return ICDeviceFromID(objc.ID(rv))
}
