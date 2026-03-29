// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SOPowerSavvyTimer] class.
var (
	_SOPowerSavvyTimerClass     SOPowerSavvyTimerClass
	_SOPowerSavvyTimerClassOnce sync.Once
)

func getSOPowerSavvyTimerClass() SOPowerSavvyTimerClass {
	_SOPowerSavvyTimerClassOnce.Do(func() {
		_SOPowerSavvyTimerClass = SOPowerSavvyTimerClass{class: objc.GetClass("SOPowerSavvyTimer")}
	})
	return _SOPowerSavvyTimerClass
}

// GetSOPowerSavvyTimerClass returns the class object for SOPowerSavvyTimer.
func GetSOPowerSavvyTimerClass() SOPowerSavvyTimerClass {
	return getSOPowerSavvyTimerClass()
}

type SOPowerSavvyTimerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SOPowerSavvyTimerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SOPowerSavvyTimerClass) Alloc() SOPowerSavvyTimer {
	rv := objc.Send[SOPowerSavvyTimer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [SOPowerSavvyTimer._target]
//   - [SOPowerSavvyTimer.IsValid]
//   - [SOPowerSavvyTimer.Repeats]
//   - [SOPowerSavvyTimer.SetRepeats]
//   - [SOPowerSavvyTimer.Selector]
//   - [SOPowerSavvyTimer.SetSelector]
//   - [SOPowerSavvyTimer.Target]
//   - [SOPowerSavvyTimer.SetTarget]
//   - [SOPowerSavvyTimer.GetTimer]
//   - [SOPowerSavvyTimer.SetGetTimer]
// See: https://developer.apple.com/documentation/SpeechObjects/SOPowerSavvyTimer
type SOPowerSavvyTimer struct {
	foundation.Timer
}

// SOPowerSavvyTimerFromID constructs a [SOPowerSavvyTimer] from an objc.ID.
func SOPowerSavvyTimerFromID(id objc.ID) SOPowerSavvyTimer {
	return SOPowerSavvyTimer{Timer: foundation.TimerFromID(id)}
}
// Ensure SOPowerSavvyTimer implements ISOPowerSavvyTimer.
var _ ISOPowerSavvyTimer = SOPowerSavvyTimer{}

// An interface definition for the [SOPowerSavvyTimer] class.
//
// # Methods
//
//   - [ISOPowerSavvyTimer._target]
//   - [ISOPowerSavvyTimer.IsValid]
//   - [ISOPowerSavvyTimer.Repeats]
//   - [ISOPowerSavvyTimer.SetRepeats]
//   - [ISOPowerSavvyTimer.Selector]
//   - [ISOPowerSavvyTimer.SetSelector]
//   - [ISOPowerSavvyTimer.Target]
//   - [ISOPowerSavvyTimer.SetTarget]
//   - [ISOPowerSavvyTimer.GetTimer]
//   - [ISOPowerSavvyTimer.SetGetTimer]
//
// See: https://developer.apple.com/documentation/SpeechObjects/SOPowerSavvyTimer
type ISOPowerSavvyTimer interface {
	foundation.ITimer

	// Topic: Methods

	_target(_target objectivec.IObject)
	IsValid() bool
	Repeats() bool
	SetRepeats(value bool)
	Selector() objc.SEL
	SetSelector(value objc.SEL)
	Target() objectivec.IObject
	SetTarget(value objectivec.IObject)
	GetTimer() *foundation.NSTimer
	SetGetTimer(value *foundation.NSTimer)
}

// Init initializes the instance.
func (s SOPowerSavvyTimer) Init() SOPowerSavvyTimer {
	rv := objc.Send[SOPowerSavvyTimer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOPowerSavvyTimer) Autorelease() SOPowerSavvyTimer {
	rv := objc.Send[SOPowerSavvyTimer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOPowerSavvyTimer creates a new SOPowerSavvyTimer instance.
func NewSOPowerSavvyTimer() SOPowerSavvyTimer {
	class := getSOPowerSavvyTimerClass()
	rv := objc.Send[SOPowerSavvyTimer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/SpeechObjects/SOPowerSavvyTimer/_target:
func (s SOPowerSavvyTimer) _target(_target objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("_target:"), _target)
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOPowerSavvyTimer/isValid
func (s SOPowerSavvyTimer) IsValid() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isValid"))
	return rv
}

//
// See: https://developer.apple.com/documentation/SpeechObjects/SOPowerSavvyTimer/requestTarget:performSelector:withObject:afterDelay:
func (_SOPowerSavvyTimerClass SOPowerSavvyTimerClass) RequestTargetPerformSelectorWithObjectAfterDelay(target objectivec.IObject, selector objc.SEL, object objectivec.IObject, delay float64) {
	objc.Send[objc.ID](objc.ID(_SOPowerSavvyTimerClass.class), objc.Sel("requestTarget:performSelector:withObject:afterDelay:"), target, selector, object, delay)
}

// See: https://developer.apple.com/documentation/SpeechObjects/SOPowerSavvyTimer/repeats
func (s SOPowerSavvyTimer) Repeats() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("repeats"))
	return rv
}
func (s SOPowerSavvyTimer) SetRepeats(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setRepeats:"), value)
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOPowerSavvyTimer/selector
func (s SOPowerSavvyTimer) Selector() objc.SEL {
	rv := objc.Send[objc.SEL](s.ID, objc.Sel("selector"))
	return rv
}
func (s SOPowerSavvyTimer) SetSelector(value objc.SEL) {
	objc.Send[struct{}](s.ID, objc.Sel("setSelector:"), value)
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOPowerSavvyTimer/target
func (s SOPowerSavvyTimer) Target() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("target"))
	return objectivec.Object{ID: rv}
}
func (s SOPowerSavvyTimer) SetTarget(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setTarget:"), value)
}
// See: https://developer.apple.com/documentation/SpeechObjects/SOPowerSavvyTimer/timer
func (s SOPowerSavvyTimer) GetTimer() *foundation.NSTimer {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("timer"))
	if rv == 0 {
		return nil
	}
	val := foundation.NSTimerFromID(objc.ID(rv))
	return &val
}
func (s SOPowerSavvyTimer) SetGetTimer(value *foundation.NSTimer) {
	if value == nil {
		objc.Send[struct{}](s.ID, objc.Sel("setTimer:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](s.ID, objc.Sel("setTimer:"), value)
}

