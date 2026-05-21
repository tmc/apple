// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [GCDirectionalGamepad] class.
var (
	_GCDirectionalGamepadClass     GCDirectionalGamepadClass
	_GCDirectionalGamepadClassOnce sync.Once
)

func getGCDirectionalGamepadClass() GCDirectionalGamepadClass {
	_GCDirectionalGamepadClassOnce.Do(func() {
		_GCDirectionalGamepadClass = GCDirectionalGamepadClass{class: objc.GetClass("GCDirectionalGamepad")}
	})
	return _GCDirectionalGamepadClass
}

// GetGCDirectionalGamepadClass returns the class object for GCDirectionalGamepad.
func GetGCDirectionalGamepadClass() GCDirectionalGamepadClass {
	return getGCDirectionalGamepadClass()
}

type GCDirectionalGamepadClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCDirectionalGamepadClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCDirectionalGamepadClass) Alloc() GCDirectionalGamepad {
	rv := objc.Send[GCDirectionalGamepad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// A profile that supports only the directional pad, without motion or
// rotation.
//
// # Overview
//
// The directional gamepad profile is similar to a micro gamepad profile
// except it doesn’t support motion or rotation. The controller’s
// [GCController.Motion] property is `nil` and the inherited
// [GCMicroGamepad.AllowsRotation] property is false.
//
// If you select Micro Gamepad when you add the Game Controllers capability
// ([GCSupportedGameControllers] ) to your project, and you also support the
// GCDirectionalGamepad profile, select Directional Gamepad as well.
//
// If you support the second-generation Siri Remote and later, set the
// [GCSupportsMultipleMicroGamepads] key to [YES] in the information property
// list in your project.
//
// In addition, the directional pad element may report digital or analog
// values. If the directional pad’s [GCControllerElement.Analog] property is
// false, it reports absolute directional pad values (the
// [GCMicroGamepad.ReportsAbsoluteDpadValues] property is true).
//
// See: https://developer.apple.com/documentation/GameController/GCDirectionalGamepad
//
// [GCSupportedGameControllers]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/GCSupportedGameControllers
// [GCSupportsMultipleMicroGamepads]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/GCSupportsMultipleMicroGamepads
type GCDirectionalGamepad struct {
	GCMicroGamepad
}

// GCDirectionalGamepadFromID constructs a [GCDirectionalGamepad] from an objc.ID.
//
// A profile that supports only the directional pad, without motion or
// rotation.
func GCDirectionalGamepadFromID(id objc.ID) GCDirectionalGamepad {
	return GCDirectionalGamepad{GCMicroGamepad: GCMicroGamepadFromID(id)}
}

// NOTE: GCDirectionalGamepad adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCDirectionalGamepad] class.
//
// See: https://developer.apple.com/documentation/GameController/GCDirectionalGamepad
type IGCDirectionalGamepad interface {
	IGCMicroGamepad
}

// Init initializes the instance.
func (g GCDirectionalGamepad) Init() GCDirectionalGamepad {
	rv := objc.Send[GCDirectionalGamepad](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCDirectionalGamepad) Autorelease() GCDirectionalGamepad {
	rv := objc.Send[GCDirectionalGamepad](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCDirectionalGamepad creates a new GCDirectionalGamepad instance.
func NewGCDirectionalGamepad() GCDirectionalGamepad {
	class := getGCDirectionalGamepadClass()
	rv := objc.Send[GCDirectionalGamepad](objc.ID(class.class), objc.Sel("new"))
	return rv
}
