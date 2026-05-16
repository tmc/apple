// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The common properties of inputs that provide absolute values along an axis with a fixed origin.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisInput
type GCAxisInput interface {
	objectivec.IObject

	// A Boolean value that indicates whether the value wraps when it reaches the range’s minimum or maximum value.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxisInput/canWrap
	CanWrap() bool

	// A Boolean value that indicates whether the input provides analog values.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxisInput/isAnalog
	IsAnalog() bool

	// The value along the axis, in unit coordinates.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxisInput/value
	Value() float32

	// The block that the input object calls when the value changes.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxisInput/valueDidChangeHandler
	ValueDidChangeHandler() func(objc.ID, float32)

	// The time of the most recent value change.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxisInput/lastValueTimestamp
	LastValueTimestamp() float64

	// The time in seconds between the last value change and the current time.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxisInput/lastValueLatency
	LastValueLatency() float64

	// One or more physical actions the user performs to manipulate the input.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxisInput/sources
	Sources() foundation.INSSet
}

// GCAxisInputObject wraps an existing Objective-C object that conforms to the GCAxisInput protocol.
type GCAxisInputObject struct {
	objectivec.Object
}

func (o GCAxisInputObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCAxisInputObjectFromID constructs a [GCAxisInputObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCAxisInputObjectFromID(id objc.ID) GCAxisInputObject {
	return GCAxisInputObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A Boolean value that indicates whether the value wraps when it reaches the
// range’s minimum or maximum value.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisInput/canWrap
func (o GCAxisInputObject) CanWrap() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("canWrap"))
	return rv
}

// A Boolean value that indicates whether the input provides analog values.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisInput/isAnalog
func (o GCAxisInputObject) IsAnalog() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAnalog"))
	return rv
}

// The value along the axis, in unit coordinates.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisInput/value
func (o GCAxisInputObject) Value() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("value"))
	return rv
}

// The block that the input object calls when the value changes.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisInput/valueDidChangeHandler
func (o GCAxisInputObject) ValueDidChangeHandler() func(objc.ID, float32) {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("valueDidChangeHandler"))
	// Block/function return - cannot convert from objc.ID to Go func
	_ = rv
	return nil
}

// The time of the most recent value change.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisInput/lastValueTimestamp
func (o GCAxisInputObject) LastValueTimestamp() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("lastValueTimestamp"))
	return rv
}

// The time in seconds between the last value change and the current time.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisInput/lastValueLatency
func (o GCAxisInputObject) LastValueLatency() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("lastValueLatency"))
	return rv
}

// One or more physical actions the user performs to manipulate the input.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisInput/sources
func (o GCAxisInputObject) Sources() foundation.INSSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sources"))
	return foundation.NSSetFromID(rv)
}
