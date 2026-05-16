// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines a common interface for game input devices.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice
type GCDevice interface {
	objectivec.IObject

	// The manufacturer-provided name for the device, or the user’s name for the device.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevice/vendorName
	VendorName() string

	// The product category that identifies the type of controller.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevice/productCategory
	ProductCategory() string

	// The dispatch queue that the framework uses to call element value change handlers.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevice/handlerQueue
	HandlerQueue() dispatch.Queue

	// The device’s physical input profile, such as a controller’s extended gamepad.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevice/physicalInputProfile
	PhysicalInputProfile() IGCPhysicalInputProfile

	// The dispatch queue that the framework uses to call element value change handlers.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevice/handlerQueue
	SetHandlerQueue(value dispatch.Queue)
}

// GCDeviceObject wraps an existing Objective-C object that conforms to the GCDevice protocol.
type GCDeviceObject struct {
	objectivec.Object
}

func (o GCDeviceObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCDeviceObjectFromID constructs a [GCDeviceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCDeviceObjectFromID(id objc.ID) GCDeviceObject {
	return GCDeviceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The manufacturer-provided name for the device, or the user’s name for the
// device.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice/vendorName
func (o GCDeviceObject) VendorName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("vendorName"))
	return foundation.NSStringFromID(rv).String()
}

// The product category that identifies the type of controller.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice/productCategory
func (o GCDeviceObject) ProductCategory() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("productCategory"))
	return foundation.NSStringFromID(rv).String()
}

// The dispatch queue that the framework uses to call element value change
// handlers.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice/handlerQueue
func (o GCDeviceObject) HandlerQueue() dispatch.Queue {
	rv := objc.Send[uintptr](o.ID, objc.Sel("handlerQueue"))
	return dispatch.QueueFromHandle(rv)
}

// The device’s physical input profile, such as a controller’s extended
// gamepad.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice/physicalInputProfile
func (o GCDeviceObject) PhysicalInputProfile() IGCPhysicalInputProfile {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("physicalInputProfile"))
	return GCPhysicalInputProfileFromID(rv)
}

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
func (o GCDeviceObject) SetHandlerQueue(value dispatch.Queue) {
	objc.Send[struct{}](o.ID, objc.Sel("setHandlerQueue:"), value)
}
