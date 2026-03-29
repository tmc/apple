// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLBackgroundWatchdog] class.
var (
	_MLBackgroundWatchdogClass     MLBackgroundWatchdogClass
	_MLBackgroundWatchdogClassOnce sync.Once
)

func getMLBackgroundWatchdogClass() MLBackgroundWatchdogClass {
	_MLBackgroundWatchdogClassOnce.Do(func() {
		_MLBackgroundWatchdogClass = MLBackgroundWatchdogClass{class: objc.GetClass("MLBackgroundWatchdog")}
	})
	return _MLBackgroundWatchdogClass
}

// GetMLBackgroundWatchdogClass returns the class object for MLBackgroundWatchdog.
func GetMLBackgroundWatchdogClass() MLBackgroundWatchdogClass {
	return getMLBackgroundWatchdogClass()
}

type MLBackgroundWatchdogClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLBackgroundWatchdogClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLBackgroundWatchdogClass) Alloc() MLBackgroundWatchdog {
	rv := objc.Send[MLBackgroundWatchdog](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLBackgroundWatchdog.Invalidate]
//   - [MLBackgroundWatchdog.Timer]
//   - [MLBackgroundWatchdog.SetTimer]
// See: https://developer.apple.com/documentation/CoreML/MLBackgroundWatchdog
type MLBackgroundWatchdog struct {
	objectivec.Object
}

// MLBackgroundWatchdogFromID constructs a [MLBackgroundWatchdog] from an objc.ID.
func MLBackgroundWatchdogFromID(id objc.ID) MLBackgroundWatchdog {
	return MLBackgroundWatchdog{objectivec.Object{ID: id}}
}
// Ensure MLBackgroundWatchdog implements IMLBackgroundWatchdog.
var _ IMLBackgroundWatchdog = MLBackgroundWatchdog{}

// An interface definition for the [MLBackgroundWatchdog] class.
//
// # Methods
//
//   - [IMLBackgroundWatchdog.Invalidate]
//   - [IMLBackgroundWatchdog.Timer]
//   - [IMLBackgroundWatchdog.SetTimer]
//
// See: https://developer.apple.com/documentation/CoreML/MLBackgroundWatchdog
type IMLBackgroundWatchdog interface {
	objectivec.IObject

	// Topic: Methods

	Invalidate()
	Timer() objectivec.Object
	SetTimer(value objectivec.Object)
}

// Init initializes the instance.
func (b MLBackgroundWatchdog) Init() MLBackgroundWatchdog {
	rv := objc.Send[MLBackgroundWatchdog](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MLBackgroundWatchdog) Autorelease() MLBackgroundWatchdog {
	rv := objc.Send[MLBackgroundWatchdog](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLBackgroundWatchdog creates a new MLBackgroundWatchdog instance.
func NewMLBackgroundWatchdog() MLBackgroundWatchdog {
	class := getMLBackgroundWatchdogClass()
	rv := objc.Send[MLBackgroundWatchdog](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundWatchdog/invalidate
func (b MLBackgroundWatchdog) Invalidate() {
	objc.Send[objc.ID](b.ID, objc.Sel("invalidate"))
}

//
// See: https://developer.apple.com/documentation/CoreML/MLBackgroundWatchdog/watchdogWithTimeout:label:queue:
func (_MLBackgroundWatchdogClass MLBackgroundWatchdogClass) WatchdogWithTimeoutLabelQueue(timeout float64, label objectivec.IObject, queue objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLBackgroundWatchdogClass.class), objc.Sel("watchdogWithTimeout:label:queue:"), timeout, label, queue)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBackgroundWatchdog/watchdogWithTimeout:queue:
func (_MLBackgroundWatchdogClass MLBackgroundWatchdogClass) WatchdogWithTimeoutQueue(timeout float64, queue objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLBackgroundWatchdogClass.class), objc.Sel("watchdogWithTimeout:queue:"), timeout, queue)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLBackgroundWatchdog/timer
func (b MLBackgroundWatchdog) Timer() objectivec.Object {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("timer"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (b MLBackgroundWatchdog) SetTimer(value objectivec.Object) {
	objc.Send[struct{}](b.ID, objc.Sel("setTimer:"), value)
}

