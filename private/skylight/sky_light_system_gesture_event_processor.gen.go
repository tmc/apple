// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightSystemGestureEventProcessor] class.
var (
	_SkyLightSystemGestureEventProcessorClass     SkyLightSystemGestureEventProcessorClass
	_SkyLightSystemGestureEventProcessorClassOnce sync.Once
)

func getSkyLightSystemGestureEventProcessorClass() SkyLightSystemGestureEventProcessorClass {
	_SkyLightSystemGestureEventProcessorClassOnce.Do(func() {
		_SkyLightSystemGestureEventProcessorClass = SkyLightSystemGestureEventProcessorClass{class: objc.GetClass("SkyLight.SystemGestureEventProcessor")}
	})
	return _SkyLightSystemGestureEventProcessorClass
}

// GetSkyLightSystemGestureEventProcessorClass returns the class object for SkyLight.SystemGestureEventProcessor.
func GetSkyLightSystemGestureEventProcessorClass() SkyLightSystemGestureEventProcessorClass {
	return getSkyLightSystemGestureEventProcessorClass()
}

type SkyLightSystemGestureEventProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightSystemGestureEventProcessorClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightSystemGestureEventProcessorClass) Alloc() SkyLightSystemGestureEventProcessor {
	rv := objc.SendIfResponds[SkyLightSystemGestureEventProcessor](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type SkyLightSystemGestureEventProcessor struct {
	objectivec.Object
}

// SkyLightSystemGestureEventProcessorFromID constructs a [SkyLightSystemGestureEventProcessor] from an objc.ID.
func SkyLightSystemGestureEventProcessorFromID(id objc.ID) SkyLightSystemGestureEventProcessor {
	return SkyLightSystemGestureEventProcessor{objectivec.Object{ID: id}}
}

// Ensure SkyLightSystemGestureEventProcessor implements ISkyLightSystemGestureEventProcessor.
var _ ISkyLightSystemGestureEventProcessor = SkyLightSystemGestureEventProcessor{}

// An interface definition for the [SkyLightSystemGestureEventProcessor] class.
type ISkyLightSystemGestureEventProcessor interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightSystemGestureEventProcessor) Init() SkyLightSystemGestureEventProcessor {
	rv := objc.SendIfResponds[SkyLightSystemGestureEventProcessor](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightSystemGestureEventProcessor) Autorelease() SkyLightSystemGestureEventProcessor {
	rv := objc.SendIfResponds[SkyLightSystemGestureEventProcessor](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightSystemGestureEventProcessor creates a new SkyLightSystemGestureEventProcessor instance.
func NewSkyLightSystemGestureEventProcessor() SkyLightSystemGestureEventProcessor {
	class := getSkyLightSystemGestureEventProcessorClass()
	rv := objc.SendIfResponds[SkyLightSystemGestureEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}
