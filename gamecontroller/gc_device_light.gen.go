// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GCDeviceLight] class.
var (
	_GCDeviceLightClass     GCDeviceLightClass
	_GCDeviceLightClassOnce sync.Once
)

func getGCDeviceLightClass() GCDeviceLightClass {
	_GCDeviceLightClassOnce.Do(func() {
		_GCDeviceLightClass = GCDeviceLightClass{class: objc.GetClass("GCDeviceLight")}
	})
	return _GCDeviceLightClass
}

// GetGCDeviceLightClass returns the class object for GCDeviceLight.
func GetGCDeviceLightClass() GCDeviceLightClass {
	return getGCDeviceLightClass()
}

type GCDeviceLightClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCDeviceLightClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCDeviceLightClass) Alloc() GCDeviceLight {
	rv := objc.Send[GCDeviceLight](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// The colored light on a device.
//
// # Getting the light’s color
//
//   - [GCDeviceLight.Color]: The color of a device’s light.
//   - [GCDeviceLight.SetColor]
//
// See: https://developer.apple.com/documentation/GameController/GCDeviceLight
type GCDeviceLight struct {
	objectivec.Object
}

// GCDeviceLightFromID constructs a [GCDeviceLight] from an objc.ID.
//
// The colored light on a device.
func GCDeviceLightFromID(id objc.ID) GCDeviceLight {
	return GCDeviceLight{objectivec.Object{ID: id}}
}

// NOTE: GCDeviceLight adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCDeviceLight] class.
//
// # Getting the light’s color
//
//   - [IGCDeviceLight.Color]: The color of a device’s light.
//   - [IGCDeviceLight.SetColor]
//
// See: https://developer.apple.com/documentation/GameController/GCDeviceLight
type IGCDeviceLight interface {
	objectivec.IObject

	// Topic: Getting the light’s color

	// The color of a device’s light.
	Color() IGCColor
	SetColor(value IGCColor)
}

// Init initializes the instance.
func (g GCDeviceLight) Init() GCDeviceLight {
	rv := objc.Send[GCDeviceLight](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCDeviceLight) Autorelease() GCDeviceLight {
	rv := objc.Send[GCDeviceLight](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCDeviceLight creates a new GCDeviceLight instance.
func NewGCDeviceLight() GCDeviceLight {
	class := getGCDeviceLightClass()
	rv := objc.Send[GCDeviceLight](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The color of a device’s light.
//
// See: https://developer.apple.com/documentation/GameController/GCDeviceLight/color
func (g GCDeviceLight) Color() IGCColor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("color"))
	return GCColorFromID(objc.ID(rv))
}
func (g GCDeviceLight) SetColor(value IGCColor) {
	objc.Send[struct{}](g.ID, objc.Sel("setColor:"), value)
}
