// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLWindowMirroringContext] class.
var (
	_SLWindowMirroringContextClass     SLWindowMirroringContextClass
	_SLWindowMirroringContextClassOnce sync.Once
)

func getSLWindowMirroringContextClass() SLWindowMirroringContextClass {
	_SLWindowMirroringContextClassOnce.Do(func() {
		_SLWindowMirroringContextClass = SLWindowMirroringContextClass{class: objc.GetClass("SLWindowMirroringContext")}
	})
	return _SLWindowMirroringContextClass
}

// GetSLWindowMirroringContextClass returns the class object for SLWindowMirroringContext.
func GetSLWindowMirroringContextClass() SLWindowMirroringContextClass {
	return getSLWindowMirroringContextClass()
}

type SLWindowMirroringContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLWindowMirroringContextClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLWindowMirroringContextClass) Alloc() SLWindowMirroringContext {
	rv := objc.Send[SLWindowMirroringContext](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLWindowMirroringContext.CurrentSession]
//   - [SLWindowMirroringContext.CurrentShieldWindowID]
//   - [SLWindowMirroringContext.DisplayID]
//   - [SLWindowMirroringContext.Extend]
//   - [SLWindowMirroringContext.Filter]
//   - [SLWindowMirroringContext.MirrorTo]
//   - [SLWindowMirroringContext.MirrorToWithFilterShowCursor]
//   - [SLWindowMirroringContext.ResetSession]
//   - [SLWindowMirroringContext.SetupSession]
//   - [SLWindowMirroringContext.ShieldWindowID]
//   - [SLWindowMirroringContext.SetShieldWindowID]
//   - [SLWindowMirroringContext.SrcDisplayID]
//   - [SLWindowMirroringContext.InitWithDisplay]
//
// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringContext
type SLWindowMirroringContext struct {
	objectivec.Object
}

// SLWindowMirroringContextFromID constructs a [SLWindowMirroringContext] from an objc.ID.
func SLWindowMirroringContextFromID(id objc.ID) SLWindowMirroringContext {
	return SLWindowMirroringContext{objectivec.Object{ID: id}}
}

// Ensure SLWindowMirroringContext implements ISLWindowMirroringContext.
var _ ISLWindowMirroringContext = SLWindowMirroringContext{}

// An interface definition for the [SLWindowMirroringContext] class.
//
// # Methods
//
//   - [ISLWindowMirroringContext.CurrentSession]
//   - [ISLWindowMirroringContext.CurrentShieldWindowID]
//   - [ISLWindowMirroringContext.DisplayID]
//   - [ISLWindowMirroringContext.Extend]
//   - [ISLWindowMirroringContext.Filter]
//   - [ISLWindowMirroringContext.MirrorTo]
//   - [ISLWindowMirroringContext.MirrorToWithFilterShowCursor]
//   - [ISLWindowMirroringContext.ResetSession]
//   - [ISLWindowMirroringContext.SetupSession]
//   - [ISLWindowMirroringContext.ShieldWindowID]
//   - [ISLWindowMirroringContext.SetShieldWindowID]
//   - [ISLWindowMirroringContext.SrcDisplayID]
//   - [ISLWindowMirroringContext.InitWithDisplay]
//
// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringContext
type ISLWindowMirroringContext interface {
	objectivec.IObject

	// Topic: Methods

	CurrentSession()
	CurrentShieldWindowID() objectivec.IObject
	DisplayID() foundation.NSNumber
	Extend() bool
	Filter() objectivec.IObject
	MirrorTo(to objectivec.IObject) bool
	MirrorToWithFilterShowCursor(to objectivec.IObject, filter objectivec.IObject, cursor bool) bool
	ResetSession()
	SetupSession()
	ShieldWindowID() foundation.NSNumber
	SetShieldWindowID(value foundation.NSNumber)
	SrcDisplayID() foundation.NSNumber
	InitWithDisplay(display objectivec.IObject) SLWindowMirroringContext
}

// Init initializes the instance.
func (s SLWindowMirroringContext) Init() SLWindowMirroringContext {
	rv := objc.Send[SLWindowMirroringContext](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLWindowMirroringContext) Autorelease() SLWindowMirroringContext {
	rv := objc.Send[SLWindowMirroringContext](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLWindowMirroringContext creates a new SLWindowMirroringContext instance.
func NewSLWindowMirroringContext() SLWindowMirroringContext {
	class := getSLWindowMirroringContextClass()
	rv := objc.Send[SLWindowMirroringContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringContext/initWithDisplay:
func NewSLWindowMirroringContextWithDisplay(display objectivec.IObject) SLWindowMirroringContext {
	instance := getSLWindowMirroringContextClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisplay:"), display)
	return SLWindowMirroringContextFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringContext/currentSession
func (s SLWindowMirroringContext) CurrentSession() {
	objc.Send[objc.ID](s.ID, objc.Sel("currentSession"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringContext/currentShieldWindowID
func (s SLWindowMirroringContext) CurrentShieldWindowID() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("currentShieldWindowID"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringContext/extend
func (s SLWindowMirroringContext) Extend() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("extend"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringContext/filter
func (s SLWindowMirroringContext) Filter() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("filter"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringContext/mirrorTo:
func (s SLWindowMirroringContext) MirrorTo(to objectivec.IObject) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("mirrorTo:"), to)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringContext/mirrorTo:withFilter:showCursor:
func (s SLWindowMirroringContext) MirrorToWithFilterShowCursor(to objectivec.IObject, filter objectivec.IObject, cursor bool) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("mirrorTo:withFilter:showCursor:"), to, filter, cursor)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringContext/resetSession
func (s SLWindowMirroringContext) ResetSession() {
	objc.Send[objc.ID](s.ID, objc.Sel("resetSession"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringContext/setupSession
func (s SLWindowMirroringContext) SetupSession() {
	objc.Send[objc.ID](s.ID, objc.Sel("setupSession"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringContext/initWithDisplay:
func (s SLWindowMirroringContext) InitWithDisplay(display objectivec.IObject) SLWindowMirroringContext {
	rv := objc.Send[SLWindowMirroringContext](s.ID, objc.Sel("initWithDisplay:"), display)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringContext/displayID
func (s SLWindowMirroringContext) DisplayID() foundation.NSNumber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayID"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringContext/shieldWindowID
func (s SLWindowMirroringContext) ShieldWindowID() foundation.NSNumber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("shieldWindowID"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (s SLWindowMirroringContext) SetShieldWindowID(value foundation.NSNumber) {
	objc.Send[struct{}](s.ID, objc.Sel("setShieldWindowID:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLWindowMirroringContext/srcDisplayID
func (s SLWindowMirroringContext) SrcDisplayID() foundation.NSNumber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("srcDisplayID"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
