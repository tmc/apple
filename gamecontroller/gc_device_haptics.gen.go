// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GCDeviceHaptics] class.
var (
	_GCDeviceHapticsClass     GCDeviceHapticsClass
	_GCDeviceHapticsClassOnce sync.Once
)

func getGCDeviceHapticsClass() GCDeviceHapticsClass {
	_GCDeviceHapticsClassOnce.Do(func() {
		_GCDeviceHapticsClass = GCDeviceHapticsClass{class: objc.GetClass("GCDeviceHaptics")}
	})
	return _GCDeviceHapticsClass
}

// GetGCDeviceHapticsClass returns the class object for GCDeviceHaptics.
func GetGCDeviceHapticsClass() GCDeviceHapticsClass {
	return getGCDeviceHapticsClass()
}

type GCDeviceHapticsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCDeviceHapticsClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCDeviceHapticsClass) Alloc() GCDeviceHaptics {
	rv := objc.Send[GCDeviceHaptics](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// The locations of haptic actuators on a game controller.
//
// # Overview
//
// Use this class to create a haptic engine with a specified locality. Any
// patterns you send to that engine play on the specified actuators.
//
// # Creating a haptics engine
//
//   - [GCDeviceHaptics.CreateEngineWithLocality]: Creates a haptics engine with the specified locality.
//   - [GCDeviceHaptics.GCHapticDurationInfinite]: An infinite duration for a haptics event.
//
// # Getting the localities
//
//   - [GCDeviceHaptics.SupportedLocalities]: The locations of haptic actuators on the device.
//
// See: https://developer.apple.com/documentation/GameController/GCDeviceHaptics
type GCDeviceHaptics struct {
	objectivec.Object
}

// GCDeviceHapticsFromID constructs a [GCDeviceHaptics] from an objc.ID.
//
// The locations of haptic actuators on a game controller.
func GCDeviceHapticsFromID(id objc.ID) GCDeviceHaptics {
	return GCDeviceHaptics{objectivec.Object{ID: id}}
}

// NOTE: GCDeviceHaptics adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCDeviceHaptics] class.
//
// # Creating a haptics engine
//
//   - [IGCDeviceHaptics.CreateEngineWithLocality]: Creates a haptics engine with the specified locality.
//   - [IGCDeviceHaptics.GCHapticDurationInfinite]: An infinite duration for a haptics event.
//
// # Getting the localities
//
//   - [IGCDeviceHaptics.SupportedLocalities]: The locations of haptic actuators on the device.
//
// See: https://developer.apple.com/documentation/GameController/GCDeviceHaptics
type IGCDeviceHaptics interface {
	objectivec.IObject

	// Topic: Creating a haptics engine

	// Creates a haptics engine with the specified locality.
	CreateEngineWithLocality(locality GCHapticsLocality) unsafe.Pointer
	// An infinite duration for a haptics event.
	GCHapticDurationInfinite() float32

	// Topic: Getting the localities

	// The locations of haptic actuators on the device.
	SupportedLocalities() foundation.INSSet

	// A Boolean value that indicates whether the device supports haptic event playback.
	SupportsHaptics() bool
	SetSupportsHaptics(value bool)
}

// Init initializes the instance.
func (g GCDeviceHaptics) Init() GCDeviceHaptics {
	rv := objc.Send[GCDeviceHaptics](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCDeviceHaptics) Autorelease() GCDeviceHaptics {
	rv := objc.Send[GCDeviceHaptics](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCDeviceHaptics creates a new GCDeviceHaptics instance.
func NewGCDeviceHaptics() GCDeviceHaptics {
	class := getGCDeviceHapticsClass()
	rv := objc.Send[GCDeviceHaptics](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a haptics engine with the specified locality.
//
// locality: The location of the haptics on the controller.
//
// # Return Value
//
// A new haptics engine with the specified locality.
//
// # Discussion
//
// If you create an engine using the [default] location, users have the
// expected haptic experience. For example, the engine uses the handle
// accuators. If you want to create different experiences, such as using the
// left handle actuator as a woofer and the right actuator as a tweeter,
// create one or more engines with different localities.
//
// See: https://developer.apple.com/documentation/GameController/GCDeviceHaptics/createEngine(withLocality:)
//
// [default]: https://developer.apple.com/documentation/GameController/GCHapticsLocality/default
func (g GCDeviceHaptics) CreateEngineWithLocality(locality GCHapticsLocality) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("createEngineWithLocality:"), objc.String(string(locality)))
	return rv
}

// An infinite duration for a haptics event.
//
// See: https://developer.apple.com/documentation/gamecontroller/gchapticdurationinfinite
func (g GCDeviceHaptics) GCHapticDurationInfinite() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("GCHapticDurationInfinite"))
	return rv
}

// The locations of haptic actuators on the device.
//
// See: https://developer.apple.com/documentation/GameController/GCDeviceHaptics/supportedLocalities
func (g GCDeviceHaptics) SupportedLocalities() foundation.INSSet {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("supportedLocalities"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the device supports haptic event
// playback.
//
// See: https://developer.apple.com/documentation/CoreHaptics/CHHapticDeviceCapability/supportsHaptics
func (g GCDeviceHaptics) SupportsHaptics() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("supportsHaptics"))
	return rv
}
func (g GCDeviceHaptics) SetSupportsHaptics(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setSupportsHaptics:"), value)
}
