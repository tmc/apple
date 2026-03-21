// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [HasTrackingID] class.
var (
	_HasTrackingIDClass     HasTrackingIDClass
	_HasTrackingIDClassOnce sync.Once
)

func getHasTrackingIDClass() HasTrackingIDClass {
	_HasTrackingIDClassOnce.Do(func() {
		_HasTrackingIDClass = HasTrackingIDClass{class: objc.GetClass("hasTrackingID")}
	})
	return _HasTrackingIDClass
}

// GetHasTrackingIDClass returns the class object for hasTrackingID.
func GetHasTrackingIDClass() HasTrackingIDClass {
	return getHasTrackingIDClass()
}

type HasTrackingIDClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (hc HasTrackingIDClass) Alloc() HasTrackingID {
	rv := objc.Send[HasTrackingID](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasTrackingID-c.ivar
type HasTrackingID struct {
	objectivec.Object
}

// HasTrackingIDFromID constructs a [HasTrackingID] from an objc.ID.
func HasTrackingIDFromID(id objc.ID) HasTrackingID {
	return HasTrackingID{objectivec.Object{ID: id}}
}
// Ensure HasTrackingID implements IHasTrackingID.
var _ IHasTrackingID = HasTrackingID{}

// An interface definition for the [HasTrackingID] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasTrackingID-c.ivar
type IHasTrackingID interface {
	objectivec.IObject
}

// Init initializes the instance.
func (h HasTrackingID) Init() HasTrackingID {
	rv := objc.Send[HasTrackingID](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h HasTrackingID) Autorelease() HasTrackingID {
	rv := objc.Send[HasTrackingID](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewHasTrackingID creates a new HasTrackingID instance.
func NewHasTrackingID() HasTrackingID {
	class := getHasTrackingIDClass()
	rv := objc.Send[HasTrackingID](objc.ID(class.class), objc.Sel("new"))
	return rv
}

