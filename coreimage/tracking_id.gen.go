// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TrackingID] class.
var (
	_TrackingIDClass     TrackingIDClass
	_TrackingIDClassOnce sync.Once
)

func getTrackingIDClass() TrackingIDClass {
	_TrackingIDClassOnce.Do(func() {
		_TrackingIDClass = TrackingIDClass{class: objc.GetClass("trackingID")}
	})
	return _TrackingIDClass
}

// GetTrackingIDClass returns the class object for trackingID.
func GetTrackingIDClass() TrackingIDClass {
	return getTrackingIDClass()
}

type TrackingIDClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TrackingIDClass) Alloc() TrackingID {
	rv := objc.Send[TrackingID](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/trackingID-c.ivar
type TrackingID struct {
	objectivec.Object
}

// TrackingIDFromID constructs a [TrackingID] from an objc.ID.
func TrackingIDFromID(id objc.ID) TrackingID {
	return TrackingID{objectivec.Object{ID: id}}
}
// Ensure TrackingID implements ITrackingID.
var _ ITrackingID = TrackingID{}

// An interface definition for the [TrackingID] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/trackingID-c.ivar
type ITrackingID interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TrackingID) Init() TrackingID {
	rv := objc.Send[TrackingID](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TrackingID) Autorelease() TrackingID {
	rv := objc.Send[TrackingID](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTrackingID creates a new TrackingID instance.
func NewTrackingID() TrackingID {
	class := getTrackingIDClass()
	rv := objc.Send[TrackingID](objc.ID(class.class), objc.Sel("new"))
	return rv
}

