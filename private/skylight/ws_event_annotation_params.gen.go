// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WSEventAnnotationParams] class.
var (
	_WSEventAnnotationParamsClass     WSEventAnnotationParamsClass
	_WSEventAnnotationParamsClassOnce sync.Once
)

func getWSEventAnnotationParamsClass() WSEventAnnotationParamsClass {
	_WSEventAnnotationParamsClassOnce.Do(func() {
		_WSEventAnnotationParamsClass = WSEventAnnotationParamsClass{class: objc.GetClass("WSEventAnnotationParams")}
	})
	return _WSEventAnnotationParamsClass
}

// GetWSEventAnnotationParamsClass returns the class object for WSEventAnnotationParams.
func GetWSEventAnnotationParamsClass() WSEventAnnotationParamsClass {
	return getWSEventAnnotationParamsClass()
}

type WSEventAnnotationParamsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WSEventAnnotationParamsClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WSEventAnnotationParamsClass) Alloc() WSEventAnnotationParams {
	rv := objc.Send[WSEventAnnotationParams](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [WSEventAnnotationParams.ForceRouting]
//   - [WSEventAnnotationParams.SetForceRouting]
//   - [WSEventAnnotationParams.NoBackground]
//   - [WSEventAnnotationParams.SetNoBackground]
//   - [WSEventAnnotationParams.SetTarget]
//   - [WSEventAnnotationParams.SetSetTarget]
//   - [WSEventAnnotationParams.ShouldProcessEvent]
//   - [WSEventAnnotationParams.SetShouldProcessEvent]
//
// See: https://developer.apple.com/documentation/SkyLight/WSEventAnnotationParams
type WSEventAnnotationParams struct {
	objectivec.Object
}

// WSEventAnnotationParamsFromID constructs a [WSEventAnnotationParams] from an objc.ID.
func WSEventAnnotationParamsFromID(id objc.ID) WSEventAnnotationParams {
	return WSEventAnnotationParams{objectivec.Object{ID: id}}
}

// Ensure WSEventAnnotationParams implements IWSEventAnnotationParams.
var _ IWSEventAnnotationParams = WSEventAnnotationParams{}

// An interface definition for the [WSEventAnnotationParams] class.
//
// # Methods
//
//   - [IWSEventAnnotationParams.ForceRouting]
//   - [IWSEventAnnotationParams.SetForceRouting]
//   - [IWSEventAnnotationParams.NoBackground]
//   - [IWSEventAnnotationParams.SetNoBackground]
//   - [IWSEventAnnotationParams.SetTarget]
//   - [IWSEventAnnotationParams.SetSetTarget]
//   - [IWSEventAnnotationParams.ShouldProcessEvent]
//   - [IWSEventAnnotationParams.SetShouldProcessEvent]
//
// See: https://developer.apple.com/documentation/SkyLight/WSEventAnnotationParams
type IWSEventAnnotationParams interface {
	objectivec.IObject

	// Topic: Methods

	ForceRouting() bool
	SetForceRouting(value bool)
	NoBackground() bool
	SetNoBackground(value bool)
	SetTarget() bool
	SetSetTarget(value bool)
	ShouldProcessEvent() bool
	SetShouldProcessEvent(value bool)
}

// Init initializes the instance.
func (w WSEventAnnotationParams) Init() WSEventAnnotationParams {
	rv := objc.Send[WSEventAnnotationParams](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WSEventAnnotationParams) Autorelease() WSEventAnnotationParams {
	rv := objc.Send[WSEventAnnotationParams](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWSEventAnnotationParams creates a new WSEventAnnotationParams instance.
func NewWSEventAnnotationParams() WSEventAnnotationParams {
	class := getWSEventAnnotationParamsClass()
	rv := objc.Send[WSEventAnnotationParams](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventAnnotationParams/forceRouting
func (w WSEventAnnotationParams) ForceRouting() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("forceRouting"))
	return rv
}
func (w WSEventAnnotationParams) SetForceRouting(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setForceRouting:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventAnnotationParams/noBackground
func (w WSEventAnnotationParams) NoBackground() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("noBackground"))
	return rv
}
func (w WSEventAnnotationParams) SetNoBackground(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setNoBackground:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventAnnotationParams/setTarget
func (w WSEventAnnotationParams) SetTarget() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("setTarget"))
	return rv
}
func (w WSEventAnnotationParams) SetSetTarget(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setSetTarget:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/WSEventAnnotationParams/shouldProcessEvent
func (w WSEventAnnotationParams) ShouldProcessEvent() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("shouldProcessEvent"))
	return rv
}
func (w WSEventAnnotationParams) SetShouldProcessEvent(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setShouldProcessEvent:"), value)
}
