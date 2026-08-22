// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MIDICIDeviceManager] class.
var (
	_MIDICIDeviceManagerClass     MIDICIDeviceManagerClass
	_MIDICIDeviceManagerClassOnce sync.Once
)

func getMIDICIDeviceManagerClass() MIDICIDeviceManagerClass {
	_MIDICIDeviceManagerClassOnce.Do(func() {
		_MIDICIDeviceManagerClass = MIDICIDeviceManagerClass{class: objc.GetClass("MIDICIDeviceManager")}
	})
	return _MIDICIDeviceManagerClass
}

// GetMIDICIDeviceManagerClass returns the class object for MIDICIDeviceManager.
func GetMIDICIDeviceManagerClass() MIDICIDeviceManagerClass {
	return getMIDICIDeviceManagerClass()
}

type MIDICIDeviceManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MIDICIDeviceManagerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MIDICIDeviceManagerClass) Alloc() MIDICIDeviceManager {
	rv := objc.Send[MIDICIDeviceManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MIDICIDeviceManager.DiscoveredCIDevices]
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceManager
type MIDICIDeviceManager struct {
	objectivec.Object
}

// MIDICIDeviceManagerFromID constructs a [MIDICIDeviceManager] from an objc.ID.
func MIDICIDeviceManagerFromID(id objc.ID) MIDICIDeviceManager {
	return MIDICIDeviceManager{objectivec.Object{ID: id}}
}

// NOTE: MIDICIDeviceManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MIDICIDeviceManager] class.
//
// # Instance Properties
//
//   - [IMIDICIDeviceManager.DiscoveredCIDevices]
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceManager
type IMIDICIDeviceManager interface {
	objectivec.IObject

	// Topic: Instance Properties

	DiscoveredCIDevices() []MIDICIDevice
}

// Init initializes the instance.
func (m MIDICIDeviceManager) Init() MIDICIDeviceManager {
	rv := objc.Send[MIDICIDeviceManager](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MIDICIDeviceManager) Autorelease() MIDICIDeviceManager {
	rv := objc.Send[MIDICIDeviceManager](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIDeviceManager creates a new MIDICIDeviceManager instance.
func NewMIDICIDeviceManager() MIDICIDeviceManager {
	class := getMIDICIDeviceManagerClass()
	rv := objc.Send[MIDICIDeviceManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceManager/discoveredCIDevices
func (m MIDICIDeviceManager) DiscoveredCIDevices() []MIDICIDevice {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("discoveredCIDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) MIDICIDevice {
		return MIDICIDeviceFromID(id)
	})
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceManager/shared
func (_MIDICIDeviceManagerClass MIDICIDeviceManagerClass) SharedInstance() MIDICIDeviceManager {
	rv := objc.Send[objc.ID](objc.ID(_MIDICIDeviceManagerClass.class), objc.Sel("sharedInstance"))
	return MIDICIDeviceManagerFromID(objc.ID(rv))
}
