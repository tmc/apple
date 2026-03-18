// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSScrollEdgeEffectStyle] class.
var (
	_NSScrollEdgeEffectStyleClass     NSScrollEdgeEffectStyleClass
	_NSScrollEdgeEffectStyleClassOnce sync.Once
)

func getNSScrollEdgeEffectStyleClass() NSScrollEdgeEffectStyleClass {
	_NSScrollEdgeEffectStyleClassOnce.Do(func() {
		_NSScrollEdgeEffectStyleClass = NSScrollEdgeEffectStyleClass{class: objc.GetClass("NSScrollEdgeEffectStyle")}
	})
	return _NSScrollEdgeEffectStyleClass
}

// GetNSScrollEdgeEffectStyleClass returns the class object for NSScrollEdgeEffectStyle.
func GetNSScrollEdgeEffectStyleClass() NSScrollEdgeEffectStyleClass {
	return getNSScrollEdgeEffectStyleClass()
}

type NSScrollEdgeEffectStyleClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScrollEdgeEffectStyleClass) Alloc() NSScrollEdgeEffectStyle {
	rv := objc.Send[NSScrollEdgeEffectStyle](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Styles for a scroll view’s edge effect.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollEdgeEffectStyle
type NSScrollEdgeEffectStyle struct {
	objectivec.Object
}

// NSScrollEdgeEffectStyleFromID constructs a [NSScrollEdgeEffectStyle] from an objc.ID.
//
// Styles for a scroll view’s edge effect.
func NSScrollEdgeEffectStyleFromID(id objc.ID) NSScrollEdgeEffectStyle {
	return NSScrollEdgeEffectStyle{objectivec.Object{ID: id}}
}
// NOTE: NSScrollEdgeEffectStyle adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSScrollEdgeEffectStyle] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollEdgeEffectStyle
type INSScrollEdgeEffectStyle interface {
	objectivec.IObject

	// The titlebar accessory’s preferred effect for content scrolling behind it.
	PreferredScrollEdgeEffectStyle() INSScrollEdgeEffectStyle
	SetPreferredScrollEdgeEffectStyle(value INSScrollEdgeEffectStyle)
}

// Init initializes the instance.
func (s NSScrollEdgeEffectStyle) Init() NSScrollEdgeEffectStyle {
	rv := objc.Send[NSScrollEdgeEffectStyle](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScrollEdgeEffectStyle) Autorelease() NSScrollEdgeEffectStyle {
	rv := objc.Send[NSScrollEdgeEffectStyle](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScrollEdgeEffectStyle creates a new NSScrollEdgeEffectStyle instance.
func NewNSScrollEdgeEffectStyle() NSScrollEdgeEffectStyle {
	class := getNSScrollEdgeEffectStyleClass()
	rv := objc.Send[NSScrollEdgeEffectStyle](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The titlebar accessory’s preferred effect for content scrolling behind
// it.
//
// See: https://developer.apple.com/documentation/appkit/nstitlebaraccessoryviewcontroller/preferredscrolledgeeffectstyle
func (s NSScrollEdgeEffectStyle) PreferredScrollEdgeEffectStyle() INSScrollEdgeEffectStyle {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("preferredScrollEdgeEffectStyle"))
	return NSScrollEdgeEffectStyleFromID(objc.ID(rv))
}
func (s NSScrollEdgeEffectStyle) SetPreferredScrollEdgeEffectStyle(value INSScrollEdgeEffectStyle) {
	objc.Send[struct{}](s.ID, objc.Sel("setPreferredScrollEdgeEffectStyle:"), value)
}

// The automatic scroll edge effect style.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollEdgeEffectStyle/automatic
func (_NSScrollEdgeEffectStyleClass NSScrollEdgeEffectStyleClass) AutomaticStyle() NSScrollEdgeEffectStyle {
	rv := objc.Send[objc.ID](objc.ID(_NSScrollEdgeEffectStyleClass.class), objc.Sel("automaticStyle"))
	return NSScrollEdgeEffectStyleFromID(objc.ID(rv))
}

// A scroll edge effect with a hard cutoff.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollEdgeEffectStyle/hard
func (_NSScrollEdgeEffectStyleClass NSScrollEdgeEffectStyleClass) HardStyle() NSScrollEdgeEffectStyle {
	rv := objc.Send[objc.ID](objc.ID(_NSScrollEdgeEffectStyleClass.class), objc.Sel("hardStyle"))
	return NSScrollEdgeEffectStyleFromID(objc.ID(rv))
}

// A scroll edge effect with a soft edge.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrollEdgeEffectStyle/soft
func (_NSScrollEdgeEffectStyleClass NSScrollEdgeEffectStyleClass) SoftStyle() NSScrollEdgeEffectStyle {
	rv := objc.Send[objc.ID](objc.ID(_NSScrollEdgeEffectStyleClass.class), objc.Sel("softStyle"))
	return NSScrollEdgeEffectStyleFromID(objc.ID(rv))
}

