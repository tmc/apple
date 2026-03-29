// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler] class.
var (
	_TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass     TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass
	_TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClassOnce sync.Once
)

func getTtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass() TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass {
	_TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClassOnce.Do(func() {
		_TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass = TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass{class: objc.GetClass("_TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler")}
	})
	return _TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass
}

// GetTtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass returns the class object for _TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler.
func GetTtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass() TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass {
	return getTtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass()
}

type TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass) Alloc() TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler {
	rv := objc.Send[TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler
type TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler struct {
	objectivec.Object
}

// TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerFromID constructs a [TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler] from an objc.ID.
func TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerFromID(id objc.ID) TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler {
	return TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler{objectivec.Object{ID: id}}
}
// NOTE: TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler struct embeds objectivec.Object (parent type unavailable) but
// ITtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler embeds the parent interface; skip compile-time assertion.

// An interface definition for the [TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler] class.
//
// See: https://developer.apple.com/documentation/CoreML/_TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler
type ITtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler) Init() TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler {
	rv := objc.Send[TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler) Autorelease() TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler {
	rv := objc.Send[TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler creates a new TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler instance.
func NewTtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler() TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler {
	class := getTtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass()
	rv := objc.Send[TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

