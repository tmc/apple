// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MIDIUMPEndpointManager] class.
var (
	_MIDIUMPEndpointManagerClass     MIDIUMPEndpointManagerClass
	_MIDIUMPEndpointManagerClassOnce sync.Once
)

func getMIDIUMPEndpointManagerClass() MIDIUMPEndpointManagerClass {
	_MIDIUMPEndpointManagerClassOnce.Do(func() {
		_MIDIUMPEndpointManagerClass = MIDIUMPEndpointManagerClass{class: objc.GetClass("MIDIUMPEndpointManager")}
	})
	return _MIDIUMPEndpointManagerClass
}

// GetMIDIUMPEndpointManagerClass returns the class object for MIDIUMPEndpointManager.
func GetMIDIUMPEndpointManagerClass() MIDIUMPEndpointManagerClass {
	return getMIDIUMPEndpointManagerClass()
}

type MIDIUMPEndpointManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MIDIUMPEndpointManagerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MIDIUMPEndpointManagerClass) Alloc() MIDIUMPEndpointManager {
	rv := objc.Send[MIDIUMPEndpointManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [MIDIUMPEndpointManager.UMPEndpoints]
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpointManager
type MIDIUMPEndpointManager struct {
	objectivec.Object
}

// MIDIUMPEndpointManagerFromID constructs a [MIDIUMPEndpointManager] from an objc.ID.
func MIDIUMPEndpointManagerFromID(id objc.ID) MIDIUMPEndpointManager {
	return MIDIUMPEndpointManager{objectivec.Object{ID: id}}
}

// NOTE: MIDIUMPEndpointManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MIDIUMPEndpointManager] class.
//
// # Instance Properties
//
//   - [IMIDIUMPEndpointManager.UMPEndpoints]
//
// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpointManager
type IMIDIUMPEndpointManager interface {
	objectivec.IObject

	// Topic: Instance Properties

	UMPEndpoints() []MIDIUMPEndpoint
}

// Init initializes the instance.
func (m MIDIUMPEndpointManager) Init() MIDIUMPEndpointManager {
	rv := objc.Send[MIDIUMPEndpointManager](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MIDIUMPEndpointManager) Autorelease() MIDIUMPEndpointManager {
	rv := objc.Send[MIDIUMPEndpointManager](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIUMPEndpointManager creates a new MIDIUMPEndpointManager instance.
func NewMIDIUMPEndpointManager() MIDIUMPEndpointManager {
	class := getMIDIUMPEndpointManagerClass()
	rv := objc.Send[MIDIUMPEndpointManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpointManager/umpEndpoints
func (m MIDIUMPEndpointManager) UMPEndpoints() []MIDIUMPEndpoint {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("UMPEndpoints"))
	return objc.ConvertSlice(rv, func(id objc.ID) MIDIUMPEndpoint {
		return MIDIUMPEndpointFromID(id)
	})
}

// See: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPEndpointManager/shared
func (_MIDIUMPEndpointManagerClass MIDIUMPEndpointManagerClass) SharedInstance() MIDIUMPEndpointManager {
	rv := objc.Send[objc.ID](objc.ID(_MIDIUMPEndpointManagerClass.class), objc.Sel("sharedInstance"))
	return MIDIUMPEndpointManagerFromID(objc.ID(rv))
}
