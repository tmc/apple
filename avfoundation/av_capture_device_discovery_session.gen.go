// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureDeviceDiscoverySession] class.
var (
	_AVCaptureDeviceDiscoverySessionClass     AVCaptureDeviceDiscoverySessionClass
	_AVCaptureDeviceDiscoverySessionClassOnce sync.Once
)

func getAVCaptureDeviceDiscoverySessionClass() AVCaptureDeviceDiscoverySessionClass {
	_AVCaptureDeviceDiscoverySessionClassOnce.Do(func() {
		_AVCaptureDeviceDiscoverySessionClass = AVCaptureDeviceDiscoverySessionClass{class: objc.GetClass("AVCaptureDeviceDiscoverySession")}
	})
	return _AVCaptureDeviceDiscoverySessionClass
}

// GetAVCaptureDeviceDiscoverySessionClass returns the class object for AVCaptureDeviceDiscoverySession.
func GetAVCaptureDeviceDiscoverySessionClass() AVCaptureDeviceDiscoverySessionClass {
	return getAVCaptureDeviceDiscoverySessionClass()
}

type AVCaptureDeviceDiscoverySessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureDeviceDiscoverySessionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureDeviceDiscoverySessionClass) Alloc() AVCaptureDeviceDiscoverySession {
	rv := objc.Send[AVCaptureDeviceDiscoverySession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that finds capture devices that match specific search criteria.
//
// # Overview
// 
// After creating a device discovery session, query its [AVCaptureDeviceDiscoverySession.Devices] property to
// find a device to use for capture. You can also key-value observe this
// property to monitor changes to the list of available devices.
//
// # Finding devices
//
//   - [AVCaptureDeviceDiscoverySession.Devices]: A list of devices that match the search criteria of the discovery session.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DiscoverySession
type AVCaptureDeviceDiscoverySession struct {
	objectivec.Object
}

// AVCaptureDeviceDiscoverySessionFromID constructs a [AVCaptureDeviceDiscoverySession] from an objc.ID.
//
// An object that finds capture devices that match specific search criteria.
func AVCaptureDeviceDiscoverySessionFromID(id objc.ID) AVCaptureDeviceDiscoverySession {
	return AVCaptureDeviceDiscoverySession{objectivec.Object{ID: id}}
}
// NOTE: AVCaptureDeviceDiscoverySession adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureDeviceDiscoverySession] class.
//
// # Finding devices
//
//   - [IAVCaptureDeviceDiscoverySession.Devices]: A list of devices that match the search criteria of the discovery session.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DiscoverySession
type IAVCaptureDeviceDiscoverySession interface {
	objectivec.IObject

	// Topic: Finding devices

	// A list of devices that match the search criteria of the discovery session.
	Devices() []AVCaptureDevice
}

// Init initializes the instance.
func (c AVCaptureDeviceDiscoverySession) Init() AVCaptureDeviceDiscoverySession {
	rv := objc.Send[AVCaptureDeviceDiscoverySession](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureDeviceDiscoverySession) Autorelease() AVCaptureDeviceDiscoverySession {
	rv := objc.Send[AVCaptureDeviceDiscoverySession](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureDeviceDiscoverySession creates a new AVCaptureDeviceDiscoverySession instance.
func NewAVCaptureDeviceDiscoverySession() AVCaptureDeviceDiscoverySession {
	class := getAVCaptureDeviceDiscoverySessionClass()
	rv := objc.Send[AVCaptureDeviceDiscoverySession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a discovery session that finds devices that match the specified
// criteria.
//
// deviceTypes: A list of device types to search for, such as [builtInWideAngleCamera] and
// [builtInMicrophone]. The array must contain at least one valid
// [AVCaptureDeviceType] value.
// //
// [builtInMicrophone]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DeviceType-swift.struct/builtInMicrophone
// [builtInWideAngleCamera]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DeviceType-swift.struct/builtInWideAngleCamera
//
// mediaType: The media type to capture, such as [video] or [audio]. Pass `nil` to search
// for devices regardless of supported media types.
// //
// [audio]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/audio
// [video]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/video
//
// position: The position of capture device to search for, relative to system hardware
// (front- or back-facing). Pass [CaptureDevicePositionUnspecified] to search
// for devices regardless of position.
//
// # Return Value
// 
// A new discovery session.
//
// # Discussion
// 
// After creating a discovery session, query its [Devices] property to examine
// matching devices and choose one for capture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DiscoverySession/init(deviceTypes:mediaType:position:)
func NewCaptureDeviceDiscoverySessionWithDeviceTypesMediaTypePosition(deviceTypes []string, mediaType AVMediaType, position AVCaptureDevicePosition) AVCaptureDeviceDiscoverySession {
	rv := objc.Send[objc.ID](objc.ID(getAVCaptureDeviceDiscoverySessionClass().class), objc.Sel("discoverySessionWithDeviceTypes:mediaType:position:"), objectivec.StringSliceToNSArray(deviceTypes), objc.String(string(mediaType)), position)
	return AVCaptureDeviceDiscoverySessionFromID(rv)
}

// A list of devices that match the search criteria of the discovery session.
//
// # Discussion
// 
// Querying this property provides an array devices currently available on the
// system. The system sorts the device list according to the order you
// specified when you created the discovery session. If you created the
// session with a position of [CaptureDevicePositionUnspecified], the system
// further sorts them by position in the [AVCaptureDevice.Position]
// enumeration.
// 
// Key-value observe this property to monitor changes to the device list.
//
// [AVCaptureDevice.Position]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Position-swift.enum
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DiscoverySession/devices
func (c AVCaptureDeviceDiscoverySession) Devices() []AVCaptureDevice {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("devices"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaptureDevice {
		return AVCaptureDeviceFromID(id)
	})
}

// A notification the system posts when a new capture device becomes
// available.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/wasconnectednotification
func (_AVCaptureDeviceDiscoverySessionClass AVCaptureDeviceDiscoverySessionClass) WasConnectedNotification() foundation.NSString {
	rv := objc.Send[objc.ID](objc.ID(_AVCaptureDeviceDiscoverySessionClass.class), objc.Sel("AVCaptureDeviceWasConnectedNotification"))
	return foundation.NSStringFromID(objc.ID(rv))
}
// A notification the system posts when an existing device becomes
// unavailable.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/wasdisconnectednotification
func (_AVCaptureDeviceDiscoverySessionClass AVCaptureDeviceDiscoverySessionClass) WasDisconnectedNotification() foundation.NSString {
	rv := objc.Send[objc.ID](objc.ID(_AVCaptureDeviceDiscoverySessionClass.class), objc.Sel("AVCaptureDeviceWasDisconnectedNotification"))
	return foundation.NSStringFromID(objc.ID(rv))
}

