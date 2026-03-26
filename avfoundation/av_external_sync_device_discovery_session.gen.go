// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVExternalSyncDeviceDiscoverySession] class.
var (
	_AVExternalSyncDeviceDiscoverySessionClass     AVExternalSyncDeviceDiscoverySessionClass
	_AVExternalSyncDeviceDiscoverySessionClassOnce sync.Once
)

func getAVExternalSyncDeviceDiscoverySessionClass() AVExternalSyncDeviceDiscoverySessionClass {
	_AVExternalSyncDeviceDiscoverySessionClassOnce.Do(func() {
		_AVExternalSyncDeviceDiscoverySessionClass = AVExternalSyncDeviceDiscoverySessionClass{class: objc.GetClass("AVExternalSyncDeviceDiscoverySession")}
	})
	return _AVExternalSyncDeviceDiscoverySessionClass
}

// GetAVExternalSyncDeviceDiscoverySessionClass returns the class object for AVExternalSyncDeviceDiscoverySession.
func GetAVExternalSyncDeviceDiscoverySessionClass() AVExternalSyncDeviceDiscoverySessionClass {
	return getAVExternalSyncDeviceDiscoverySessionClass()
}

type AVExternalSyncDeviceDiscoverySessionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVExternalSyncDeviceDiscoverySessionClass) Alloc() AVExternalSyncDeviceDiscoverySession {
	rv := objc.Send[AVExternalSyncDeviceDiscoverySession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A means of discovering and monitoring connection / disconnection of
// external sync devices to the host.
//
// # Overview
// 
// [AVExternalSyncDeviceDiscoverySession] is a singleton that lists the
// external sync devices connected to the host. The client is expected to
// key-value observe the [AVExternalSyncDeviceDiscoverySession.Devices] property for changes to the external sync
// devices list.
//
// # Finding devices
//
//   - [AVExternalSyncDeviceDiscoverySession.Devices]: An array of external sync devices connected to this host.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/DiscoverySession
type AVExternalSyncDeviceDiscoverySession struct {
	objectivec.Object
}

// AVExternalSyncDeviceDiscoverySessionFromID constructs a [AVExternalSyncDeviceDiscoverySession] from an objc.ID.
//
// A means of discovering and monitoring connection / disconnection of
// external sync devices to the host.
func AVExternalSyncDeviceDiscoverySessionFromID(id objc.ID) AVExternalSyncDeviceDiscoverySession {
	return AVExternalSyncDeviceDiscoverySession{objectivec.Object{ID: id}}
}
// NOTE: AVExternalSyncDeviceDiscoverySession adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVExternalSyncDeviceDiscoverySession] class.
//
// # Finding devices
//
//   - [IAVExternalSyncDeviceDiscoverySession.Devices]: An array of external sync devices connected to this host.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/DiscoverySession
type IAVExternalSyncDeviceDiscoverySession interface {
	objectivec.IObject

	// Topic: Finding devices

	// An array of external sync devices connected to this host.
	Devices() []AVExternalSyncDevice
}

// Init initializes the instance.
func (e AVExternalSyncDeviceDiscoverySession) Init() AVExternalSyncDeviceDiscoverySession {
	rv := objc.Send[AVExternalSyncDeviceDiscoverySession](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e AVExternalSyncDeviceDiscoverySession) Autorelease() AVExternalSyncDeviceDiscoverySession {
	rv := objc.Send[AVExternalSyncDeviceDiscoverySession](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVExternalSyncDeviceDiscoverySession creates a new AVExternalSyncDeviceDiscoverySession instance.
func NewAVExternalSyncDeviceDiscoverySession() AVExternalSyncDeviceDiscoverySession {
	class := getAVExternalSyncDeviceDiscoverySessionClass()
	rv := objc.Send[AVExternalSyncDeviceDiscoverySession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An array of external sync devices connected to this host.
//
// # Discussion
// 
// The list is updated when external sync devices are connected to the host
// and they remain in the list until they become unavailable. This property is
// key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/DiscoverySession/devices
func (e AVExternalSyncDeviceDiscoverySession) Devices() []AVExternalSyncDevice {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("devices"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVExternalSyncDevice {
		return AVExternalSyncDeviceFromID(id)
	})
}

// The singleton instance of the external sync source device discovery
// session.
//
// # Discussion
// 
// Access the one and only external sync device discovery session on this host
// device using this method. `sharedSession` returns `nil` if the host device
// doesn’t support external sync devices.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/DiscoverySession/shared
func (_AVExternalSyncDeviceDiscoverySessionClass AVExternalSyncDeviceDiscoverySessionClass) SharedSession() AVExternalSyncDeviceDiscoverySession {
	rv := objc.Send[objc.ID](objc.ID(_AVExternalSyncDeviceDiscoverySessionClass.class), objc.Sel("sharedSession"))
	return AVExternalSyncDeviceDiscoverySessionFromID(objc.ID(rv))
}
// Whether external sync devices are supported by this device.
//
// # Discussion
// 
// A value of `true` indicates that external sync devices are supported while
// `false` indicates they are not.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDevice/DiscoverySession/isSupported
func (_AVExternalSyncDeviceDiscoverySessionClass AVExternalSyncDeviceDiscoverySessionClass) Supported() bool {
	rv := objc.Send[bool](objc.ID(_AVExternalSyncDeviceDiscoverySessionClass.class), objc.Sel("isSupported"))
	return rv
}

