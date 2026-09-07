// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightPKGWindowInteraction] class.
var (
	_SkyLightPKGWindowInteractionClass     SkyLightPKGWindowInteractionClass
	_SkyLightPKGWindowInteractionClassOnce sync.Once
)

func getSkyLightPKGWindowInteractionClass() SkyLightPKGWindowInteractionClass {
	_SkyLightPKGWindowInteractionClassOnce.Do(func() {
		_SkyLightPKGWindowInteractionClass = SkyLightPKGWindowInteractionClass{class: objc.GetClass("SkyLight.PKGWindowInteraction")}
	})
	return _SkyLightPKGWindowInteractionClass
}

// GetSkyLightPKGWindowInteractionClass returns the class object for SkyLight.PKGWindowInteraction.
func GetSkyLightPKGWindowInteractionClass() SkyLightPKGWindowInteractionClass {
	return getSkyLightPKGWindowInteractionClass()
}

type SkyLightPKGWindowInteractionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightPKGWindowInteractionClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightPKGWindowInteractionClass) Alloc() SkyLightPKGWindowInteraction {
	rv := objc.SendIfResponds[SkyLightPKGWindowInteraction](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type SkyLightPKGWindowInteraction struct {
	objectivec.Object
}

// SkyLightPKGWindowInteractionFromID constructs a [SkyLightPKGWindowInteraction] from an objc.ID.
func SkyLightPKGWindowInteractionFromID(id objc.ID) SkyLightPKGWindowInteraction {
	return SkyLightPKGWindowInteraction{objectivec.Object{ID: id}}
}

// Ensure SkyLightPKGWindowInteraction implements ISkyLightPKGWindowInteraction.
var _ ISkyLightPKGWindowInteraction = SkyLightPKGWindowInteraction{}

// An interface definition for the [SkyLightPKGWindowInteraction] class.
type ISkyLightPKGWindowInteraction interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightPKGWindowInteraction) Init() SkyLightPKGWindowInteraction {
	rv := objc.SendIfResponds[SkyLightPKGWindowInteraction](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightPKGWindowInteraction) Autorelease() SkyLightPKGWindowInteraction {
	rv := objc.SendIfResponds[SkyLightPKGWindowInteraction](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightPKGWindowInteraction creates a new SkyLightPKGWindowInteraction instance.
func NewSkyLightPKGWindowInteraction() SkyLightPKGWindowInteraction {
	class := getSkyLightPKGWindowInteractionClass()
	rv := objc.SendIfResponds[SkyLightPKGWindowInteraction](objc.ID(class.class), objc.Sel("new"))
	return rv
}
