// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightIndirectSystemGestureEventProcessor] class.
var (
	_SkyLightIndirectSystemGestureEventProcessorClass     SkyLightIndirectSystemGestureEventProcessorClass
	_SkyLightIndirectSystemGestureEventProcessorClassOnce sync.Once
)

func getSkyLightIndirectSystemGestureEventProcessorClass() SkyLightIndirectSystemGestureEventProcessorClass {
	_SkyLightIndirectSystemGestureEventProcessorClassOnce.Do(func() {
		_SkyLightIndirectSystemGestureEventProcessorClass = SkyLightIndirectSystemGestureEventProcessorClass{class: objc.GetClass("SkyLight.IndirectSystemGestureEventProcessor")}
	})
	return _SkyLightIndirectSystemGestureEventProcessorClass
}

// GetSkyLightIndirectSystemGestureEventProcessorClass returns the class object for SkyLight.IndirectSystemGestureEventProcessor.
func GetSkyLightIndirectSystemGestureEventProcessorClass() SkyLightIndirectSystemGestureEventProcessorClass {
	return getSkyLightIndirectSystemGestureEventProcessorClass()
}

type SkyLightIndirectSystemGestureEventProcessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightIndirectSystemGestureEventProcessorClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightIndirectSystemGestureEventProcessorClass) Alloc() SkyLightIndirectSystemGestureEventProcessor {
	rv := objc.SendIfResponds[SkyLightIndirectSystemGestureEventProcessor](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type SkyLightIndirectSystemGestureEventProcessor struct {
	objectivec.Object
}

// SkyLightIndirectSystemGestureEventProcessorFromID constructs a [SkyLightIndirectSystemGestureEventProcessor] from an objc.ID.
func SkyLightIndirectSystemGestureEventProcessorFromID(id objc.ID) SkyLightIndirectSystemGestureEventProcessor {
	return SkyLightIndirectSystemGestureEventProcessor{objectivec.Object{ID: id}}
}

// Ensure SkyLightIndirectSystemGestureEventProcessor implements ISkyLightIndirectSystemGestureEventProcessor.
var _ ISkyLightIndirectSystemGestureEventProcessor = SkyLightIndirectSystemGestureEventProcessor{}

// An interface definition for the [SkyLightIndirectSystemGestureEventProcessor] class.
type ISkyLightIndirectSystemGestureEventProcessor interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightIndirectSystemGestureEventProcessor) Init() SkyLightIndirectSystemGestureEventProcessor {
	rv := objc.SendIfResponds[SkyLightIndirectSystemGestureEventProcessor](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightIndirectSystemGestureEventProcessor) Autorelease() SkyLightIndirectSystemGestureEventProcessor {
	rv := objc.SendIfResponds[SkyLightIndirectSystemGestureEventProcessor](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightIndirectSystemGestureEventProcessor creates a new SkyLightIndirectSystemGestureEventProcessor instance.
func NewSkyLightIndirectSystemGestureEventProcessor() SkyLightIndirectSystemGestureEventProcessor {
	class := getSkyLightIndirectSystemGestureEventProcessorClass()
	rv := objc.SendIfResponds[SkyLightIndirectSystemGestureEventProcessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}
