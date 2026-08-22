// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler] class.
var (
	_EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass     EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass
	_EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClassOnce sync.Once
)

func getEA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass() EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass {
	_EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClassOnce.Do(func() {
		_EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass = EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass{class: objc.GetClass("_TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler")}
	})
	return _EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass
}

// GetEA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass returns the class object for _TtC6CoreMLP33_E8DC9B772C73548EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler.
func GetEA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass() EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass {
	return getEA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass()
}

type EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass) Alloc() EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler {
	rv := objc.SendIfResponds[EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

type EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler struct {
	objectivec.Object
}

// EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerFromID constructs a [EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler] from an objc.ID.
func EA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerFromID(id objc.ID) EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler {
	return EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler{objectivec.Object{ID: id}}
}

// Ensure EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler implements IEA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler.
var _ IEA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler = EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler{}

// An interface definition for the [EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler] class.
type IEA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler interface {
	objectivec.IObject
}

// Init initializes the instance.
func (e EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler) Init() EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler {
	rv := objc.SendIfResponds[EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler) Autorelease() EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler {
	rv := objc.SendIfResponds[EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler creates a new EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler instance.
func NewEA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler() EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler {
	class := getEA5F31AD4BCEEE43A29AsyncComputeFunctionSchedulerClass()
	rv := objc.SendIfResponds[EA5F31AD4BCEEE43A29AsyncComputeFunctionScheduler](objc.ID(class.class), objc.Sel("new"))
	return rv
}
