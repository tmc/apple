// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The common properties of inputs that switch between two or more positions.
//
// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput
type GCSwitchPositionInput interface {
	objectivec.IObject

	// The range of possible values for the switch.
	//
	// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/positionRange
	PositionRange() foundation.NSRange

	// A Boolean value that indicates whether the position change is sequential.
	//
	// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/isSequential
	IsSequential() bool

	// A Boolean value that indicates whether the position value wraps when it reaches the range’s minimum or maximum value.
	//
	// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/canWrap
	CanWrap() bool

	// The position of the switch.
	//
	// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/position
	Position() int

	// The block that the profile calls when the value of the switch changes.
	//
	// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/positionDidChangeHandler
	PositionDidChangeHandler() func(objc.ID, int64)

	// A timestamp for when the profile reports the last position.
	//
	// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/lastPositionTimestamp
	LastPositionTimestamp() float64

	// The time in seconds between the current and previous positions.
	//
	// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/lastPositionLatency
	LastPositionLatency() float64

	// One or more physical actions the user performs to manipulate the input.
	//
	// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/sources
	Sources() foundation.INSSet
}

// GCSwitchPositionInputObject wraps an existing Objective-C object that conforms to the GCSwitchPositionInput protocol.
type GCSwitchPositionInputObject struct {
	objectivec.Object
}

func (o GCSwitchPositionInputObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCSwitchPositionInputObjectFromID constructs a [GCSwitchPositionInputObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCSwitchPositionInputObjectFromID(id objc.ID) GCSwitchPositionInputObject {
	return GCSwitchPositionInputObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The range of possible values for the switch.
//
// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/positionRange
func (o GCSwitchPositionInputObject) PositionRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("positionRange"))
	return rv
}

// A Boolean value that indicates whether the position change is sequential.
//
// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/isSequential
func (o GCSwitchPositionInputObject) IsSequential() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isSequential"))
	return rv
}

// A Boolean value that indicates whether the position value wraps when it
// reaches the range’s minimum or maximum value.
//
// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/canWrap
func (o GCSwitchPositionInputObject) CanWrap() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("canWrap"))
	return rv
}

// The position of the switch.
//
// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/position
func (o GCSwitchPositionInputObject) Position() int {
	rv := objc.Send[int](o.ID, objc.Sel("position"))
	return rv
}

// The block that the profile calls when the value of the switch changes.
//
// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/positionDidChangeHandler
func (o GCSwitchPositionInputObject) PositionDidChangeHandler() func(objc.ID, int64) {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("positionDidChangeHandler"))
	// Block/function return - cannot convert from objc.ID to Go func
	_ = rv
	return nil
}

// A timestamp for when the profile reports the last position.
//
// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/lastPositionTimestamp
func (o GCSwitchPositionInputObject) LastPositionTimestamp() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("lastPositionTimestamp"))
	return rv
}

// The time in seconds between the current and previous positions.
//
// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/lastPositionLatency
func (o GCSwitchPositionInputObject) LastPositionLatency() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("lastPositionLatency"))
	return rv
}

// One or more physical actions the user performs to manipulate the input.
//
// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/sources
func (o GCSwitchPositionInputObject) Sources() foundation.INSSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sources"))
	return foundation.NSSetFromID(rv)
}
