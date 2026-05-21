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

	// A Boolean value that indicates whether the position change is sequential.
	//
	// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/isSequential
	IsSequential() bool

	// The range of possible values for the switch.
	//
	// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/positionRange
	PositionRange() foundation.NSRange

	// A Boolean value that indicates whether the position change is sequential.
	//
	// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/isSequential
	Sequential() bool

	// A Boolean value that indicates whether the position value wraps when it reaches the range’s minimum or maximum value.
	//
	// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/canWrap
	CanWrap() bool

	// The position of the switch.
	//
	// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/position
	Position() int

	// A timestamp for when the profile reports the last position.
	//
	// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/lastPositionTimestamp
	LastPositionTimestamp() foundation.NSTimeInterval

	// The time in seconds between the current and previous positions.
	//
	// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/lastPositionLatency
	LastPositionLatency() foundation.NSTimeInterval

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

// A Boolean value that indicates whether the position change is sequential.
//
// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/isSequential
func (o GCSwitchPositionInputObject) IsSequential() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isSequential"))
	return rv
}

// The range of possible values for the switch.
//
// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/positionRange
func (o GCSwitchPositionInputObject) PositionRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("positionRange"))
	return foundation.NSRange(rv)
}

// A Boolean value that indicates whether the position change is sequential.
//
// # Discussion
//
// A sequential gear shift requires the user to move through the gears in
// sequence.
//
// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/isSequential
func (o GCSwitchPositionInputObject) Sequential() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isSequential"))
	return bool(rv)
}

// A Boolean value that indicates whether the position value wraps when it
// reaches the range’s minimum or maximum value.
//
// # Discussion
//
// For non-sequential switches, this property is always true.
//
// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/canWrap
func (o GCSwitchPositionInputObject) CanWrap() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("canWrap"))
	return bool(rv)
}

// The position of the switch.
//
// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/position
func (o GCSwitchPositionInputObject) Position() int {
	rv := objc.Send[int](o.ID, objc.Sel("position"))
	return int(rv)
}

// A timestamp for when the profile reports the last position.
//
// # Discussion
//
// This property isn’t a specific date and time. To determine the time
// between positions, subtract a previous value from the current value.
//
// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/lastPositionTimestamp
func (o GCSwitchPositionInputObject) LastPositionTimestamp() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](o.ID, objc.Sel("lastPositionTimestamp"))
	return foundation.NSTimeInterval(rv)
}

// The time in seconds between the current and previous positions.
//
// # Discussion
//
// Use this property as a minimum latency value that may not include latency
// that accrues on the device or when it transmits the event.
//
// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/lastPositionLatency
func (o GCSwitchPositionInputObject) LastPositionLatency() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](o.ID, objc.Sel("lastPositionLatency"))
	return foundation.NSTimeInterval(rv)
}

// One or more physical actions the user performs to manipulate the input.
//
// See: https://developer.apple.com/documentation/GameController/GCSwitchPositionInput/sources
func (o GCSwitchPositionInputObject) Sources() foundation.INSSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sources"))
	return foundation.NSSetFromID(rv)
}
