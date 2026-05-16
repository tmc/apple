// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Physical extents scale the normalized value reported by [GCLinearInput] into physical units.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputExtents
type GCPhysicalInputExtents interface {
	objectivec.IObject

	// The maximum value for the physical extent of the input.
	//
	// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputExtents/maximumValue
	MaximumValue() float64

	// The minimum value for the physical extent of the input.
	//
	// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputExtents/minimumValue
	MinimumValue() float64

	// The value of the input, scaled into physical units.
	//
	// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputExtents/scaledValue
	ScaledValue() float64
}

// GCPhysicalInputExtentsObject wraps an existing Objective-C object that conforms to the GCPhysicalInputExtents protocol.
type GCPhysicalInputExtentsObject struct {
	objectivec.Object
}

func (o GCPhysicalInputExtentsObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCPhysicalInputExtentsObjectFromID constructs a [GCPhysicalInputExtentsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCPhysicalInputExtentsObjectFromID(id objc.ID) GCPhysicalInputExtentsObject {
	return GCPhysicalInputExtentsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The maximum value for the physical extent of the input.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputExtents/maximumValue
func (o GCPhysicalInputExtentsObject) MaximumValue() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("maximumValue"))
	return rv
}

// The minimum value for the physical extent of the input.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputExtents/minimumValue
func (o GCPhysicalInputExtentsObject) MinimumValue() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("minimumValue"))
	return rv
}

// The value of the input, scaled into physical units.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputExtents/scaledValue
func (o GCPhysicalInputExtentsObject) ScaledValue() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("scaledValue"))
	return rv
}
