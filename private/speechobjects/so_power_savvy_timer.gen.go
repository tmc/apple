// Code generated from Apple documentation for speechobjects. DO NOT EDIT.

package speechobjects

import (
	"sync"
	"unsafe"

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
	rv := objc.SendIfResponds[SOPowerSavvyTimer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SOPowerSavvyTimer._target]
//   - [SOPowerSavvyTimer.Fire]
//   - [SOPowerSavvyTimer.FireDate]
//   - [SOPowerSavvyTimer.Invalidate]
//   - [SOPowerSavvyTimer.IsValid]
//   - [SOPowerSavvyTimer.Repeats]
//   - [SOPowerSavvyTimer.SetRepeats]
//   - [SOPowerSavvyTimer.Selector]
//   - [SOPowerSavvyTimer.SetSelector]
//   - [SOPowerSavvyTimer.SetFireDate]
//   - [SOPowerSavvyTimer.Target]
//   - [SOPowerSavvyTimer.SetTarget]
//   - [SOPowerSavvyTimer.TimeInterval]
//   - [SOPowerSavvyTimer.Timer]
//   - [SOPowerSavvyTimer.SetTimer]
//   - [SOPowerSavvyTimer.UserInfo]
type SOPowerSavvyTimer struct {
	objectivec.Object
}

// SOPowerSavvyTimerFromID constructs a [SOPowerSavvyTimer] from an objc.ID.
func SOPowerSavvyTimerFromID(id objc.ID) SOPowerSavvyTimer {
	return SOPowerSavvyTimer{objectivec.Object{ID: id}}
}

// NOTE: SOPowerSavvyTimer embeds objectivec.Object because the parent type is
// unavailable, but ISOPowerSavvyTimer embeds INSTimer, which that fallback
// cannot satisfy; skip compile-time assertion.

// An interface definition for the [SOPowerSavvyTimer] class.
//
// # Methods
//
//   - [ISOPowerSavvyTimer._target]
//   - [ISOPowerSavvyTimer.Fire]
//   - [ISOPowerSavvyTimer.FireDate]
//   - [ISOPowerSavvyTimer.Invalidate]
//   - [ISOPowerSavvyTimer.IsValid]
//   - [ISOPowerSavvyTimer.Repeats]
//   - [ISOPowerSavvyTimer.SetRepeats]
//   - [ISOPowerSavvyTimer.Selector]
//   - [ISOPowerSavvyTimer.SetSelector]
//   - [ISOPowerSavvyTimer.SetFireDate]
//   - [ISOPowerSavvyTimer.Target]
//   - [ISOPowerSavvyTimer.SetTarget]
//   - [ISOPowerSavvyTimer.TimeInterval]
//   - [ISOPowerSavvyTimer.Timer]
//   - [ISOPowerSavvyTimer.SetTimer]
//   - [ISOPowerSavvyTimer.UserInfo]
type ISOPowerSavvyTimer interface {
	INSTimer

	// Topic: Methods

	_target(_target objectivec.IObject)
	Fire()
	FireDate() objectivec.IObject
	Invalidate()
	IsValid() bool
	Repeats() bool
	SetRepeats(value bool)
	Selector() objectivec.SEL
	SetSelector(value objectivec.SEL)
	SetFireDate(date objectivec.IObject)
	Target() objectivec.IObject
	SetTarget(value objectivec.IObject)
	TimeInterval() float64
	Timer() unsafe.Pointer
	SetTimer(value unsafe.Pointer)
	UserInfo() objectivec.IObject
}

// Init initializes the instance.
func (s SOPowerSavvyTimer) Init() SOPowerSavvyTimer {
	rv := objc.SendIfResponds[SOPowerSavvyTimer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SOPowerSavvyTimer) Autorelease() SOPowerSavvyTimer {
	rv := objc.SendIfResponds[SOPowerSavvyTimer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSOPowerSavvyTimer creates a new SOPowerSavvyTimer instance.
func NewSOPowerSavvyTimer() SOPowerSavvyTimer {
	class := getSOPowerSavvyTimerClass()
	rv := objc.SendIfResponds[SOPowerSavvyTimer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (s SOPowerSavvyTimer) _target(_target objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("_target:"), _target)
}
func (s SOPowerSavvyTimer) Fire() {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("fire"))
}
func (s SOPowerSavvyTimer) FireDate() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("fireDate"))
	return objectivec.Object{ID: rv}
}
func (s SOPowerSavvyTimer) Invalidate() {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("invalidate"))
}
func (s SOPowerSavvyTimer) IsValid() bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("isValid"))
	return rv
}
func (s SOPowerSavvyTimer) SetFireDate(date objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setFireDate:"), date)
}
func (s SOPowerSavvyTimer) TimeInterval() float64 {
	rv := objc.SendIfResponds[float64](s.ID, objc.Sel("timeInterval"))
	return rv
}
func (s SOPowerSavvyTimer) UserInfo() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("userInfo"))
	return objectivec.Object{ID: rv}
}

func (_SOPowerSavvyTimerClass SOPowerSavvyTimerClass) RequestTargetPerformSelectorWithObjectAfterDelay(target objectivec.IObject, selector objc.SEL, object objectivec.IObject, delay float64) {
	objc.SendIfResponds[objc.ID](objc.ID(_SOPowerSavvyTimerClass.class), objc.Sel("requestTarget:performSelector:withObject:afterDelay:"), target, selector, object, delay)
}
func (_SOPowerSavvyTimerClass SOPowerSavvyTimerClass) ScheduledTimerWithTimeIntervalTargetSelectorUserInfoRepeats(interval float64, target objectivec.IObject, selector objc.SEL, info objectivec.IObject, repeats bool) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_SOPowerSavvyTimerClass.class), objc.Sel("scheduledTimerWithTimeInterval:target:selector:userInfo:repeats:"), interval, target, selector, info, repeats)
	return objectivec.Object{ID: rv}
}

func (s SOPowerSavvyTimer) Repeats() bool {
	rv := objc.SendIfResponds[bool](s.ID, objc.Sel("repeats"))
	return rv
}
func (s SOPowerSavvyTimer) SetRepeats(value bool) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setRepeats:"), value)
}
func (s SOPowerSavvyTimer) Selector() objectivec.SEL {
	rv := objc.SendIfResponds[objc.SEL](s.ID, objc.Sel("selector"))
	return objectivec.SEL(rv)
}
func (s SOPowerSavvyTimer) SetSelector(value objectivec.SEL) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setSelector:"), value)
}
func (s SOPowerSavvyTimer) Target() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("target"))
	return objectivec.Object{ID: rv}
}
func (s SOPowerSavvyTimer) SetTarget(value objectivec.IObject) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setTarget:"), value)
}
func (s SOPowerSavvyTimer) Timer() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](s.ID, objc.Sel("timer"))
	return rv
}
func (s SOPowerSavvyTimer) SetTimer(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](s.ID, objc.Sel("setTimer:"), value)
}
