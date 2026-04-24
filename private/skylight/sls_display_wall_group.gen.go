// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSDisplayWallGroup] class.
var (
	_SLSDisplayWallGroupClass     SLSDisplayWallGroupClass
	_SLSDisplayWallGroupClassOnce sync.Once
)

func getSLSDisplayWallGroupClass() SLSDisplayWallGroupClass {
	_SLSDisplayWallGroupClassOnce.Do(func() {
		_SLSDisplayWallGroupClass = SLSDisplayWallGroupClass{class: objc.GetClass("SLSDisplayWallGroup")}
	})
	return _SLSDisplayWallGroupClass
}

// GetSLSDisplayWallGroupClass returns the class object for SLSDisplayWallGroup.
func GetSLSDisplayWallGroupClass() SLSDisplayWallGroupClass {
	return getSLSDisplayWallGroupClass()
}

type SLSDisplayWallGroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSDisplayWallGroupClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSDisplayWallGroupClass) Alloc() SLSDisplayWallGroup {
	rv := objc.Send[SLSDisplayWallGroup](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSDisplayWallGroup.DisplayIDs]
//   - [SLSDisplayWallGroup.GroupID]
//   - [SLSDisplayWallGroup.InitWithCGDisplayWallGroup]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayWallGroup
type SLSDisplayWallGroup struct {
	objectivec.Object
}

// SLSDisplayWallGroupFromID constructs a [SLSDisplayWallGroup] from an objc.ID.
func SLSDisplayWallGroupFromID(id objc.ID) SLSDisplayWallGroup {
	return SLSDisplayWallGroup{objectivec.Object{ID: id}}
}

// Ensure SLSDisplayWallGroup implements ISLSDisplayWallGroup.
var _ ISLSDisplayWallGroup = SLSDisplayWallGroup{}

// An interface definition for the [SLSDisplayWallGroup] class.
//
// # Methods
//
//   - [ISLSDisplayWallGroup.DisplayIDs]
//   - [ISLSDisplayWallGroup.GroupID]
//   - [ISLSDisplayWallGroup.InitWithCGDisplayWallGroup]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayWallGroup
type ISLSDisplayWallGroup interface {
	objectivec.IObject

	// Topic: Methods

	DisplayIDs() foundation.INSArray
	GroupID() foundation.NSNumber
	InitWithCGDisplayWallGroup(group objectivec.IObject) SLSDisplayWallGroup
}

// Init initializes the instance.
func (s SLSDisplayWallGroup) Init() SLSDisplayWallGroup {
	rv := objc.Send[SLSDisplayWallGroup](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSDisplayWallGroup) Autorelease() SLSDisplayWallGroup {
	rv := objc.Send[SLSDisplayWallGroup](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSDisplayWallGroup creates a new SLSDisplayWallGroup instance.
func NewSLSDisplayWallGroup() SLSDisplayWallGroup {
	class := getSLSDisplayWallGroupClass()
	rv := objc.Send[SLSDisplayWallGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayWallGroup/initWithCGDisplayWallGroup:
func NewSLSDisplayWallGroupWithCGDisplayWallGroup(group objectivec.IObject) SLSDisplayWallGroup {
	instance := getSLSDisplayWallGroupClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCGDisplayWallGroup:"), group)
	return SLSDisplayWallGroupFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayWallGroup/initWithCGDisplayWallGroup:
func (s SLSDisplayWallGroup) InitWithCGDisplayWallGroup(group objectivec.IObject) SLSDisplayWallGroup {
	rv := objc.Send[SLSDisplayWallGroup](s.ID, objc.Sel("initWithCGDisplayWallGroup:"), group)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayWallGroup/displayIDs
func (s SLSDisplayWallGroup) DisplayIDs() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayIDs"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSDisplayWallGroup/groupID
func (s SLSDisplayWallGroup) GroupID() foundation.NSNumber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("groupID"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
