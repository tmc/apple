// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The common properties for physical devices with elements.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState
type GCDevicePhysicalInputState interface {
	objectivec.IObject

	// Returns the element that the key specifies.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/subscript(_:)
	ObjectForKeyedSubscript(key string) GCPhysicalInputElement

	// The physical device that this profile represents.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/device
	Device() GCDevice

	// The time of the most recent event.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/lastEventTimestamp
	LastEventTimestamp() foundation.NSTimeInterval

	// The time in seconds between the last event and the current time.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/lastEventLatency
	LastEventLatency() foundation.NSTimeInterval

	// The device’s elements as key-value pairs for lookup by name.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/elements-1shp2
	Elements() IGCPhysicalInputElementCollection

	// The device’s axes as key-value pairs for lookup by name.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/axes-80rx
	Axes() IGCPhysicalInputElementCollection

	// The device’s buttons as key-value pairs for lookup by name.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/buttons-3257g
	Buttons() IGCPhysicalInputElementCollection

	// The device’s directional pads as key-value pairs for lookup by name.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/dpads-5yr9x
	Dpads() IGCPhysicalInputElementCollection

	// The device’s switches as key-value pairs for lookup by name.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/switches-6bws2
	Switches() IGCPhysicalInputElementCollection
}

// GCDevicePhysicalInputStateObject wraps an existing Objective-C object that conforms to the GCDevicePhysicalInputState protocol.
type GCDevicePhysicalInputStateObject struct {
	objectivec.Object
}

func (o GCDevicePhysicalInputStateObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCDevicePhysicalInputStateObjectFromID constructs a [GCDevicePhysicalInputStateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCDevicePhysicalInputStateObjectFromID(id objc.ID) GCDevicePhysicalInputStateObject {
	return GCDevicePhysicalInputStateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the element that the key specifies.
//
// key: A key that identifies an element.
//
// # Return Value
//
// The element that matches the key.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/subscript(_:)
func (o GCDevicePhysicalInputStateObject) ObjectForKeyedSubscript(key string) GCPhysicalInputElement {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("objectForKeyedSubscript:"), objc.String(key))
	return GCPhysicalInputElementObjectFromID(rv)
}

// The physical device that this profile represents.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/device
func (o GCDevicePhysicalInputStateObject) Device() GCDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return GCDeviceObjectFromID(rv)
}

// The time of the most recent event.
//
// # Discussion
//
// This property isn’t relative to any specific date and time. To determine
// the time between events, subtract a previous value of this property from
// the current value. You can also compare [LastEventTimestamp] properties of
// two different devices to determine which event occurs first.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/lastEventTimestamp
func (o GCDevicePhysicalInputStateObject) LastEventTimestamp() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](o.ID, objc.Sel("lastEventTimestamp"))
	return foundation.NSTimeInterval(rv)
}

// The time in seconds between the last event and the current time.
//
// # Discussion
//
// Use this property as a minimum latency value that may not include latency
// that accrues on the device or when it transmits the event. If the host goes
// to sleep between when the event occurs and when you get this property, the
// value may not be accurate.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/lastEventLatency
func (o GCDevicePhysicalInputStateObject) LastEventLatency() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](o.ID, objc.Sel("lastEventLatency"))
	return foundation.NSTimeInterval(rv)
}

// The device’s elements as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/elements-1shp2
func (o GCDevicePhysicalInputStateObject) Elements() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("elements"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s axes as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/axes-80rx
func (o GCDevicePhysicalInputStateObject) Axes() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("axes"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s buttons as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/buttons-3257g
func (o GCDevicePhysicalInputStateObject) Buttons() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("buttons"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s directional pads as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/dpads-5yr9x
func (o GCDevicePhysicalInputStateObject) Dpads() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("dpads"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s switches as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/switches-6bws2
func (o GCDevicePhysicalInputStateObject) Switches() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("switches"))
	return GCPhysicalInputElementCollectionFromID(rv)
}
