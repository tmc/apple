// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// Methods that model a hierarchical timing system, allowing objects to map time between their parent and local time.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming
type CAMediaTiming interface {
	objectivec.IObject

	// Specifies the begin time of the receiver in relation to its parent object, if applicable.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/beginTime
	BeginTime() float64

	// Specifies an additional time offset in active local time.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/timeOffset
	TimeOffset() float64

	// Determines the number of times the animation will repeat.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/repeatCount
	RepeatCount() float32

	// Determines how many seconds the animation will repeat for.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/repeatDuration
	RepeatDuration() float64

	// Specifies the basic duration of the animation, in seconds.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/duration
	Duration() float64

	// Specifies how time is mapped to receiver’s time space from the parent time space.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/speed
	Speed() float32

	// Determines if the receiver plays in the reverse upon completion.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/autoreverses
	Autoreverses() bool

	// Determines if the receiver’s presentation is frozen or removed once its active duration has completed.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/fillMode
	FillMode() CAMediaTimingFillMode

	// Specifies the begin time of the receiver in relation to its parent object, if applicable.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/beginTime
	SetBeginTime(value float64)

	// Specifies an additional time offset in active local time.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/timeOffset
	SetTimeOffset(value float64)

	// Determines the number of times the animation will repeat.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/repeatCount
	SetRepeatCount(value float32)

	// Determines how many seconds the animation will repeat for.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/repeatDuration
	SetRepeatDuration(value float64)

	// Specifies the basic duration of the animation, in seconds.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/duration
	SetDuration(value float64)

	// Specifies how time is mapped to receiver’s time space from the parent time space.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/speed
	SetSpeed(value float32)

	// Determines if the receiver plays in the reverse upon completion.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/autoreverses
	SetAutoreverses(value bool)

	// Determines if the receiver’s presentation is frozen or removed once its active duration has completed.
	//
	// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/fillMode
	SetFillMode(value CAMediaTimingFillMode)
}

// CAMediaTimingObject wraps an existing Objective-C object that conforms to the CAMediaTiming protocol.
type CAMediaTimingObject struct {
	objectivec.Object
}
func (o CAMediaTimingObject) BaseObject() objectivec.Object {
	return o.Object
}

// CAMediaTimingObjectFromID constructs a [CAMediaTimingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CAMediaTimingObjectFromID(id objc.ID) CAMediaTimingObject {
	return CAMediaTimingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Specifies the begin time of the receiver in relation to its parent object,
// if applicable.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/beginTime
func (o CAMediaTimingObject) BeginTime() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("beginTime"))
	return rv
	}
// Specifies an additional time offset in active local time.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/timeOffset
func (o CAMediaTimingObject) TimeOffset() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("timeOffset"))
	return rv
	}
// Determines the number of times the animation will repeat.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/repeatCount
func (o CAMediaTimingObject) RepeatCount() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("repeatCount"))
	return rv
	}
// Determines how many seconds the animation will repeat for.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/repeatDuration
func (o CAMediaTimingObject) RepeatDuration() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("repeatDuration"))
	return rv
	}
// Specifies the basic duration of the animation, in seconds.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/duration
func (o CAMediaTimingObject) Duration() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("duration"))
	return rv
	}
// Specifies how time is mapped to receiver’s time space from the parent
// time space.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/speed
func (o CAMediaTimingObject) Speed() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("speed"))
	return rv
	}
// Determines if the receiver plays in the reverse upon completion.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/autoreverses
func (o CAMediaTimingObject) Autoreverses() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("autoreverses"))
	return rv
	}
// Determines if the receiver’s presentation is frozen or removed once its
// active duration has completed.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/fillMode
func (o CAMediaTimingObject) FillMode() CAMediaTimingFillMode {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("fillMode"))
	return CAMediaTimingFillMode(foundation.NSStringFromID(rv).String())
	}

func (o CAMediaTimingObject) SetBeginTime(value float64) {
	objc.Send[struct{}](o.ID, objc.Sel("setBeginTime:"), value)
}

func (o CAMediaTimingObject) SetTimeOffset(value float64) {
	objc.Send[struct{}](o.ID, objc.Sel("setTimeOffset:"), value)
}

func (o CAMediaTimingObject) SetRepeatCount(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRepeatCount:"), value)
}

func (o CAMediaTimingObject) SetRepeatDuration(value float64) {
	objc.Send[struct{}](o.ID, objc.Sel("setRepeatDuration:"), value)
}

func (o CAMediaTimingObject) SetDuration(value float64) {
	objc.Send[struct{}](o.ID, objc.Sel("setDuration:"), value)
}

func (o CAMediaTimingObject) SetSpeed(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSpeed:"), value)
}

func (o CAMediaTimingObject) SetAutoreverses(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAutoreverses:"), value)
}

func (o CAMediaTimingObject) SetFillMode(value CAMediaTimingFillMode) {
	objc.Send[struct{}](o.ID, objc.Sel("setFillMode:"), objc.String(string(value)))
}

