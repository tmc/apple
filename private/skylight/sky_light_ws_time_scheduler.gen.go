// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightWSTimeScheduler] class.
var (
	_SkyLightWSTimeSchedulerClass     SkyLightWSTimeSchedulerClass
	_SkyLightWSTimeSchedulerClassOnce sync.Once
)

func getSkyLightWSTimeSchedulerClass() SkyLightWSTimeSchedulerClass {
	_SkyLightWSTimeSchedulerClassOnce.Do(func() {
		_SkyLightWSTimeSchedulerClass = SkyLightWSTimeSchedulerClass{class: objc.GetClass("SkyLight.WSTimeScheduler")}
	})
	return _SkyLightWSTimeSchedulerClass
}

// GetSkyLightWSTimeSchedulerClass returns the class object for SkyLight.WSTimeScheduler.
func GetSkyLightWSTimeSchedulerClass() SkyLightWSTimeSchedulerClass {
	return getSkyLightWSTimeSchedulerClass()
}

type SkyLightWSTimeSchedulerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightWSTimeSchedulerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightWSTimeSchedulerClass) Alloc() SkyLightWSTimeScheduler {
	rv := objc.SendIfResponds[SkyLightWSTimeScheduler](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type SkyLightWSTimeScheduler struct {
	objectivec.Object
}

// SkyLightWSTimeSchedulerFromID constructs a [SkyLightWSTimeScheduler] from an objc.ID.
func SkyLightWSTimeSchedulerFromID(id objc.ID) SkyLightWSTimeScheduler {
	return SkyLightWSTimeScheduler{objectivec.Object{ID: id}}
}

// Ensure SkyLightWSTimeScheduler implements ISkyLightWSTimeScheduler.
var _ ISkyLightWSTimeScheduler = SkyLightWSTimeScheduler{}

// An interface definition for the [SkyLightWSTimeScheduler] class.
type ISkyLightWSTimeScheduler interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightWSTimeScheduler) Init() SkyLightWSTimeScheduler {
	rv := objc.SendIfResponds[SkyLightWSTimeScheduler](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightWSTimeScheduler) Autorelease() SkyLightWSTimeScheduler {
	rv := objc.SendIfResponds[SkyLightWSTimeScheduler](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightWSTimeScheduler creates a new SkyLightWSTimeScheduler instance.
func NewSkyLightWSTimeScheduler() SkyLightWSTimeScheduler {
	class := getSkyLightWSTimeSchedulerClass()
	rv := objc.SendIfResponds[SkyLightWSTimeScheduler](objc.ID(class.class), objc.Sel("new"))
	return rv
}
