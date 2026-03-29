// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TrackingFrameCount] class.
var (
	_TrackingFrameCountClass     TrackingFrameCountClass
	_TrackingFrameCountClassOnce sync.Once
)

func getTrackingFrameCountClass() TrackingFrameCountClass {
	_TrackingFrameCountClassOnce.Do(func() {
		_TrackingFrameCountClass = TrackingFrameCountClass{class: objc.GetClass("trackingFrameCount")}
	})
	return _TrackingFrameCountClass
}

// GetTrackingFrameCountClass returns the class object for trackingFrameCount.
func GetTrackingFrameCountClass() TrackingFrameCountClass {
	return getTrackingFrameCountClass()
}

type TrackingFrameCountClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TrackingFrameCountClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TrackingFrameCountClass) Alloc() TrackingFrameCount {
	rv := objc.Send[TrackingFrameCount](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/trackingFrameCount-c.ivar
type TrackingFrameCount struct {
	objectivec.Object
}

// TrackingFrameCountFromID constructs a [TrackingFrameCount] from an objc.ID.
func TrackingFrameCountFromID(id objc.ID) TrackingFrameCount {
	return TrackingFrameCount{objectivec.Object{ID: id}}
}
// Ensure TrackingFrameCount implements ITrackingFrameCount.
var _ ITrackingFrameCount = TrackingFrameCount{}

// An interface definition for the [TrackingFrameCount] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/trackingFrameCount-c.ivar
type ITrackingFrameCount interface {
	objectivec.IObject
}

// Init initializes the instance.
func (t TrackingFrameCount) Init() TrackingFrameCount {
	rv := objc.Send[TrackingFrameCount](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TrackingFrameCount) Autorelease() TrackingFrameCount {
	rv := objc.Send[TrackingFrameCount](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTrackingFrameCount creates a new TrackingFrameCount instance.
func NewTrackingFrameCount() TrackingFrameCount {
	class := getTrackingFrameCountClass()
	rv := objc.Send[TrackingFrameCount](objc.ID(class.class), objc.Sel("new"))
	return rv
}

