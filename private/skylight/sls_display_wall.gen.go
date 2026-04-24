// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSDisplayWall] class.
var (
	_SLSDisplayWallClass     SLSDisplayWallClass
	_SLSDisplayWallClassOnce sync.Once
)

func getSLSDisplayWallClass() SLSDisplayWallClass {
	_SLSDisplayWallClassOnce.Do(func() {
		_SLSDisplayWallClass = SLSDisplayWallClass{class: objc.GetClass("SLSDisplayWall")}
	})
	return _SLSDisplayWallClass
}

// GetSLSDisplayWallClass returns the class object for SLSDisplayWall.
func GetSLSDisplayWallClass() SLSDisplayWallClass {
	return getSLSDisplayWallClass()
}

type SLSDisplayWallClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSDisplayWallClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSDisplayWallClass) Alloc() SLSDisplayWall {
	rv := objc.Send[SLSDisplayWall](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayWall
type SLSDisplayWall struct {
	objectivec.Object
}

// SLSDisplayWallFromID constructs a [SLSDisplayWall] from an objc.ID.
func SLSDisplayWallFromID(id objc.ID) SLSDisplayWall {
	return SLSDisplayWall{objectivec.Object{ID: id}}
}

// Ensure SLSDisplayWall implements ISLSDisplayWall.
var _ ISLSDisplayWall = SLSDisplayWall{}

// An interface definition for the [SLSDisplayWall] class.
//
// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayWall
type ISLSDisplayWall interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SLSDisplayWall) Init() SLSDisplayWall {
	rv := objc.Send[SLSDisplayWall](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSDisplayWall) Autorelease() SLSDisplayWall {
	rv := objc.Send[SLSDisplayWall](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSDisplayWall creates a new SLSDisplayWall instance.
func NewSLSDisplayWall() SLSDisplayWall {
	class := getSLSDisplayWallClass()
	rv := objc.Send[SLSDisplayWall](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayWall/availableDisplayWallGroups
func (_SLSDisplayWallClass SLSDisplayWallClass) AvailableDisplayWallGroups() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_SLSDisplayWallClass.class), objc.Sel("availableDisplayWallGroups"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayWall/disableDisplayWall:
func (_SLSDisplayWallClass SLSDisplayWallClass) DisableDisplayWall(wall objectivec.IObject) int {
	rv := objc.Send[int](objc.ID(_SLSDisplayWallClass.class), objc.Sel("disableDisplayWall:"), wall)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayWall/enableDisplayWallWithConfiguration:
func (_SLSDisplayWallClass SLSDisplayWallClass) EnableDisplayWallWithConfiguration(configuration objectivec.IObject) int {
	rv := objc.Send[int](objc.ID(_SLSDisplayWallClass.class), objc.Sel("enableDisplayWallWithConfiguration:"), configuration)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayWall/isDisplayWall:
func (_SLSDisplayWallClass SLSDisplayWallClass) IsDisplayWall(wall objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_SLSDisplayWallClass.class), objc.Sel("isDisplayWall:"), wall)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayWall/validateDisplayWallConfiguration:
func (_SLSDisplayWallClass SLSDisplayWallClass) ValidateDisplayWallConfiguration(configuration objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_SLSDisplayWallClass.class), objc.Sel("validateDisplayWallConfiguration:"), configuration)
	return rv
}
