// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightSystemGestureClientWindow] class.
var (
	_SkyLightSystemGestureClientWindowClass     SkyLightSystemGestureClientWindowClass
	_SkyLightSystemGestureClientWindowClassOnce sync.Once
)

func getSkyLightSystemGestureClientWindowClass() SkyLightSystemGestureClientWindowClass {
	_SkyLightSystemGestureClientWindowClassOnce.Do(func() {
		_SkyLightSystemGestureClientWindowClass = SkyLightSystemGestureClientWindowClass{class: objc.GetClass("SkyLight.SystemGestureClientWindow")}
	})
	return _SkyLightSystemGestureClientWindowClass
}

// GetSkyLightSystemGestureClientWindowClass returns the class object for SkyLight.SystemGestureClientWindow.
func GetSkyLightSystemGestureClientWindowClass() SkyLightSystemGestureClientWindowClass {
	return getSkyLightSystemGestureClientWindowClass()
}

type SkyLightSystemGestureClientWindowClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightSystemGestureClientWindowClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightSystemGestureClientWindowClass) Alloc() SkyLightSystemGestureClientWindow {
	rv := objc.SendIfResponds[SkyLightSystemGestureClientWindow](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type SkyLightSystemGestureClientWindow struct {
	objectivec.Object
}

// SkyLightSystemGestureClientWindowFromID constructs a [SkyLightSystemGestureClientWindow] from an objc.ID.
func SkyLightSystemGestureClientWindowFromID(id objc.ID) SkyLightSystemGestureClientWindow {
	return SkyLightSystemGestureClientWindow{objectivec.Object{ID: id}}
}

// Ensure SkyLightSystemGestureClientWindow implements ISkyLightSystemGestureClientWindow.
var _ ISkyLightSystemGestureClientWindow = SkyLightSystemGestureClientWindow{}

// An interface definition for the [SkyLightSystemGestureClientWindow] class.
type ISkyLightSystemGestureClientWindow interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightSystemGestureClientWindow) Init() SkyLightSystemGestureClientWindow {
	rv := objc.SendIfResponds[SkyLightSystemGestureClientWindow](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightSystemGestureClientWindow) Autorelease() SkyLightSystemGestureClientWindow {
	rv := objc.SendIfResponds[SkyLightSystemGestureClientWindow](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightSystemGestureClientWindow creates a new SkyLightSystemGestureClientWindow instance.
func NewSkyLightSystemGestureClientWindow() SkyLightSystemGestureClientWindow {
	class := getSkyLightSystemGestureClientWindowClass()
	rv := objc.SendIfResponds[SkyLightSystemGestureClientWindow](objc.ID(class.class), objc.Sel("new"))
	return rv
}
