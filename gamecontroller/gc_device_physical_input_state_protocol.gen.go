// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The common properties for physical devices with elements.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState
type GCDevicePhysicalInputState interface {
	objectivec.IObject

	// The physical device that this profile represents.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/device
	Device() GCDevice

	// The time of the most recent event.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/lastEventTimestamp
	LastEventTimestamp() float64

	// The time in seconds between the last event and the current time.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/lastEventLatency
	LastEventLatency() float64

	// Returns the element that the key specifies.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/subscript(_:)
	ObjectForKeyedSubscript(key string) GCPhysicalInputElement
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

// The physical device that this profile represents.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/device
func (o GCDevicePhysicalInputStateObject) Device() GCDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return GCDeviceObjectFromID(rv)
}

// The time of the most recent event.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/lastEventTimestamp
func (o GCDevicePhysicalInputStateObject) LastEventTimestamp() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("lastEventTimestamp"))
	return rv
}

// The time in seconds between the last event and the current time.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/lastEventLatency
func (o GCDevicePhysicalInputStateObject) LastEventLatency() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("lastEventLatency"))
	return rv
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
