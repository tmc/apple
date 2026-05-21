// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [GCDeviceCursor] class.
var (
	_GCDeviceCursorClass     GCDeviceCursorClass
	_GCDeviceCursorClassOnce sync.Once
)

func getGCDeviceCursorClass() GCDeviceCursorClass {
	_GCDeviceCursorClassOnce.Do(func() {
		_GCDeviceCursorClass = GCDeviceCursorClass{class: objc.GetClass("GCDeviceCursor")}
	})
	return _GCDeviceCursorClass
}

// GetGCDeviceCursorClass returns the class object for GCDeviceCursor.
func GetGCDeviceCursorClass() GCDeviceCursorClass {
	return getGCDeviceCursorClass()
}

type GCDeviceCursorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCDeviceCursorClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCDeviceCursorClass) Alloc() GCDeviceCursor {
	rv := objc.Send[GCDeviceCursor](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// A control element for the cursor used as a directional pad.
//
// # Overview
//
// This controller element allows you to use the cursor as a directional pad
// with the values of the [GCControllerDirectionPad.XAxis] and
// [GCControllerDirectionPad.YAxis] elements scaled to the width and height of
// the screen, not ranging from `-1` to `1`.
//
// See: https://developer.apple.com/documentation/GameController/GCDeviceCursor
type GCDeviceCursor struct {
	GCControllerDirectionPad
}

// GCDeviceCursorFromID constructs a [GCDeviceCursor] from an objc.ID.
//
// A control element for the cursor used as a directional pad.
func GCDeviceCursorFromID(id objc.ID) GCDeviceCursor {
	return GCDeviceCursor{GCControllerDirectionPad: GCControllerDirectionPadFromID(id)}
}

// NOTE: GCDeviceCursor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCDeviceCursor] class.
//
// See: https://developer.apple.com/documentation/GameController/GCDeviceCursor
type IGCDeviceCursor interface {
	IGCControllerDirectionPad
}

// Init initializes the instance.
func (g GCDeviceCursor) Init() GCDeviceCursor {
	rv := objc.Send[GCDeviceCursor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCDeviceCursor) Autorelease() GCDeviceCursor {
	rv := objc.Send[GCDeviceCursor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCDeviceCursor creates a new GCDeviceCursor instance.
func NewGCDeviceCursor() GCDeviceCursor {
	class := getGCDeviceCursorClass()
	rv := objc.Send[GCDeviceCursor](objc.ID(class.class), objc.Sel("new"))
	return rv
}
