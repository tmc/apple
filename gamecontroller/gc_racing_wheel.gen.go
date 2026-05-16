// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GCRacingWheel] class.
var (
	_GCRacingWheelClass     GCRacingWheelClass
	_GCRacingWheelClassOnce sync.Once
)

func getGCRacingWheelClass() GCRacingWheelClass {
	_GCRacingWheelClassOnce.Do(func() {
		_GCRacingWheelClass = GCRacingWheelClass{class: objc.GetClass("GCRacingWheel")}
	})
	return _GCRacingWheelClass
}

// GetGCRacingWheelClass returns the class object for GCRacingWheel.
func GetGCRacingWheelClass() GCRacingWheelClass {
	return getGCRacingWheelClass()
}

type GCRacingWheelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCRacingWheelClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCRacingWheelClass) Alloc() GCRacingWheel {
	rv := objc.Send[GCRacingWheel](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a physical racing wheel controller connected to a
// device.
//
// # Discovering racing wheels
//
//   - [GCRacingWheel.GCRacingWheelDidConnect]: A notification that posts after a racing wheel controller connects to the device.
//   - [GCRacingWheel.GCRacingWheelDidDisconnect]: A notification that posts after a racing wheel controller disconnects from the device.
//
// # Getting events
//
//   - [GCRacingWheel.AcquireDeviceWithError]: Starts receiving events from the racing wheel.
//   - [GCRacingWheel.RelinquishDevice]: Stops receiving events from the racing wheel.
//   - [GCRacingWheel.Acquired]: A Boolean value that indicates whether the racing wheel sends events to the app.
//
// # Accessing the controller profile
//
//   - [GCRacingWheel.WheelInput]: The physical input profile for the racing wheel.
//
// # Creating snapshots
//
//   - [GCRacingWheel.Capture]: Returns a snapshot of the racing wheel with its current element values.
//   - [GCRacingWheel.Snapshot]: A Boolean value that indicates whether the object is a snapshot of a racing wheel.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheel
type GCRacingWheel struct {
	objectivec.Object
}

// GCRacingWheelFromID constructs a [GCRacingWheel] from an objc.ID.
//
// An object that represents a physical racing wheel controller connected to a
// device.
func GCRacingWheelFromID(id objc.ID) GCRacingWheel {
	return GCRacingWheel{objectivec.Object{ID: id}}
}

// NOTE: GCRacingWheel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCRacingWheel] class.
//
// # Discovering racing wheels
//
//   - [IGCRacingWheel.GCRacingWheelDidConnect]: A notification that posts after a racing wheel controller connects to the device.
//   - [IGCRacingWheel.GCRacingWheelDidDisconnect]: A notification that posts after a racing wheel controller disconnects from the device.
//
// # Getting events
//
//   - [IGCRacingWheel.AcquireDeviceWithError]: Starts receiving events from the racing wheel.
//   - [IGCRacingWheel.RelinquishDevice]: Stops receiving events from the racing wheel.
//   - [IGCRacingWheel.Acquired]: A Boolean value that indicates whether the racing wheel sends events to the app.
//
// # Accessing the controller profile
//
//   - [IGCRacingWheel.WheelInput]: The physical input profile for the racing wheel.
//
// # Creating snapshots
//
//   - [IGCRacingWheel.Capture]: Returns a snapshot of the racing wheel with its current element values.
//   - [IGCRacingWheel.Snapshot]: A Boolean value that indicates whether the object is a snapshot of a racing wheel.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheel
type IGCRacingWheel interface {
	objectivec.IObject
	GCDevice

	// Topic: Discovering racing wheels

	// A notification that posts after a racing wheel controller connects to the device.
	GCRacingWheelDidConnect() foundation.NSString
	// A notification that posts after a racing wheel controller disconnects from the device.
	GCRacingWheelDidDisconnect() foundation.NSString

	// Topic: Getting events

	// Starts receiving events from the racing wheel.
	AcquireDeviceWithError() (bool, error)
	// Stops receiving events from the racing wheel.
	RelinquishDevice()
	// A Boolean value that indicates whether the racing wheel sends events to the app.
	Acquired() bool

	// Topic: Accessing the controller profile

	// The physical input profile for the racing wheel.
	WheelInput() IGCRacingWheelInput

	// Topic: Creating snapshots

	// Returns a snapshot of the racing wheel with its current element values.
	Capture() IGCRacingWheel
	// A Boolean value that indicates whether the object is a snapshot of a racing wheel.
	Snapshot() bool
}

// Init initializes the instance.
func (g GCRacingWheel) Init() GCRacingWheel {
	rv := objc.Send[GCRacingWheel](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCRacingWheel) Autorelease() GCRacingWheel {
	rv := objc.Send[GCRacingWheel](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCRacingWheel creates a new GCRacingWheel instance.
func NewGCRacingWheel() GCRacingWheel {
	class := getGCRacingWheelClass()
	rv := objc.Send[GCRacingWheel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Starts receiving events from the racing wheel.
//
// # Discussion
//
// Before invoking this method, the racing wheel doesn’t deliver events to
// your app. Since only one app may receive racing wheel events at a time,
// this method can fail to acquire the device.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheel/acquireDevice()
func (g GCRacingWheel) AcquireDeviceWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](g.ID, objc.Sel("acquireDeviceWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("acquireDeviceWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// Stops receiving events from the racing wheel.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheel/relinquishDevice()
func (g GCRacingWheel) RelinquishDevice() {
	objc.Send[objc.ID](g.ID, objc.Sel("relinquishDevice"))
}

// Returns a snapshot of the racing wheel with its current element values.
//
// # Return Value
//
// A snapshot of the racing wheel.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheel/capture()
func (g GCRacingWheel) Capture() IGCRacingWheel {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("capture"))
	return GCRacingWheelFromID(rv)
}

// The device’s physical input profile, such as a controller’s extended
// gamepad.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice/physicalInputProfile
func (g GCRacingWheel) PhysicalInputProfile() IGCPhysicalInputProfile {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("physicalInputProfile"))
	return GCPhysicalInputProfileFromID(rv)
}

// A notification that posts after a racing wheel controller connects to the
// device.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/GCRacingWheelDidConnect
func (g GCRacingWheel) GCRacingWheelDidConnect() foundation.NSString {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("GCRacingWheelDidConnect"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// A notification that posts after a racing wheel controller disconnects from
// the device.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/GCRacingWheelDidDisconnect
func (g GCRacingWheel) GCRacingWheelDidDisconnect() foundation.NSString {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("GCRacingWheelDidDisconnect"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the racing wheel sends events to the
// app.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheel/isAcquired
func (g GCRacingWheel) Acquired() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isAcquired"))
	return rv
}

// The physical input profile for the racing wheel.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheel/wheelInput
func (g GCRacingWheel) WheelInput() IGCRacingWheelInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("wheelInput"))
	return GCRacingWheelInputFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the object is a snapshot of a racing
// wheel.
//
// # Discussion
//
// If true, the racing wheel is a snapshot at a moment in time of a real
// device; otherwise, it’s an actual racing wheel.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheel/isSnapshot
func (g GCRacingWheel) Snapshot() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isSnapshot"))
	return rv
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
func (g GCRacingWheel) HandlerQueue() dispatch.Queue {
	rv := objc.Send[uintptr](g.ID, objc.Sel("handlerQueue"))
	return dispatch.QueueFromHandle(rv)
}
func (g GCRacingWheel) SetHandlerQueue(value dispatch.Queue) {
	objc.Send[struct{}](g.ID, objc.Sel("setHandlerQueue:"), uintptr(value.Handle()))
}

// The product category that identifies the type of controller.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice/productCategory
func (g GCRacingWheel) ProductCategory() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("productCategory"))
	return foundation.NSStringFromID(rv).String()
}

// The manufacturer-provided name for the device, or the user’s name for the
// device.
//
// # Discussion
//
// The value of this property may be `nil` and may not be unique. Use this
// property to present information about the device to the user.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice/vendorName
func (g GCRacingWheel) VendorName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("vendorName"))
	return foundation.NSStringFromID(rv).String()
}

// The racing wheels connected to the device.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheel/connectedRacingWheels
func (_GCRacingWheelClass GCRacingWheelClass) ConnectedRacingWheels() foundation.INSSet {
	rv := objc.Send[objc.ID](objc.ID(_GCRacingWheelClass.class), objc.Sel("connectedRacingWheels"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// Protocol methods for GCDevice
