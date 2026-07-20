// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
//   - [SOPowerSavvyTimer.Timer]
//   - [SOPowerSavvyTimer.SetTimer]
type SOPowerSavvyTimer struct {
	foundation.NSTimer
}

// SOPowerSavvyTimerFromID constructs a [SOPowerSavvyTimer] from an objc.ID.
func SOPowerSavvyTimerFromID(id objc.ID) SOPowerSavvyTimer {
	return SOPowerSavvyTimer{NSTimer: foundation.NSTimerFromID(id)}
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
//   - [ISOPowerSavvyTimer.Timer]
//   - [ISOPowerSavvyTimer.SetTimer]
type ISOPowerSavvyTimer interface {
	foundation.INSTimer

	// Topic: Methods

	_target(_target objectivec.IObject)
	IsValid() bool
	Repeats() bool
	SetRepeats(value bool)
	Selector() objectivec.SEL
	SetSelector(value objectivec.SEL)
	Target() objectivec.IObject
	SetTarget(value objectivec.IObject)
	Timer() foundation.Timer
	SetTimer(value foundation.Timer)
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

func (s SOPowerSavvyTimer) _target(_target objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("_target:"), _target)
}
func (s SOPowerSavvyTimer) IsValid() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isValid"))
	return rv
}

func (_SOPowerSavvyTimerClass SOPowerSavvyTimerClass) RequestTargetPerformSelectorWithObjectAfterDelay(target objectivec.IObject, selector objc.SEL, object objectivec.IObject, delay float64) {
	objc.Send[objc.ID](objc.ID(_SOPowerSavvyTimerClass.class), objc.Sel("requestTarget:performSelector:withObject:afterDelay:"), target, selector, object, delay)
}

func (s SOPowerSavvyTimer) Repeats() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("repeats"))
	return rv
}
func (s SOPowerSavvyTimer) SetRepeats(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setRepeats:"), value)
}
func (s SOPowerSavvyTimer) Selector() objectivec.SEL {
	rv := objc.Send[objc.SEL](s.ID, objc.Sel("selector"))
	return objectivec.SEL(rv)
}
func (s SOPowerSavvyTimer) SetSelector(value objectivec.SEL) {
	objc.Send[struct{}](s.ID, objc.Sel("setSelector:"), value)
}
func (s SOPowerSavvyTimer) Target() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("target"))
	return objectivec.Object{ID: rv}
}
func (s SOPowerSavvyTimer) SetTarget(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setTarget:"), value)
}
func (s SOPowerSavvyTimer) Timer() foundation.Timer {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("timer"))
	return foundation.TimerFromID(objc.ID(rv))
}
func (s SOPowerSavvyTimer) SetTimer(value foundation.Timer) {
	objc.Send[struct{}](s.ID, objc.Sel("setTimer:"), value)
}
