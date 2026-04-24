// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSWMBridgedWindow] class.
var (
	_SLSWMBridgedWindowClass     SLSWMBridgedWindowClass
	_SLSWMBridgedWindowClassOnce sync.Once
)

func getSLSWMBridgedWindowClass() SLSWMBridgedWindowClass {
	_SLSWMBridgedWindowClassOnce.Do(func() {
		_SLSWMBridgedWindowClass = SLSWMBridgedWindowClass{class: objc.GetClass("SLSWMBridgedWindow")}
	})
	return _SLSWMBridgedWindowClass
}

// GetSLSWMBridgedWindowClass returns the class object for SLSWMBridgedWindow.
func GetSLSWMBridgedWindowClass() SLSWMBridgedWindowClass {
	return getSLSWMBridgedWindowClass()
}

type SLSWMBridgedWindowClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSWMBridgedWindowClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSWMBridgedWindowClass) Alloc() SLSWMBridgedWindow {
	rv := objc.Send[SLSWMBridgedWindow](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSWMBridgedWindow._rebuildChildWindowInfos]
//   - [SLSWMBridgedWindow.AddChildWindowOrdered]
//   - [SLSWMBridgedWindow.ClearOrderingGroup]
//   - [SLSWMBridgedWindow.OrderWindowRelativeToIDRelativeToOrderGroup]
//   - [SLSWMBridgedWindow.RemoveChildWindow]
//   - [SLSWMBridgedWindow.RemoveFromParent]
//   - [SLSWMBridgedWindow.SetFrameForceAsync]
//   - [SLSWMBridgedWindow.SetWindowLevel]
//   - [SLSWMBridgedWindow.WindowDidUpdateWithChangedProperties]
//   - [SLSWMBridgedWindow.InitWithWindowID]
//   - [SLSWMBridgedWindow.DebugDescription]
//   - [SLSWMBridgedWindow.Description]
//   - [SLSWMBridgedWindow.Hash]
//   - [SLSWMBridgedWindow.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindow
type SLSWMBridgedWindow struct {
	objectivec.Object
}

// SLSWMBridgedWindowFromID constructs a [SLSWMBridgedWindow] from an objc.ID.
func SLSWMBridgedWindowFromID(id objc.ID) SLSWMBridgedWindow {
	return SLSWMBridgedWindow{objectivec.Object{ID: id}}
}

// Ensure SLSWMBridgedWindow implements ISLSWMBridgedWindow.
var _ ISLSWMBridgedWindow = SLSWMBridgedWindow{}

// An interface definition for the [SLSWMBridgedWindow] class.
//
// # Methods
//
//   - [ISLSWMBridgedWindow._rebuildChildWindowInfos]
//   - [ISLSWMBridgedWindow.AddChildWindowOrdered]
//   - [ISLSWMBridgedWindow.ClearOrderingGroup]
//   - [ISLSWMBridgedWindow.OrderWindowRelativeToIDRelativeToOrderGroup]
//   - [ISLSWMBridgedWindow.RemoveChildWindow]
//   - [ISLSWMBridgedWindow.RemoveFromParent]
//   - [ISLSWMBridgedWindow.SetFrameForceAsync]
//   - [ISLSWMBridgedWindow.SetWindowLevel]
//   - [ISLSWMBridgedWindow.WindowDidUpdateWithChangedProperties]
//   - [ISLSWMBridgedWindow.InitWithWindowID]
//   - [ISLSWMBridgedWindow.DebugDescription]
//   - [ISLSWMBridgedWindow.Description]
//   - [ISLSWMBridgedWindow.Hash]
//   - [ISLSWMBridgedWindow.Superclass]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindow
type ISLSWMBridgedWindow interface {
	objectivec.IObject

	// Topic: Methods

	_rebuildChildWindowInfos()
	AddChildWindowOrdered(window objectivec.IObject, ordered int)
	ClearOrderingGroup()
	OrderWindowRelativeToIDRelativeToOrderGroup(window int, id uint32, to objectivec.IObject, group bool)
	RemoveChildWindow(window objectivec.IObject)
	RemoveFromParent()
	SetFrameForceAsync(frame corefoundation.CGRect, async bool)
	SetWindowLevel(level int)
	WindowDidUpdateWithChangedProperties(window objectivec.IObject, properties uint64)
	InitWithWindowID(id uint32) SLSWMBridgedWindow
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (s SLSWMBridgedWindow) Init() SLSWMBridgedWindow {
	rv := objc.Send[SLSWMBridgedWindow](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSWMBridgedWindow) Autorelease() SLSWMBridgedWindow {
	rv := objc.Send[SLSWMBridgedWindow](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSWMBridgedWindow creates a new SLSWMBridgedWindow instance.
func NewSLSWMBridgedWindow() SLSWMBridgedWindow {
	class := getSLSWMBridgedWindowClass()
	rv := objc.Send[SLSWMBridgedWindow](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindow/initWithWindowID:
func NewSLSWMBridgedWindowWithWindowID(id uint32) SLSWMBridgedWindow {
	instance := getSLSWMBridgedWindowClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWindowID:"), id)
	return SLSWMBridgedWindowFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindow/_rebuildChildWindowInfos
func (s SLSWMBridgedWindow) _rebuildChildWindowInfos() {
	objc.Send[objc.ID](s.ID, objc.Sel("_rebuildChildWindowInfos"))
}

// RebuildChildWindowInfos is an exported wrapper for the private method _rebuildChildWindowInfos.
func (s SLSWMBridgedWindow) RebuildChildWindowInfos() {
	s._rebuildChildWindowInfos()
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindow/addChildWindow:ordered:
func (s SLSWMBridgedWindow) AddChildWindowOrdered(window objectivec.IObject, ordered int) {
	objc.Send[objc.ID](s.ID, objc.Sel("addChildWindow:ordered:"), window, ordered)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindow/clearOrderingGroup
func (s SLSWMBridgedWindow) ClearOrderingGroup() {
	objc.Send[objc.ID](s.ID, objc.Sel("clearOrderingGroup"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindow/orderWindow:relativeToID:relativeTo:orderGroup:
func (s SLSWMBridgedWindow) OrderWindowRelativeToIDRelativeToOrderGroup(window int, id uint32, to objectivec.IObject, group bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("orderWindow:relativeToID:relativeTo:orderGroup:"), window, id, to, group)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindow/removeChildWindow:
func (s SLSWMBridgedWindow) RemoveChildWindow(window objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("removeChildWindow:"), window)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindow/removeFromParent
func (s SLSWMBridgedWindow) RemoveFromParent() {
	objc.Send[objc.ID](s.ID, objc.Sel("removeFromParent"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindow/setFrame:forceAsync:
func (s SLSWMBridgedWindow) SetFrameForceAsync(frame corefoundation.CGRect, async bool) {
	objc.Send[objc.ID](s.ID, objc.Sel("setFrame:forceAsync:"), frame, async)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindow/setWindowLevel:
func (s SLSWMBridgedWindow) SetWindowLevel(level int) {
	objc.Send[objc.ID](s.ID, objc.Sel("setWindowLevel:"), level)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindow/window:didUpdateWithChangedProperties:
func (s SLSWMBridgedWindow) WindowDidUpdateWithChangedProperties(window objectivec.IObject, properties uint64) {
	objc.Send[objc.ID](s.ID, objc.Sel("window:didUpdateWithChangedProperties:"), window, properties)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindow/initWithWindowID:
func (s SLSWMBridgedWindow) InitWithWindowID(id uint32) SLSWMBridgedWindow {
	rv := objc.Send[SLSWMBridgedWindow](s.ID, objc.Sel("initWithWindowID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindow/debugDescription
func (s SLSWMBridgedWindow) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindow/description
func (s SLSWMBridgedWindow) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindow/hash
func (s SLSWMBridgedWindow) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSWMBridgedWindow/superclass
func (s SLSWMBridgedWindow) Superclass() objc.Class {
	rv := objc.Send[objc.Class](s.ID, objc.Sel("superclass"))
	return rv
}
