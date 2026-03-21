// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [HasTrackingFrameCount] class.
var (
	_HasTrackingFrameCountClass     HasTrackingFrameCountClass
	_HasTrackingFrameCountClassOnce sync.Once
)

func getHasTrackingFrameCountClass() HasTrackingFrameCountClass {
	_HasTrackingFrameCountClassOnce.Do(func() {
		_HasTrackingFrameCountClass = HasTrackingFrameCountClass{class: objc.GetClass("hasTrackingFrameCount")}
	})
	return _HasTrackingFrameCountClass
}

// GetHasTrackingFrameCountClass returns the class object for hasTrackingFrameCount.
func GetHasTrackingFrameCountClass() HasTrackingFrameCountClass {
	return getHasTrackingFrameCountClass()
}

type HasTrackingFrameCountClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (hc HasTrackingFrameCountClass) Alloc() HasTrackingFrameCount {
	rv := objc.Send[HasTrackingFrameCount](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasTrackingFrameCount-c.ivar
type HasTrackingFrameCount struct {
	objectivec.Object
}

// HasTrackingFrameCountFromID constructs a [HasTrackingFrameCount] from an objc.ID.
func HasTrackingFrameCountFromID(id objc.ID) HasTrackingFrameCount {
	return HasTrackingFrameCount{objectivec.Object{ID: id}}
}
// Ensure HasTrackingFrameCount implements IHasTrackingFrameCount.
var _ IHasTrackingFrameCount = HasTrackingFrameCount{}

// An interface definition for the [HasTrackingFrameCount] class.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasTrackingFrameCount-c.ivar
type IHasTrackingFrameCount interface {
	objectivec.IObject
}

// Init initializes the instance.
func (h HasTrackingFrameCount) Init() HasTrackingFrameCount {
	rv := objc.Send[HasTrackingFrameCount](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h HasTrackingFrameCount) Autorelease() HasTrackingFrameCount {
	rv := objc.Send[HasTrackingFrameCount](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewHasTrackingFrameCount creates a new HasTrackingFrameCount instance.
func NewHasTrackingFrameCount() HasTrackingFrameCount {
	class := getHasTrackingFrameCountClass()
	rv := objc.Send[HasTrackingFrameCount](objc.ID(class.class), objc.Sel("new"))
	return rv
}

