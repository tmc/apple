// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GCDeviceBattery] class.
var (
	_GCDeviceBatteryClass     GCDeviceBatteryClass
	_GCDeviceBatteryClassOnce sync.Once
)

func getGCDeviceBatteryClass() GCDeviceBatteryClass {
	_GCDeviceBatteryClassOnce.Do(func() {
		_GCDeviceBatteryClass = GCDeviceBatteryClass{class: objc.GetClass("GCDeviceBattery")}
	})
	return _GCDeviceBatteryClass
}

// GetGCDeviceBatteryClass returns the class object for GCDeviceBattery.
func GetGCDeviceBatteryClass() GCDeviceBatteryClass {
	return getGCDeviceBatteryClass()
}

type GCDeviceBatteryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCDeviceBatteryClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCDeviceBatteryClass) Alloc() GCDeviceBattery {
	rv := objc.Send[GCDeviceBattery](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// The charge level and state of a device’s battery.
//
// # Overview
//
// Use this class to display the state of a device’s battery to a player.
//
// # Getting the battery level and state
//
//   - [GCDeviceBattery.BatteryLevel]: The charge level of a device’s battery.
//   - [GCDeviceBattery.BatteryState]: The state of a device’s battery.
//
// See: https://developer.apple.com/documentation/GameController/GCDeviceBattery
type GCDeviceBattery struct {
	objectivec.Object
}

// GCDeviceBatteryFromID constructs a [GCDeviceBattery] from an objc.ID.
//
// The charge level and state of a device’s battery.
func GCDeviceBatteryFromID(id objc.ID) GCDeviceBattery {
	return GCDeviceBattery{objectivec.Object{ID: id}}
}

// NOTE: GCDeviceBattery adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCDeviceBattery] class.
//
// # Getting the battery level and state
//
//   - [IGCDeviceBattery.BatteryLevel]: The charge level of a device’s battery.
//   - [IGCDeviceBattery.BatteryState]: The state of a device’s battery.
//
// See: https://developer.apple.com/documentation/GameController/GCDeviceBattery
type IGCDeviceBattery interface {
	objectivec.IObject

	// Topic: Getting the battery level and state

	// The charge level of a device’s battery.
	BatteryLevel() float32
	// The state of a device’s battery.
	BatteryState() GCDeviceBatteryState
}

// Init initializes the instance.
func (g GCDeviceBattery) Init() GCDeviceBattery {
	rv := objc.Send[GCDeviceBattery](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCDeviceBattery) Autorelease() GCDeviceBattery {
	rv := objc.Send[GCDeviceBattery](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCDeviceBattery creates a new GCDeviceBattery instance.
func NewGCDeviceBattery() GCDeviceBattery {
	class := getGCDeviceBatteryClass()
	rv := objc.Send[GCDeviceBattery](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The charge level of a device’s battery.
//
// # Discussion
//
// The battery level is a percentage ranging from `0.0` (fully discharged) to
// `1.0` (100% charged). The default value for this property is `0.0`.
//
// See: https://developer.apple.com/documentation/GameController/GCDeviceBattery/batteryLevel
func (g GCDeviceBattery) BatteryLevel() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("batteryLevel"))
	return rv
}

// The state of a device’s battery.
//
// See: https://developer.apple.com/documentation/GameController/GCDeviceBattery/batteryState
func (g GCDeviceBattery) BatteryState() GCDeviceBatteryState {
	rv := objc.Send[GCDeviceBatteryState](g.ID, objc.Sel("batteryState"))
	return GCDeviceBatteryState(rv)
}
