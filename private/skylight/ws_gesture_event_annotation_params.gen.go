// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
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
	rv := objc.SendIfResponds[WSGestureEventAnnotationParams](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSGestureEventAnnotationParams.GestureStreamState]
//   - [WSGestureEventAnnotationParams.SetGestureStreamState]
//   - [WSGestureEventAnnotationParams.HitTestTarget]
//   - [WSGestureEventAnnotationParams.SetHitTestTarget]
//   - [WSGestureEventAnnotationParams.PassthroughTargets]
//   - [WSGestureEventAnnotationParams.SetPassthroughTargets]
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
//   - [IWSGestureEventAnnotationParams.HitTestTarget]
//   - [IWSGestureEventAnnotationParams.SetHitTestTarget]
//   - [IWSGestureEventAnnotationParams.PassthroughTargets]
//   - [IWSGestureEventAnnotationParams.SetPassthroughTargets]
type IWSGestureEventAnnotationParams interface {
	IWSEventAnnotationParams

	// Topic: Methods

	GestureStreamState() int32
	SetGestureStreamState(value int32)
	HitTestTarget() unsafe.Pointer
	SetHitTestTarget(value unsafe.Pointer)
	PassthroughTargets() foundation.INSArray
	SetPassthroughTargets(value foundation.INSArray)
}

// Init initializes the instance.
func (w WSGestureEventAnnotationParams) Init() WSGestureEventAnnotationParams {
	rv := objc.SendIfResponds[WSGestureEventAnnotationParams](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSGestureEventAnnotationParams) Autorelease() WSGestureEventAnnotationParams {
	rv := objc.SendIfResponds[WSGestureEventAnnotationParams](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSGestureEventAnnotationParams creates a new WSGestureEventAnnotationParams instance.
func NewWSGestureEventAnnotationParams() WSGestureEventAnnotationParams {
	class := getWSGestureEventAnnotationParamsClass()
	rv := objc.SendIfResponds[WSGestureEventAnnotationParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (w WSGestureEventAnnotationParams) GestureStreamState() int32 {
	rv := objc.SendIfResponds[int32](w.ID, objc.Sel("gestureStreamState"))
	return rv
}
func (w WSGestureEventAnnotationParams) SetGestureStreamState(value int32) {
	objc.SendIfResponds[struct{}](w.ID, objc.Sel("setGestureStreamState:"), value)
}
func (w WSGestureEventAnnotationParams) HitTestTarget() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](w.ID, objc.Sel("hitTestTarget"))
	return rv
}
func (w WSGestureEventAnnotationParams) SetHitTestTarget(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](w.ID, objc.Sel("setHitTestTarget:"), value)
}
func (w WSGestureEventAnnotationParams) PassthroughTargets() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](w.ID, objc.Sel("passthroughTargets"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (w WSGestureEventAnnotationParams) SetPassthroughTargets(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](w.ID, objc.Sel("setPassthroughTargets:"), value)
}
