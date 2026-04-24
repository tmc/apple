// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSWindowManagerDragContext] class.
var (
	_SLSWindowManagerDragContextClass     SLSWindowManagerDragContextClass
	_SLSWindowManagerDragContextClassOnce sync.Once
)

func getSLSWindowManagerDragContextClass() SLSWindowManagerDragContextClass {
	_SLSWindowManagerDragContextClassOnce.Do(func() {
		_SLSWindowManagerDragContextClass = SLSWindowManagerDragContextClass{class: objc.GetClass("SLSWindowManagerDragContext")}
	})
	return _SLSWindowManagerDragContextClass
}

// GetSLSWindowManagerDragContextClass returns the class object for SLSWindowManagerDragContext.
func GetSLSWindowManagerDragContextClass() SLSWindowManagerDragContextClass {
	return getSLSWindowManagerDragContextClass()
}

type SLSWindowManagerDragContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSWindowManagerDragContextClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSWindowManagerDragContextClass) Alloc() SLSWindowManagerDragContext {
	rv := objc.Send[SLSWindowManagerDragContext](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSWindowManagerDragContext.AccessibilityState]
//   - [SLSWindowManagerDragContext.SetAccessibilityState]
//   - [SLSWindowManagerDragContext.DefaultWindowOrigin]
//   - [SLSWindowManagerDragContext.DragOffset]
//   - [SLSWindowManagerDragContext.SetDragOffset]
//   - [SLSWindowManagerDragContext.MouseLocation]
//   - [SLSWindowManagerDragContext.SetMouseLocation]
//   - [SLSWindowManagerDragContext.ProposedWindowOrigin]
//   - [SLSWindowManagerDragContext.SetProposedWindowOrigin]
//   - [SLSWindowManagerDragContext.Timestamp]
//   - [SLSWindowManagerDragContext.SetTimestamp]
//   - [SLSWindowManagerDragContext.WindowID]
//   - [SLSWindowManagerDragContext.SetWindowID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerDragContext
type SLSWindowManagerDragContext struct {
	objectivec.Object
}

// SLSWindowManagerDragContextFromID constructs a [SLSWindowManagerDragContext] from an objc.ID.
func SLSWindowManagerDragContextFromID(id objc.ID) SLSWindowManagerDragContext {
	return SLSWindowManagerDragContext{objectivec.Object{ID: id}}
}

// Ensure SLSWindowManagerDragContext implements ISLSWindowManagerDragContext.
var _ ISLSWindowManagerDragContext = SLSWindowManagerDragContext{}

// An interface definition for the [SLSWindowManagerDragContext] class.
//
// # Methods
//
//   - [ISLSWindowManagerDragContext.AccessibilityState]
//   - [ISLSWindowManagerDragContext.SetAccessibilityState]
//   - [ISLSWindowManagerDragContext.DefaultWindowOrigin]
//   - [ISLSWindowManagerDragContext.DragOffset]
//   - [ISLSWindowManagerDragContext.SetDragOffset]
//   - [ISLSWindowManagerDragContext.MouseLocation]
//   - [ISLSWindowManagerDragContext.SetMouseLocation]
//   - [ISLSWindowManagerDragContext.ProposedWindowOrigin]
//   - [ISLSWindowManagerDragContext.SetProposedWindowOrigin]
//   - [ISLSWindowManagerDragContext.Timestamp]
//   - [ISLSWindowManagerDragContext.SetTimestamp]
//   - [ISLSWindowManagerDragContext.WindowID]
//   - [ISLSWindowManagerDragContext.SetWindowID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerDragContext
type ISLSWindowManagerDragContext interface {
	objectivec.IObject

	// Topic: Methods

	AccessibilityState() uint64
	SetAccessibilityState(value uint64)
	DefaultWindowOrigin() corefoundation.CGPoint
	DragOffset() corefoundation.CGPoint
	SetDragOffset(value corefoundation.CGPoint)
	MouseLocation() corefoundation.CGPoint
	SetMouseLocation(value corefoundation.CGPoint)
	ProposedWindowOrigin() corefoundation.CGPoint
	SetProposedWindowOrigin(value corefoundation.CGPoint)
	Timestamp() uint64
	SetTimestamp(value uint64)
	WindowID() uint32
	SetWindowID(value uint32)
}

// Init initializes the instance.
func (s SLSWindowManagerDragContext) Init() SLSWindowManagerDragContext {
	rv := objc.Send[SLSWindowManagerDragContext](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSWindowManagerDragContext) Autorelease() SLSWindowManagerDragContext {
	rv := objc.Send[SLSWindowManagerDragContext](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSWindowManagerDragContext creates a new SLSWindowManagerDragContext instance.
func NewSLSWindowManagerDragContext() SLSWindowManagerDragContext {
	class := getSLSWindowManagerDragContextClass()
	rv := objc.Send[SLSWindowManagerDragContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerDragContext/accessibilityState
func (s SLSWindowManagerDragContext) AccessibilityState() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("accessibilityState"))
	return rv
}
func (s SLSWindowManagerDragContext) SetAccessibilityState(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setAccessibilityState:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerDragContext/defaultWindowOrigin
func (s SLSWindowManagerDragContext) DefaultWindowOrigin() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](s.ID, objc.Sel("defaultWindowOrigin"))
	return corefoundation.CGPoint(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerDragContext/dragOffset
func (s SLSWindowManagerDragContext) DragOffset() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](s.ID, objc.Sel("dragOffset"))
	return corefoundation.CGPoint(rv)
}
func (s SLSWindowManagerDragContext) SetDragOffset(value corefoundation.CGPoint) {
	objc.Send[struct{}](s.ID, objc.Sel("setDragOffset:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerDragContext/mouseLocation
func (s SLSWindowManagerDragContext) MouseLocation() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](s.ID, objc.Sel("mouseLocation"))
	return corefoundation.CGPoint(rv)
}
func (s SLSWindowManagerDragContext) SetMouseLocation(value corefoundation.CGPoint) {
	objc.Send[struct{}](s.ID, objc.Sel("setMouseLocation:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerDragContext/proposedWindowOrigin
func (s SLSWindowManagerDragContext) ProposedWindowOrigin() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](s.ID, objc.Sel("proposedWindowOrigin"))
	return corefoundation.CGPoint(rv)
}
func (s SLSWindowManagerDragContext) SetProposedWindowOrigin(value corefoundation.CGPoint) {
	objc.Send[struct{}](s.ID, objc.Sel("setProposedWindowOrigin:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerDragContext/timestamp
func (s SLSWindowManagerDragContext) Timestamp() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("timestamp"))
	return rv
}
func (s SLSWindowManagerDragContext) SetTimestamp(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setTimestamp:"), value)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWindowManagerDragContext/windowID
func (s SLSWindowManagerDragContext) WindowID() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("windowID"))
	return rv
}
func (s SLSWindowManagerDragContext) SetWindowID(value uint32) {
	objc.Send[struct{}](s.ID, objc.Sel("setWindowID:"), value)
}
