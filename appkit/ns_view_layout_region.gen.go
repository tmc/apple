// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSViewLayoutRegion] class.
var (
	_NSViewLayoutRegionClass     NSViewLayoutRegionClass
	_NSViewLayoutRegionClassOnce sync.Once
)

func getNSViewLayoutRegionClass() NSViewLayoutRegionClass {
	_NSViewLayoutRegionClassOnce.Do(func() {
		_NSViewLayoutRegionClass = NSViewLayoutRegionClass{class: objc.GetClass("NSViewLayoutRegion")}
	})
	return _NSViewLayoutRegionClass
}

// GetNSViewLayoutRegionClass returns the class object for NSViewLayoutRegion.
func GetNSViewLayoutRegionClass() NSViewLayoutRegionClass {
	return getNSViewLayoutRegionClass()
}

type NSViewLayoutRegionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSViewLayoutRegionClass) Alloc() NSViewLayoutRegion {
	rv := objc.Send[NSViewLayoutRegion](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegion
type NSViewLayoutRegion struct {
	objectivec.Object
}

// NSViewLayoutRegionFromID constructs a [NSViewLayoutRegion] from an objc.ID.
func NSViewLayoutRegionFromID(id objc.ID) NSViewLayoutRegion {
	return NSViewLayoutRegion{objectivec.Object{ID: id}}
}
// Ensure NSViewLayoutRegion implements INSViewLayoutRegion.
var _ INSViewLayoutRegion = NSViewLayoutRegion{}

// An interface definition for the [NSViewLayoutRegion] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegion
type INSViewLayoutRegion interface {
	objectivec.IObject
}

// Init initializes the instance.
func (v NSViewLayoutRegion) Init() NSViewLayoutRegion {
	rv := objc.Send[NSViewLayoutRegion](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v NSViewLayoutRegion) Autorelease() NSViewLayoutRegion {
	rv := objc.Send[NSViewLayoutRegion](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSViewLayoutRegion creates a new NSViewLayoutRegion instance.
func NewNSViewLayoutRegion() NSViewLayoutRegion {
	class := getNSViewLayoutRegionClass()
	rv := objc.Send[NSViewLayoutRegion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegion/marginsLayoutRegionWithCornerAdaptation:
func (_NSViewLayoutRegionClass NSViewLayoutRegionClass) MarginsLayoutRegionWithCornerAdaptation(adaptivityAxis NSViewLayoutRegionAdaptivityAxis) NSViewLayoutRegion {
	rv := objc.Send[objc.ID](objc.ID(_NSViewLayoutRegionClass.class), objc.Sel("marginsLayoutRegionWithCornerAdaptation:"), adaptivityAxis)
	return NSViewLayoutRegionFromID(rv)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegion/safeAreaLayoutRegionWithCornerAdaptation:
func (_NSViewLayoutRegionClass NSViewLayoutRegionClass) SafeAreaLayoutRegionWithCornerAdaptation(adaptivityAxis NSViewLayoutRegionAdaptivityAxis) NSViewLayoutRegion {
	rv := objc.Send[objc.ID](objc.ID(_NSViewLayoutRegionClass.class), objc.Sel("safeAreaLayoutRegionWithCornerAdaptation:"), adaptivityAxis)
	return NSViewLayoutRegionFromID(rv)
}

