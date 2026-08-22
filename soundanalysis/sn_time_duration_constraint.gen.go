// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SNTimeDurationConstraint] class.
var (
	_SNTimeDurationConstraintClass     SNTimeDurationConstraintClass
	_SNTimeDurationConstraintClassOnce sync.Once
)

func getSNTimeDurationConstraintClass() SNTimeDurationConstraintClass {
	_SNTimeDurationConstraintClassOnce.Do(func() {
		_SNTimeDurationConstraintClass = SNTimeDurationConstraintClass{class: objc.GetClass("SNTimeDurationConstraint")}
	})
	return _SNTimeDurationConstraintClass
}

// GetSNTimeDurationConstraintClass returns the class object for SNTimeDurationConstraint.
func GetSNTimeDurationConstraintClass() SNTimeDurationConstraintClass {
	return getSNTimeDurationConstraintClass()
}

type SNTimeDurationConstraintClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SNTimeDurationConstraintClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SNTimeDurationConstraintClass) Alloc() SNTimeDurationConstraint {
	rv := objc.Send[SNTimeDurationConstraint](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// Defines the time duration windows the request’s underlying sound
// classifier accepts with a range, or an array, of durations.
//
// # Overview
//
// Inspect the constraint’s [SNTimeDurationConstraint.Type] property first
// to determine whether to check [SNTimeDurationConstraint.DurationRange] or
// [SNTimeDurationConstraint.EnumeratedDurations] next.
//
// # Inspecting a Constraint
//
//   - [SNTimeDurationConstraint.Type]: An enumeration that tells you which constraint property to inspect.
//   - [SNTimeDurationConstraint.DurationRange]: A time duration range the request’s underlying sound classifier accepts.
//   - [SNTimeDurationConstraint.EnumeratedDurations]: An array of time durations the request’s underlying sound classifier accepts.
//
// # Creating a Time Duration Constraint
//
//   - [SNTimeDurationConstraint.InitWithDurationRange]: Creates a constraint with a time duration range.
//   - [SNTimeDurationConstraint.InitWithEnumeratedDurations]: Creates a constraint with discrete time durations.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNTimeDurationConstraint-c.class
type SNTimeDurationConstraint struct {
	objectivec.Object
}

// SNTimeDurationConstraintFromID constructs a [SNTimeDurationConstraint] from an objc.ID.
//
// Defines the time duration windows the request’s underlying sound
// classifier accepts with a range, or an array, of durations.
func SNTimeDurationConstraintFromID(id objc.ID) SNTimeDurationConstraint {
	return SNTimeDurationConstraint{objectivec.Object{ID: id}}
}

// Ensure SNTimeDurationConstraint implements ISNTimeDurationConstraint.
var _ ISNTimeDurationConstraint = SNTimeDurationConstraint{}

// An interface definition for the [SNTimeDurationConstraint] class.
//
// # Inspecting a Constraint
//
//   - [ISNTimeDurationConstraint.Type]: An enumeration that tells you which constraint property to inspect.
//   - [ISNTimeDurationConstraint.DurationRange]: A time duration range the request’s underlying sound classifier accepts.
//   - [ISNTimeDurationConstraint.EnumeratedDurations]: An array of time durations the request’s underlying sound classifier accepts.
//
// # Creating a Time Duration Constraint
//
//   - [ISNTimeDurationConstraint.InitWithDurationRange]: Creates a constraint with a time duration range.
//   - [ISNTimeDurationConstraint.InitWithEnumeratedDurations]: Creates a constraint with discrete time durations.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNTimeDurationConstraint-c.class
type ISNTimeDurationConstraint interface {
	objectivec.IObject

	// Topic: Inspecting a Constraint

	// An enumeration that tells you which constraint property to inspect.
	Type() SNTimeDurationConstraintType
	// A time duration range the request’s underlying sound classifier accepts.
	DurationRange() coremedia.CMTimeRange
	// An array of time durations the request’s underlying sound classifier accepts.
	EnumeratedDurations() []foundation.NSValue

	// Topic: Creating a Time Duration Constraint

	// Creates a constraint with a time duration range.
	InitWithDurationRange(durationRange coremedia.CMTimeRange) SNTimeDurationConstraint
	// Creates a constraint with discrete time durations.
	InitWithEnumeratedDurations(enumeratedDurations []foundation.NSValue) SNTimeDurationConstraint
}

// Init initializes the instance.
func (t SNTimeDurationConstraint) Init() SNTimeDurationConstraint {
	rv := objc.Send[SNTimeDurationConstraint](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t SNTimeDurationConstraint) Autorelease() SNTimeDurationConstraint {
	rv := objc.Send[SNTimeDurationConstraint](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewSNTimeDurationConstraint creates a new SNTimeDurationConstraint instance.
func NewSNTimeDurationConstraint() SNTimeDurationConstraint {
	class := getSNTimeDurationConstraintClass()
	rv := objc.Send[SNTimeDurationConstraint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a constraint with a time duration range.
//
// durationRange: A range of time durations.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNTimeDurationConstraint-c.class/initWithDurationRange:
func NewTimeDurationConstraintWithDurationRange(durationRange coremedia.CMTimeRange) SNTimeDurationConstraint {
	instance := getSNTimeDurationConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDurationRange:"), durationRange)
	return SNTimeDurationConstraintFromID(rv)
}

// Creates a constraint with discrete time durations.
//
// enumeratedDurations: An array of time durations.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNTimeDurationConstraint-c.class/initWithEnumeratedDurations:
func NewTimeDurationConstraintWithEnumeratedDurations(enumeratedDurations []foundation.NSValue) SNTimeDurationConstraint {
	instance := getSNTimeDurationConstraintClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEnumeratedDurations:"), objectivec.IObjectSliceToNSArray(enumeratedDurations))
	return SNTimeDurationConstraintFromID(rv)
}

// Creates a constraint with a time duration range.
//
// durationRange: A range of time durations.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNTimeDurationConstraint-c.class/initWithDurationRange:
func (t SNTimeDurationConstraint) InitWithDurationRange(durationRange coremedia.CMTimeRange) SNTimeDurationConstraint {
	rv := objc.Send[SNTimeDurationConstraint](t.ID, objc.Sel("initWithDurationRange:"), durationRange)
	return rv
}

// Creates a constraint with discrete time durations.
//
// enumeratedDurations: An array of time durations.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNTimeDurationConstraint-c.class/initWithEnumeratedDurations:
func (t SNTimeDurationConstraint) InitWithEnumeratedDurations(enumeratedDurations []foundation.NSValue) SNTimeDurationConstraint {
	rv := objc.Send[SNTimeDurationConstraint](t.ID, objc.Sel("initWithEnumeratedDurations:"), objectivec.IObjectSliceToNSArray(enumeratedDurations))
	return rv
}

// An enumeration that tells you which constraint property to inspect.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNTimeDurationConstraint-c.class/type
func (t SNTimeDurationConstraint) Type() SNTimeDurationConstraintType {
	rv := objc.Send[SNTimeDurationConstraintType](t.ID, objc.Sel("type"))
	return SNTimeDurationConstraintType(rv)
}

// A time duration range the request’s underlying sound classifier accepts.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNTimeDurationConstraint-c.class/durationRange
func (t SNTimeDurationConstraint) DurationRange() coremedia.CMTimeRange {
	rv := objc.Send[coremedia.CMTimeRange](t.ID, objc.Sel("durationRange"))
	return coremedia.CMTimeRange(rv)
}

// An array of time durations the request’s underlying sound classifier
// accepts.
//
// See: https://developer.apple.com/documentation/SoundAnalysis/SNTimeDurationConstraint-c.class/enumeratedDurations
func (t SNTimeDurationConstraint) EnumeratedDurations() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("enumeratedDurations"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}
