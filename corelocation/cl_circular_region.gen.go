// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [CLCircularRegion] class.
var (
	_CLCircularRegionClass     CLCircularRegionClass
	_CLCircularRegionClassOnce sync.Once
)

func getCLCircularRegionClass() CLCircularRegionClass {
	_CLCircularRegionClassOnce.Do(func() {
		_CLCircularRegionClass = CLCircularRegionClass{class: objc.GetClass("CLCircularRegion")}
	})
	return _CLCircularRegionClass
}

// GetCLCircularRegionClass returns the class object for CLCircularRegion.
func GetCLCircularRegionClass() CLCircularRegionClass {
	return getCLCircularRegionClass()
}

type CLCircularRegionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLCircularRegionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLCircularRegionClass) Alloc() CLCircularRegion {
	rv := objc.Send[CLCircularRegion](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A circular geographic region that a center point and radius deine.
//
// # Overview
//
// The [CLCircularRegion] class defines the location and boundaries for a
// circular geographic region. You can use instances of this class to define
// geofences for a specific location. The crossing of a geofence’s boundary
// causes the location manager to notify its delegate.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLCircularRegion
type CLCircularRegion struct {
	CLRegion
}

// CLCircularRegionFromID constructs a [CLCircularRegion] from an objc.ID.
//
// A circular geographic region that a center point and radius deine.
func CLCircularRegionFromID(id objc.ID) CLCircularRegion {
	return CLCircularRegion{CLRegion: CLRegionFromID(id)}
}

// NOTE: CLCircularRegion adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CLCircularRegion] class.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLCircularRegion
type ICLCircularRegion interface {
	ICLRegion
}

// Init initializes the instance.
func (c CLCircularRegion) Init() CLCircularRegion {
	rv := objc.Send[CLCircularRegion](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CLCircularRegion) Autorelease() CLCircularRegion {
	rv := objc.Send[CLCircularRegion](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLCircularRegion creates a new CLCircularRegion instance.
func NewCLCircularRegion() CLCircularRegion {
	class := getCLCircularRegionClass()
	rv := objc.Send[CLCircularRegion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreLocation/CLRegion/init(coder:)
func NewCircularRegionWithCoder(coder foundation.INSCoder) CLCircularRegion {
	instance := getCLCircularRegionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return CLCircularRegionFromID(rv)
}
