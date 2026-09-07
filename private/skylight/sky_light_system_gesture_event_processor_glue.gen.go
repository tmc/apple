// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightSystemGestureEventProcessorGlue] class.
var (
	_SkyLightSystemGestureEventProcessorGlueClass     SkyLightSystemGestureEventProcessorGlueClass
	_SkyLightSystemGestureEventProcessorGlueClassOnce sync.Once
)

func getSkyLightSystemGestureEventProcessorGlueClass() SkyLightSystemGestureEventProcessorGlueClass {
	_SkyLightSystemGestureEventProcessorGlueClassOnce.Do(func() {
		_SkyLightSystemGestureEventProcessorGlueClass = SkyLightSystemGestureEventProcessorGlueClass{class: objc.GetClass("SkyLight.SystemGestureEventProcessorGlue")}
	})
	return _SkyLightSystemGestureEventProcessorGlueClass
}

// GetSkyLightSystemGestureEventProcessorGlueClass returns the class object for SkyLight.SystemGestureEventProcessorGlue.
func GetSkyLightSystemGestureEventProcessorGlueClass() SkyLightSystemGestureEventProcessorGlueClass {
	return getSkyLightSystemGestureEventProcessorGlueClass()
}

type SkyLightSystemGestureEventProcessorGlueClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightSystemGestureEventProcessorGlueClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightSystemGestureEventProcessorGlueClass) Alloc() SkyLightSystemGestureEventProcessorGlue {
	rv := objc.SendIfResponds[SkyLightSystemGestureEventProcessorGlue](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type SkyLightSystemGestureEventProcessorGlue struct {
	objectivec.Object
}

// SkyLightSystemGestureEventProcessorGlueFromID constructs a [SkyLightSystemGestureEventProcessorGlue] from an objc.ID.
func SkyLightSystemGestureEventProcessorGlueFromID(id objc.ID) SkyLightSystemGestureEventProcessorGlue {
	return SkyLightSystemGestureEventProcessorGlue{objectivec.Object{ID: id}}
}

// Ensure SkyLightSystemGestureEventProcessorGlue implements ISkyLightSystemGestureEventProcessorGlue.
var _ ISkyLightSystemGestureEventProcessorGlue = SkyLightSystemGestureEventProcessorGlue{}

// An interface definition for the [SkyLightSystemGestureEventProcessorGlue] class.
type ISkyLightSystemGestureEventProcessorGlue interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightSystemGestureEventProcessorGlue) Init() SkyLightSystemGestureEventProcessorGlue {
	rv := objc.SendIfResponds[SkyLightSystemGestureEventProcessorGlue](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightSystemGestureEventProcessorGlue) Autorelease() SkyLightSystemGestureEventProcessorGlue {
	rv := objc.SendIfResponds[SkyLightSystemGestureEventProcessorGlue](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightSystemGestureEventProcessorGlue creates a new SkyLightSystemGestureEventProcessorGlue instance.
func NewSkyLightSystemGestureEventProcessorGlue() SkyLightSystemGestureEventProcessorGlue {
	class := getSkyLightSystemGestureEventProcessorGlueClass()
	rv := objc.SendIfResponds[SkyLightSystemGestureEventProcessorGlue](objc.ID(class.class), objc.Sel("new"))
	return rv
}
