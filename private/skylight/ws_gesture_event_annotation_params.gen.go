// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [WSGestureEventAnnotationParams] class.
var (
	_WSGestureEventAnnotationParamsClass     WSGestureEventAnnotationParamsClass
	_WSGestureEventAnnotationParamsClassOnce sync.Once
)

func getWSGestureEventAnnotationParamsClass() WSGestureEventAnnotationParamsClass {
	_WSGestureEventAnnotationParamsClassOnce.Do(func() {
		_WSGestureEventAnnotationParamsClass = WSGestureEventAnnotationParamsClass{class: objc.GetClass("WSGestureEventAnnotationParams")}
	})
	return _WSGestureEventAnnotationParamsClass
}

// GetWSGestureEventAnnotationParamsClass returns the class object for WSGestureEventAnnotationParams.
func GetWSGestureEventAnnotationParamsClass() WSGestureEventAnnotationParamsClass {
	return getWSGestureEventAnnotationParamsClass()
}

type WSGestureEventAnnotationParamsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSGestureEventAnnotationParamsClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSGestureEventAnnotationParamsClass) Alloc() WSGestureEventAnnotationParams {
	rv := objc.Send[WSGestureEventAnnotationParams](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSGestureEventAnnotationParams.GestureStreamState]
//   - [WSGestureEventAnnotationParams.SetGestureStreamState]
//
// See: https://developer.apple.com/documentation/SkyLight/WSGestureEventAnnotationParams
type WSGestureEventAnnotationParams struct {
	WSEventAnnotationParams
}

// WSGestureEventAnnotationParamsFromID constructs a [WSGestureEventAnnotationParams] from an objc.ID.
func WSGestureEventAnnotationParamsFromID(id objc.ID) WSGestureEventAnnotationParams {
	return WSGestureEventAnnotationParams{WSEventAnnotationParams: WSEventAnnotationParamsFromID(id)}
}

// Ensure WSGestureEventAnnotationParams implements IWSGestureEventAnnotationParams.
var _ IWSGestureEventAnnotationParams = WSGestureEventAnnotationParams{}

// An interface definition for the [WSGestureEventAnnotationParams] class.
//
// # Methods
//
//   - [IWSGestureEventAnnotationParams.GestureStreamState]
//   - [IWSGestureEventAnnotationParams.SetGestureStreamState]
//
// See: https://developer.apple.com/documentation/SkyLight/WSGestureEventAnnotationParams
type IWSGestureEventAnnotationParams interface {
	IWSEventAnnotationParams

	// Topic: Methods

	GestureStreamState() int
	SetGestureStreamState(value int)
}

// Init initializes the instance.
func (w WSGestureEventAnnotationParams) Init() WSGestureEventAnnotationParams {
	rv := objc.Send[WSGestureEventAnnotationParams](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSGestureEventAnnotationParams) Autorelease() WSGestureEventAnnotationParams {
	rv := objc.Send[WSGestureEventAnnotationParams](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSGestureEventAnnotationParams creates a new WSGestureEventAnnotationParams instance.
func NewWSGestureEventAnnotationParams() WSGestureEventAnnotationParams {
	class := getWSGestureEventAnnotationParamsClass()
	rv := objc.Send[WSGestureEventAnnotationParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSGestureEventAnnotationParams/gestureStreamState
func (w WSGestureEventAnnotationParams) GestureStreamState() int {
	rv := objc.Send[int](w.ID, objc.Sel("gestureStreamState"))
	return rv
}
func (w WSGestureEventAnnotationParams) SetGestureStreamState(value int) {
	objc.Send[struct{}](w.ID, objc.Sel("setGestureStreamState:"), value)
}
