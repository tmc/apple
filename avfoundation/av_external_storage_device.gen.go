// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVExternalStorageDevice] class.
var (
	_AVExternalStorageDeviceClass     AVExternalStorageDeviceClass
	_AVExternalStorageDeviceClassOnce sync.Once
)

func getAVExternalStorageDeviceClass() AVExternalStorageDeviceClass {
	_AVExternalStorageDeviceClassOnce.Do(func() {
		_AVExternalStorageDeviceClass = AVExternalStorageDeviceClass{class: objc.GetClass("AVExternalStorageDevice")}
	})
	return _AVExternalStorageDeviceClass
}

// GetAVExternalStorageDeviceClass returns the class object for AVExternalStorageDevice.
func GetAVExternalStorageDeviceClass() AVExternalStorageDeviceClass {
	return getAVExternalStorageDeviceClass()
}

type AVExternalStorageDeviceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVExternalStorageDeviceClass) Alloc() AVExternalStorageDevice {
	rv := objc.Send[AVExternalStorageDevice](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// Represents a physical external storage device that stores media assets.
//
// # Overview
// 
// Each storage device instance corresponds to a physical external storage
// device where the system can media assets. You can access all of the
// currently available external storage devices with the
// [AVExternalStorageDeviceDiscoverySession] object’s
// [AVExternalStorageDevice.ExternalStorageDevices] property.
//
// # Generating URLs for image assets
//
//   - [AVExternalStorageDevice.NextAvailableURLsWithPathExtensionsError]: Generates an array of security scoped URLs that are compliant for digital camera formats, where each element has a different path extension.
//
// # Inspecting a storage device
//
//   - [AVExternalStorageDevice.Connected]: A Boolean value that indicates whether the system has a connection to the external storage device.
//   - [AVExternalStorageDevice.DisplayName]: The name of an external storage device that’s appropriate for a user interface.
//   - [AVExternalStorageDevice.Uuid]: The external storage device’s unique identifier.
//   - [AVExternalStorageDevice.FreeSize]: The amount of free storage space, in bytes, that’s available on the external storage device.
//   - [AVExternalStorageDevice.TotalSize]: The total amount of storage space, in bytes, that’s available on the external storage device.
//   - [AVExternalStorageDevice.NotRecommendedForCaptureUse]: A Boolean value that indicates whether the external storage device is suitable for camera capture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice
type AVExternalStorageDevice struct {
	objectivec.Object
}

// AVExternalStorageDeviceFromID constructs a [AVExternalStorageDevice] from an objc.ID.
//
// Represents a physical external storage device that stores media assets.
func AVExternalStorageDeviceFromID(id objc.ID) AVExternalStorageDevice {
	return AVExternalStorageDevice{objectivec.Object{ID: id}}
}
// NOTE: AVExternalStorageDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVExternalStorageDevice] class.
//
// # Generating URLs for image assets
//
//   - [IAVExternalStorageDevice.NextAvailableURLsWithPathExtensionsError]: Generates an array of security scoped URLs that are compliant for digital camera formats, where each element has a different path extension.
//
// # Inspecting a storage device
//
//   - [IAVExternalStorageDevice.Connected]: A Boolean value that indicates whether the system has a connection to the external storage device.
//   - [IAVExternalStorageDevice.DisplayName]: The name of an external storage device that’s appropriate for a user interface.
//   - [IAVExternalStorageDevice.Uuid]: The external storage device’s unique identifier.
//   - [IAVExternalStorageDevice.FreeSize]: The amount of free storage space, in bytes, that’s available on the external storage device.
//   - [IAVExternalStorageDevice.TotalSize]: The total amount of storage space, in bytes, that’s available on the external storage device.
//   - [IAVExternalStorageDevice.NotRecommendedForCaptureUse]: A Boolean value that indicates whether the external storage device is suitable for camera capture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice
type IAVExternalStorageDevice interface {
	objectivec.IObject

	// Topic: Generating URLs for image assets

	// Generates an array of security scoped URLs that are compliant for digital camera formats, where each element has a different path extension.
	NextAvailableURLsWithPathExtensionsError(extensionArray []string) ([]foundation.NSURL, error)

	// Topic: Inspecting a storage device

	// A Boolean value that indicates whether the system has a connection to the external storage device.
	Connected() bool
	// The name of an external storage device that’s appropriate for a user interface.
	DisplayName() string
	// The external storage device’s unique identifier.
	Uuid() foundation.NSUUID
	// The amount of free storage space, in bytes, that’s available on the external storage device.
	FreeSize() int
	// The total amount of storage space, in bytes, that’s available on the external storage device.
	TotalSize() int
	// A Boolean value that indicates whether the external storage device is suitable for camera capture.
	NotRecommendedForCaptureUse() bool

	// An array of external storage devices the session updates as individual devices connect or disconnect from the system.
	ExternalStorageDevices() IAVExternalStorageDevice
	SetExternalStorageDevices(value IAVExternalStorageDevice)
}

// Init initializes the instance.
func (e AVExternalStorageDevice) Init() AVExternalStorageDevice {
	rv := objc.Send[AVExternalStorageDevice](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e AVExternalStorageDevice) Autorelease() AVExternalStorageDevice {
	rv := objc.Send[AVExternalStorageDevice](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVExternalStorageDevice creates a new AVExternalStorageDevice instance.
func NewAVExternalStorageDevice() AVExternalStorageDevice {
	class := getAVExternalStorageDeviceClass()
	rv := objc.Send[AVExternalStorageDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Generates an array of security scoped URLs that are compliant for digital
// camera formats, where each element has a different path extension.
//
// extensionArray: An array of path extensions the method generates URLs for.
//
// # Return Value
// 
// An array of digital camera format (DCF) compliant URLs with security
// scoping, one for each path extension element in `extensionArray`.
//
// # Discussion
// 
// The method generates a digital camera format (DCF) compliant URL with
// security scoping for each file extension element in `extensionArray`. It
// does this by configuring the folder structure and, if necessary, creates a
// digital camera image (DCIM) folder on the external storage device.
// 
// # Request access to the storage device before request
// 
// Your app can request authorization before calling the method if
// [AuthorizationStatus] is [AuthorizationStatusNotDetermined] by calling the
// [RequestAccessWithCompletionHandler] method first.
// 
// # Start and stop access to a URL around your code
// 
// To access one of the security-scoped URLs the method returns, you need to
// call the [startAccessingSecurityScopedResource()], and
// [stopAccessingSecurityScopedResource()] methods before and after your code.
//
// [startAccessingSecurityScopedResource()]: https://developer.apple.com/documentation/Foundation/URL/startAccessingSecurityScopedResource()
// [stopAccessingSecurityScopedResource()]: https://developer.apple.com/documentation/Foundation/URL/stopAccessingSecurityScopedResource()
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/nextAvailableURLs(withPathExtensions:)
func (e AVExternalStorageDevice) NextAvailableURLsWithPathExtensionsError(extensionArray []string) ([]foundation.NSURL, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("nextAvailableURLsWithPathExtensions:error:"), objectivec.StringSliceToNSArray(extensionArray), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSURL {
		return foundation.NSURLFromID(id)
	}), nil

}

// Requests access to an external storage device on behalf of your app, which
// can present a dialog to a person on their device’s display.
//
// handler: A closure you provide the system calls to inform your app whether the
// person authorizes it to access the storage device. The system can call your
// closure on any dispatch queue and it’s your app’s responsibility to
// update its UI on the main thread or queue.
//
// # Discussion
// 
// Use this method to request access to save image assets to the external
// storage device. Your app can’t access the external storage device without
// generating an error until a person gives it permission.
// 
// The system only presents the dialog to a person the first time your app
// calls the method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/requestAccess(completionHandler:)
func (_AVExternalStorageDeviceClass AVExternalStorageDeviceClass) RequestAccessWithCompletionHandler(handler BoolHandler) {
_block0, _cleanup0 := NewBoolBlock(handler)
	defer _cleanup0()
	objc.Send[objc.ID](objc.ID(_AVExternalStorageDeviceClass.class), objc.Sel("requestAccessWithCompletionHandler:"), _block0)
}

// A Boolean value that indicates whether the system has a connection to the
// external storage device.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/isConnected
func (e AVExternalStorageDevice) Connected() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("isConnected"))
	return rv
}
// The name of an external storage device that’s appropriate for a user
// interface.
//
// # Discussion
// 
// The value is `nil` when the system can’t retrieve information from
// external storage device.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/displayName
func (e AVExternalStorageDevice) DisplayName() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("displayName"))
	return foundation.NSStringFromID(rv).String()
}
// The external storage device’s unique identifier.
//
// # Discussion
// 
// The value is `nil` when the system can’t retrieve information from
// external storage device.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/uuid
func (e AVExternalStorageDevice) Uuid() foundation.NSUUID {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("uuid"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
// The amount of free storage space, in bytes, that’s available on the
// external storage device.
//
// # Discussion
// 
// The value is `-1` when the system can’t retrieve information from
// external storage device.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/freeSize
func (e AVExternalStorageDevice) FreeSize() int {
	rv := objc.Send[int](e.ID, objc.Sel("freeSize"))
	return rv
}
// The total amount of storage space, in bytes, that’s available on the
// external storage device.
//
// # Discussion
// 
// The value is `-1` when the system can’t retrieve information from
// external storage device.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/totalSize
func (e AVExternalStorageDevice) TotalSize() int {
	rv := objc.Send[int](e.ID, objc.Sel("totalSize"))
	return rv
}
// A Boolean value that indicates whether the external storage device is
// suitable for camera capture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/isNotRecommendedForCaptureUse
func (e AVExternalStorageDevice) NotRecommendedForCaptureUse() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("isNotRecommendedForCaptureUse"))
	return rv
}
// An array of external storage devices the session updates as individual
// devices connect or disconnect from the system.
//
// See: https://developer.apple.com/documentation/avfoundation/avexternalstoragedevicediscoverysession/externalstoragedevices
func (e AVExternalStorageDevice) ExternalStorageDevices() IAVExternalStorageDevice {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("externalStorageDevices"))
	return AVExternalStorageDeviceFromID(objc.ID(rv))
}
func (e AVExternalStorageDevice) SetExternalStorageDevices(value IAVExternalStorageDevice) {
	objc.Send[struct{}](e.ID, objc.Sel("setExternalStorageDevices:"), value)
}

// Your app’s authorization status for the external storage device.
//
// # Discussion
// 
// If the value is [AuthorizationStatusNotDetermined], you can request access
// by calling the [RequestAccessWithCompletionHandler] method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice/authorizationStatus
func (_AVExternalStorageDeviceClass AVExternalStorageDeviceClass) AuthorizationStatus() AVAuthorizationStatus {
	rv := objc.Send[AVAuthorizationStatus](objc.ID(_AVExternalStorageDeviceClass.class), objc.Sel("authorizationStatus"))
	return AVAuthorizationStatus(rv)
}

// RequestAccess is a synchronous wrapper around [AVExternalStorageDevice.RequestAccessWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (ec AVExternalStorageDeviceClass) RequestAccess(ctx context.Context) (bool, error) {
	done := make(chan bool, 1)
	ec.RequestAccessWithCompletionHandler(func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

