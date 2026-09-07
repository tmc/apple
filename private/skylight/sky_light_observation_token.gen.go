// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightObservationToken] class.
var (
	_SkyLightObservationTokenClass     SkyLightObservationTokenClass
	_SkyLightObservationTokenClassOnce sync.Once
)

func getSkyLightObservationTokenClass() SkyLightObservationTokenClass {
	_SkyLightObservationTokenClassOnce.Do(func() {
		_SkyLightObservationTokenClass = SkyLightObservationTokenClass{class: objc.GetClass("SkyLight.ObservationToken")}
	})
	return _SkyLightObservationTokenClass
}

// GetSkyLightObservationTokenClass returns the class object for SkyLight.ObservationToken.
func GetSkyLightObservationTokenClass() SkyLightObservationTokenClass {
	return getSkyLightObservationTokenClass()
}

type SkyLightObservationTokenClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightObservationTokenClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightObservationTokenClass) Alloc() SkyLightObservationToken {
	rv := objc.SendIfResponds[SkyLightObservationToken](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type SkyLightObservationToken struct {
	objectivec.Object
}

// SkyLightObservationTokenFromID constructs a [SkyLightObservationToken] from an objc.ID.
func SkyLightObservationTokenFromID(id objc.ID) SkyLightObservationToken {
	return SkyLightObservationToken{objectivec.Object{ID: id}}
}

// Ensure SkyLightObservationToken implements ISkyLightObservationToken.
var _ ISkyLightObservationToken = SkyLightObservationToken{}

// An interface definition for the [SkyLightObservationToken] class.
type ISkyLightObservationToken interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightObservationToken) Init() SkyLightObservationToken {
	rv := objc.SendIfResponds[SkyLightObservationToken](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightObservationToken) Autorelease() SkyLightObservationToken {
	rv := objc.SendIfResponds[SkyLightObservationToken](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightObservationToken creates a new SkyLightObservationToken instance.
func NewSkyLightObservationToken() SkyLightObservationToken {
	class := getSkyLightObservationTokenClass()
	rv := objc.SendIfResponds[SkyLightObservationToken](objc.ID(class.class), objc.Sel("new"))
	return rv
}
