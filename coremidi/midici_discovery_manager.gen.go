// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MIDICIDiscoveryManager] class.
var (
	_MIDICIDiscoveryManagerClass     MIDICIDiscoveryManagerClass
	_MIDICIDiscoveryManagerClassOnce sync.Once
)

func getMIDICIDiscoveryManagerClass() MIDICIDiscoveryManagerClass {
	_MIDICIDiscoveryManagerClassOnce.Do(func() {
		_MIDICIDiscoveryManagerClass = MIDICIDiscoveryManagerClass{class: objc.GetClass("MIDICIDiscoveryManager")}
	})
	return _MIDICIDiscoveryManagerClass
}

// GetMIDICIDiscoveryManagerClass returns the class object for MIDICIDiscoveryManager.
func GetMIDICIDiscoveryManagerClass() MIDICIDiscoveryManagerClass {
	return getMIDICIDiscoveryManagerClass()
}

type MIDICIDiscoveryManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MIDICIDiscoveryManagerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MIDICIDiscoveryManagerClass) Alloc() MIDICIDiscoveryManager {
	rv := objc.Send[MIDICIDiscoveryManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A singleton object that performs systemwide MIDI-CI discovery.
//
// # Overview
//
// Use this class to retrieve information about MIDI-CI–capable nodes in the
// MIDI subsystem. You can create [MIDICISession] objects only from the
// destinations discovered using this API.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveryManager
type MIDICIDiscoveryManager struct {
	objectivec.Object
}

// MIDICIDiscoveryManagerFromID constructs a [MIDICIDiscoveryManager] from an objc.ID.
//
// A singleton object that performs systemwide MIDI-CI discovery.
func MIDICIDiscoveryManagerFromID(id objc.ID) MIDICIDiscoveryManager {
	return MIDICIDiscoveryManager{objectivec.Object{ID: id}}
}

// NOTE: MIDICIDiscoveryManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MIDICIDiscoveryManager] class.
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveryManager
type IMIDICIDiscoveryManager interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MIDICIDiscoveryManager) Init() MIDICIDiscoveryManager {
	rv := objc.Send[MIDICIDiscoveryManager](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MIDICIDiscoveryManager) Autorelease() MIDICIDiscoveryManager {
	rv := objc.Send[MIDICIDiscoveryManager](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIDiscoveryManager creates a new MIDICIDiscoveryManager instance.
func NewMIDICIDiscoveryManager() MIDICIDiscoveryManager {
	class := getMIDICIDiscoveryManagerClass()
	rv := objc.Send[MIDICIDiscoveryManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}
