// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The common properties of inputs that provide values in unit coordinates.
//
// See: https://developer.apple.com/documentation/GameController/GCLinearInput
type GCLinearInput interface {
	objectivec.IObject

	// A Boolean value that indicates whether the input value wraps when it reaches the range’s minimum or maximum value.
	//
	// See: https://developer.apple.com/documentation/GameController/GCLinearInput/canWrap
	CanWrap() bool

	// A Boolean value that indicates whether the input provides analog values.
	//
	// See: https://developer.apple.com/documentation/GameController/GCLinearInput/isAnalog
	IsAnalog() bool

	// The value in unit coordinates.
	//
	// See: https://developer.apple.com/documentation/GameController/GCLinearInput/value
	Value() float32

	// The block that the profile calls when an element’s value changes.
	//
	// See: https://developer.apple.com/documentation/GameController/GCLinearInput/valueDidChangeHandler
	ValueDidChangeHandler() func(objc.ID, float32)

	// The time of the most recent value change.
	//
	// See: https://developer.apple.com/documentation/GameController/GCLinearInput/lastValueTimestamp
	LastValueTimestamp() float64

	// The time in seconds between the last value change and the current time.
	//
	// See: https://developer.apple.com/documentation/GameController/GCLinearInput/lastValueLatency
	LastValueLatency() float64

	// One or more physical actions the user performs to manipulate the input.
	//
	// See: https://developer.apple.com/documentation/GameController/GCLinearInput/sources
	Sources() foundation.INSSet

	// An object describing the physical extents of the input, if the input represents a physical unit of measurement.
	//
	// See: https://developer.apple.com/documentation/GameController/GCLinearInput/physicalExtents
	PhysicalExtents() GCPhysicalInputExtents
}

// GCLinearInputObject wraps an existing Objective-C object that conforms to the GCLinearInput protocol.
type GCLinearInputObject struct {
	objectivec.Object
}

func (o GCLinearInputObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCLinearInputObjectFromID constructs a [GCLinearInputObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCLinearInputObjectFromID(id objc.ID) GCLinearInputObject {
	return GCLinearInputObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A Boolean value that indicates whether the input value wraps when it
// reaches the range’s minimum or maximum value.
//
// See: https://developer.apple.com/documentation/GameController/GCLinearInput/canWrap
func (o GCLinearInputObject) CanWrap() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("canWrap"))
	return rv
}

// A Boolean value that indicates whether the input provides analog values.
//
// See: https://developer.apple.com/documentation/GameController/GCLinearInput/isAnalog
func (o GCLinearInputObject) IsAnalog() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAnalog"))
	return rv
}

// The value in unit coordinates.
//
// See: https://developer.apple.com/documentation/GameController/GCLinearInput/value
func (o GCLinearInputObject) Value() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("value"))
	return rv
}

// The block that the profile calls when an element’s value changes.
//
// See: https://developer.apple.com/documentation/GameController/GCLinearInput/valueDidChangeHandler
func (o GCLinearInputObject) ValueDidChangeHandler() func(objc.ID, float32) {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("valueDidChangeHandler"))
	// Block/function return - cannot convert from objc.ID to Go func
	_ = rv
	return nil
}

// The time of the most recent value change.
//
// See: https://developer.apple.com/documentation/GameController/GCLinearInput/lastValueTimestamp
func (o GCLinearInputObject) LastValueTimestamp() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("lastValueTimestamp"))
	return rv
}

// The time in seconds between the last value change and the current time.
//
// See: https://developer.apple.com/documentation/GameController/GCLinearInput/lastValueLatency
func (o GCLinearInputObject) LastValueLatency() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("lastValueLatency"))
	return rv
}

// One or more physical actions the user performs to manipulate the input.
//
// See: https://developer.apple.com/documentation/GameController/GCLinearInput/sources
func (o GCLinearInputObject) Sources() foundation.INSSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sources"))
	return foundation.NSSetFromID(rv)
}

// An object describing the physical extents of the input, if the input
// represents a physical unit of measurement.
//
// See: https://developer.apple.com/documentation/GameController/GCLinearInput/physicalExtents
func (o GCLinearInputObject) PhysicalExtents() GCPhysicalInputExtents {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("physicalExtents"))
	return GCPhysicalInputExtentsObjectFromID(rv)
}
