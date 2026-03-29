// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZTouch] class.
var (
	_VZTouchClass     VZTouchClass
	_VZTouchClassOnce sync.Once
)

func getVZTouchClass() VZTouchClass {
	_VZTouchClassOnce.Do(func() {
		_VZTouchClass = VZTouchClass{class: objc.GetClass("_VZTouch")}
	})
	return _VZTouchClass
}

// GetVZTouchClass returns the class object for _VZTouch.
func GetVZTouchClass() VZTouchClass {
	return getVZTouchClass()
}

type VZTouchClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZTouchClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZTouchClass) Alloc() VZTouch {
	rv := objc.Send[VZTouch](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZTouch.Location]
//   - [VZTouch.Phase]
//   - [VZTouch.SwipeAim]
//   - [VZTouch.Timestamp]
//   - [VZTouch.InitWithViewIndexPhaseLocationSwipeAimTimestamp]
// See: https://developer.apple.com/documentation/Virtualization/_VZTouch
type VZTouch struct {
	objectivec.Object
}

// VZTouchFromID constructs a [VZTouch] from an objc.ID.
func VZTouchFromID(id objc.ID) VZTouch {
	return VZTouch{objectivec.Object{ID: id}}
}
// Ensure VZTouch implements IVZTouch.
var _ IVZTouch = VZTouch{}

// An interface definition for the [VZTouch] class.
//
// # Methods
//
//   - [IVZTouch.Location]
//   - [IVZTouch.Phase]
//   - [IVZTouch.SwipeAim]
//   - [IVZTouch.Timestamp]
//   - [IVZTouch.InitWithViewIndexPhaseLocationSwipeAimTimestamp]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZTouch
type IVZTouch interface {
	objectivec.IObject

	// Topic: Methods

	Location() corefoundation.CGPoint
	Phase() int64
	SwipeAim() int64
	Timestamp() float64
	InitWithViewIndexPhaseLocationSwipeAimTimestamp(view objectivec.IObject, index byte, phase int64, location corefoundation.CGPoint, aim int64, timestamp float64) VZTouch
}

// Init initializes the instance.
func (v VZTouch) Init() VZTouch {
	rv := objc.Send[VZTouch](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZTouch) Autorelease() VZTouch {
	rv := objc.Send[VZTouch](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZTouch creates a new VZTouch instance.
func NewVZTouch() VZTouch {
	class := getVZTouchClass()
	rv := objc.Send[VZTouch](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZTouch/initWithView:index:phase:location:swipeAim:timestamp:
func NewVZTouchWithViewIndexPhaseLocationSwipeAimTimestamp(view objectivec.IObject, index byte, phase int64, location corefoundation.CGPoint, aim int64, timestamp float64) VZTouch {
	instance := getVZTouchClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithView:index:phase:location:swipeAim:timestamp:"), view, index, phase, location, aim, timestamp)
	return VZTouchFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZTouch/initWithView:index:phase:location:swipeAim:timestamp:
func (v VZTouch) InitWithViewIndexPhaseLocationSwipeAimTimestamp(view objectivec.IObject, index byte, phase int64, location corefoundation.CGPoint, aim int64, timestamp float64) VZTouch {
	rv := objc.Send[VZTouch](v.ID, objc.Sel("initWithView:index:phase:location:swipeAim:timestamp:"), view, index, phase, location, aim, timestamp)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZTouch/location
func (v VZTouch) Location() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](v.ID, objc.Sel("location"))
	return corefoundation.CGPoint(rv)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZTouch/phase
func (v VZTouch) Phase() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("phase"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZTouch/swipeAim
func (v VZTouch) SwipeAim() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("swipeAim"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZTouch/timestamp
func (v VZTouch) Timestamp() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("timestamp"))
	return rv
}

