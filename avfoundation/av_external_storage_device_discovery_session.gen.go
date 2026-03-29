// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVExternalStorageDeviceDiscoverySession] class.
var (
	_AVExternalStorageDeviceDiscoverySessionClass     AVExternalStorageDeviceDiscoverySessionClass
	_AVExternalStorageDeviceDiscoverySessionClassOnce sync.Once
)

func getAVExternalStorageDeviceDiscoverySessionClass() AVExternalStorageDeviceDiscoverySessionClass {
	_AVExternalStorageDeviceDiscoverySessionClassOnce.Do(func() {
		_AVExternalStorageDeviceDiscoverySessionClass = AVExternalStorageDeviceDiscoverySessionClass{class: objc.GetClass("AVExternalStorageDeviceDiscoverySession")}
	})
	return _AVExternalStorageDeviceDiscoverySessionClass
}

// GetAVExternalStorageDeviceDiscoverySessionClass returns the class object for AVExternalStorageDeviceDiscoverySession.
func GetAVExternalStorageDeviceDiscoverySessionClass() AVExternalStorageDeviceDiscoverySessionClass {
	return getAVExternalStorageDeviceDiscoverySessionClass()
}

type AVExternalStorageDeviceDiscoverySessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVExternalStorageDeviceDiscoverySessionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVExternalStorageDeviceDiscoverySessionClass) Alloc() AVExternalStorageDeviceDiscoverySession {
	rv := objc.Send[AVExternalStorageDeviceDiscoverySession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// Informs your app when the external storage devices connect to and
// disconnect from the system.
//
// # Monitoring for storage device updates
//
//   - [AVExternalStorageDeviceDiscoverySession.ExternalStorageDevices]: An array of external storage devices the session updates as individual devices connect or disconnect from the system.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDeviceDiscoverySession
type AVExternalStorageDeviceDiscoverySession struct {
	objectivec.Object
}

// AVExternalStorageDeviceDiscoverySessionFromID constructs a [AVExternalStorageDeviceDiscoverySession] from an objc.ID.
//
// Informs your app when the external storage devices connect to and
// disconnect from the system.
func AVExternalStorageDeviceDiscoverySessionFromID(id objc.ID) AVExternalStorageDeviceDiscoverySession {
	return AVExternalStorageDeviceDiscoverySession{objectivec.Object{ID: id}}
}
// NOTE: AVExternalStorageDeviceDiscoverySession adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVExternalStorageDeviceDiscoverySession] class.
//
// # Monitoring for storage device updates
//
//   - [IAVExternalStorageDeviceDiscoverySession.ExternalStorageDevices]: An array of external storage devices the session updates as individual devices connect or disconnect from the system.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDeviceDiscoverySession
type IAVExternalStorageDeviceDiscoverySession interface {
	objectivec.IObject

	// Topic: Monitoring for storage device updates

	// An array of external storage devices the session updates as individual devices connect or disconnect from the system.
	ExternalStorageDevices() []AVExternalStorageDevice
}

// Init initializes the instance.
func (e AVExternalStorageDeviceDiscoverySession) Init() AVExternalStorageDeviceDiscoverySession {
	rv := objc.Send[AVExternalStorageDeviceDiscoverySession](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e AVExternalStorageDeviceDiscoverySession) Autorelease() AVExternalStorageDeviceDiscoverySession {
	rv := objc.Send[AVExternalStorageDeviceDiscoverySession](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVExternalStorageDeviceDiscoverySession creates a new AVExternalStorageDeviceDiscoverySession instance.
func NewAVExternalStorageDeviceDiscoverySession() AVExternalStorageDeviceDiscoverySession {
	class := getAVExternalStorageDeviceDiscoverySessionClass()
	rv := objc.Send[AVExternalStorageDeviceDiscoverySession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An array of external storage devices the session updates as individual
// devices connect or disconnect from the system.
//
// # Discussion
// 
// Your app can monitor the changes to this array with a key-value
// observation.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDeviceDiscoverySession/externalStorageDevices
func (e AVExternalStorageDeviceDiscoverySession) ExternalStorageDevices() []AVExternalStorageDevice {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("externalStorageDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVExternalStorageDevice {
		return AVExternalStorageDeviceFromID(id)
	})
}

// A Boolean value that indicates whether the system supports external storage
// devices.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDeviceDiscoverySession/isSupported
func (_AVExternalStorageDeviceDiscoverySessionClass AVExternalStorageDeviceDiscoverySessionClass) Supported() bool {
	rv := objc.Send[bool](objc.ID(_AVExternalStorageDeviceDiscoverySessionClass.class), objc.Sel("isSupported"))
	return rv
}
// The system’s singleton device discovery session instance.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDeviceDiscoverySession/shared
func (_AVExternalStorageDeviceDiscoverySessionClass AVExternalStorageDeviceDiscoverySessionClass) SharedSession() AVExternalStorageDeviceDiscoverySession {
	rv := objc.Send[objc.ID](objc.ID(_AVExternalStorageDeviceDiscoverySessionClass.class), objc.Sel("sharedSession"))
	return AVExternalStorageDeviceDiscoverySessionFromID(objc.ID(rv))
}

