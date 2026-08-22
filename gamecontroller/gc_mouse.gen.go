// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GCMouse] class.
var (
	_GCMouseClass     GCMouseClass
	_GCMouseClassOnce sync.Once
)

func getGCMouseClass() GCMouseClass {
	_GCMouseClassOnce.Do(func() {
		_GCMouseClass = GCMouseClass{class: objc.GetClass("GCMouse")}
	})
	return _GCMouseClass
}

// GetGCMouseClass returns the class object for GCMouse.
func GetGCMouseClass() GCMouseClass {
	return getGCMouseClass()
}

type GCMouseClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCMouseClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCMouseClass) Alloc() GCMouse {
	rv := objc.Send[GCMouse](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a physical mouse connected to a device.
//
// # Overview
//
// To get a mouse object and its input values, register for the
// [GCMouseDidConnect] (Swift) or [GCMouseDidConnectNotification]
// (Objective-C) notification for when a mouse connects to the device. Then
// register for the [GCMouseDidBecomeCurrent] (Swift) or
// [GCMouseDidBecomeCurrentNotification] (Objective-C) notification for when
// it becomes the [GCMouseClass.Current] mouse. Alternatively, use the
// [GCMouseClass.Current] class property or the [GCMouseClass.Mice] class
// method to get a mouse object. Then get the current input values from the
// mouse object’s [GCMouse.MouseInput] controller profile.
//
// # Getting input values
//
//   - [GCMouse.MouseInput]: The controller profile for the mouse device.
//
// See: https://developer.apple.com/documentation/GameController/GCMouse
//
// [GCMouseDidBecomeCurrentNotification]: https://developer.apple.com/documentation/GameController/GCMouseDidBecomeCurrentNotification
// [GCMouseDidBecomeCurrent]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/GCMouseDidBecomeCurrent
// [GCMouseDidConnectNotification]: https://developer.apple.com/documentation/GameController/GCMouseDidConnectNotification
// [GCMouseDidConnect]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/GCMouseDidConnect
type GCMouse struct {
	objectivec.Object
}

// GCMouseFromID constructs a [GCMouse] from an objc.ID.
//
// An object that represents a physical mouse connected to a device.
func GCMouseFromID(id objc.ID) GCMouse {
	return GCMouse{objectivec.Object{ID: id}}
}

// NOTE: GCMouse adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCMouse] class.
//
// # Getting input values
//
//   - [IGCMouse.MouseInput]: The controller profile for the mouse device.
//
// See: https://developer.apple.com/documentation/GameController/GCMouse
type IGCMouse interface {
	objectivec.IObject

	// Topic: Getting input values

	// The controller profile for the mouse device.
	MouseInput() IGCMouseInput

	// The dispatch queue that the framework uses to call element value change handlers.
	HandlerQueue() dispatch.Queue
	// The device’s physical input profile, such as a controller’s extended gamepad.
	PhysicalInputProfile() IGCPhysicalInputProfile
	// The product category that identifies the type of controller.
	ProductCategory() string
	// The manufacturer-provided name for the device, or the user’s name for the device.
	VendorName() string
}

// Init initializes the instance.
func (g GCMouse) Init() GCMouse {
	rv := objc.Send[GCMouse](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCMouse) Autorelease() GCMouse {
	rv := objc.Send[GCMouse](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCMouse creates a new GCMouse instance.
func NewGCMouse() GCMouse {
	class := getGCMouseClass()
	rv := objc.Send[GCMouse](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The dispatch queue that the framework uses to call element value change
// handlers.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice/handlerQueue
func (g GCMouse) HandlerQueue() dispatch.Queue {
	rv := objc.Send[uintptr](g.ID, objc.Sel("handlerQueue"))
	return dispatch.QueueFromHandle(rv)
}

// The device’s physical input profile, such as a controller’s extended
// gamepad.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice/physicalInputProfile
func (g GCMouse) PhysicalInputProfile() IGCPhysicalInputProfile {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("physicalInputProfile"))
	return GCPhysicalInputProfileFromID(rv)
}

// The product category that identifies the type of controller.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice/productCategory
func (g GCMouse) ProductCategory() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("productCategory"))
	return foundation.NSStringFromID(rv).String()
}

// The manufacturer-provided name for the device, or the user’s name for the
// device.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice/vendorName
func (g GCMouse) VendorName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("vendorName"))
	return foundation.NSStringFromID(rv).String()
}

// Returns any mice that the user connects to the device.
//
// # Return Value
//
// The currently connected mouse devices.
//
// See: https://developer.apple.com/documentation/GameController/GCMouse/mice()
func (_GCMouseClass GCMouseClass) Mice() []GCMouse {
	rv := objc.Send[[]objc.ID](objc.ID(_GCMouseClass.class), objc.Sel("mice"))
	return objc.ConvertSlice(rv, func(id objc.ID) GCMouse {
		return GCMouseFromID(id)
	})
}

// The controller profile for the mouse device.
//
// # Discussion
//
// Get the mouse’s current state, and input values for its buttons and
// scroll wheel, from this object.
//
// See: https://developer.apple.com/documentation/GameController/GCMouse/mouseInput
func (g GCMouse) MouseInput() IGCMouseInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("mouseInput"))
	return GCMouseInputFromID(objc.ID(rv))
}

// The most recent mouse that the user connects.
//
// See: https://developer.apple.com/documentation/GameController/GCMouse/current
func (_GCMouseClass GCMouseClass) Current() GCMouse {
	rv := objc.Send[objc.ID](objc.ID(_GCMouseClass.class), objc.Sel("current"))
	return GCMouseFromID(objc.ID(rv))
}

// Protocol methods for GCDevice

// The dispatch queue that the framework uses to call element value change
// handlers.
//
// # Discussion
//
// The default queue is the main queue. Set this property to another queue to
// asynchronously call value change handlers (see [GCControllerAxisInput],
// [GCControllerButtonInput], [GCControllerDirectionPad], and [GCMotion]). For
// example, if you handle input on another queue, set this property when you
// first access the input device.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice/handlerQueue
func (o GCMouse) SetHandlerQueue(value dispatch.Queue) {
	objc.Send[struct{}](o.ID, objc.Sel("setHandlerQueue:"), value)
}
