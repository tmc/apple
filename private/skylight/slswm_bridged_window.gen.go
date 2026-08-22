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
	rv := objc.SendIfResponds[SLSWMBridgedWindow](objc.ID(sc.class), objc.Sel("alloc"))
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
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (s SLSWMBridgedWindow) Init() SLSWMBridgedWindow {
	rv := objc.SendIfResponds[SLSWMBridgedWindow](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSWMBridgedWindow) Autorelease() SLSWMBridgedWindow {
	rv := objc.SendIfResponds[SLSWMBridgedWindow](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSWMBridgedWindow creates a new SLSWMBridgedWindow instance.
func NewSLSWMBridgedWindow() SLSWMBridgedWindow {
	class := getSLSWMBridgedWindowClass()
	rv := objc.SendIfResponds[SLSWMBridgedWindow](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSWMBridgedWindowWithWindowID(id uint32) SLSWMBridgedWindow {
	instance := getSLSWMBridgedWindowClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithWindowID:"), id)
	return SLSWMBridgedWindowFromID(rv)
}

func (s SLSWMBridgedWindow) _rebuildChildWindowInfos() {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("_rebuildChildWindowInfos"))
}

// RebuildChildWindowInfos is an exported wrapper for the private method _rebuildChildWindowInfos.
func (s SLSWMBridgedWindow) RebuildChildWindowInfos() error {
	if !objc.RespondsToSelector(s.ID, objc.Sel("_rebuildChildWindowInfos")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_rebuildChildWindowInfos"}
		return err
	}
	s._rebuildChildWindowInfos()
	return nil
}

// CanRebuildChildWindowInfos reports whether the receiver responds to the private selector _rebuildChildWindowInfos.
func (s SLSWMBridgedWindow) CanRebuildChildWindowInfos() bool {
	return objc.RespondsToSelector(s.ID, objc.Sel("_rebuildChildWindowInfos"))
}
func (s SLSWMBridgedWindow) AddChildWindowOrdered(window objectivec.IObject, ordered int) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("addChildWindow:ordered:"), window, ordered)
}
func (s SLSWMBridgedWindow) ClearOrderingGroup() {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("clearOrderingGroup"))
}
func (s SLSWMBridgedWindow) OrderWindowRelativeToIDRelativeToOrderGroup(window int, id uint32, to objectivec.IObject, group bool) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("orderWindow:relativeToID:relativeTo:orderGroup:"), window, id, to, group)
}
func (s SLSWMBridgedWindow) RemoveChildWindow(window objectivec.IObject) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("removeChildWindow:"), window)
}
func (s SLSWMBridgedWindow) RemoveFromParent() {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("removeFromParent"))
}
func (s SLSWMBridgedWindow) SetFrameForceAsync(frame corefoundation.CGRect, async bool) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setFrame:forceAsync:"), frame, async)
}
func (s SLSWMBridgedWindow) SetWindowLevel(level int) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("setWindowLevel:"), level)
}
func (s SLSWMBridgedWindow) WindowDidUpdateWithChangedProperties(window objectivec.IObject, properties uint64) {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("window:didUpdateWithChangedProperties:"), window, properties)
}
func (s SLSWMBridgedWindow) InitWithWindowID(id uint32) SLSWMBridgedWindow {
	rv := objc.SendIfResponds[SLSWMBridgedWindow](s.ID, objc.Sel("initWithWindowID:"), id)
	return rv
}

func (s SLSWMBridgedWindow) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (s SLSWMBridgedWindow) Description() string {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (s SLSWMBridgedWindow) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("hash"))
	return rv
}
func (s SLSWMBridgedWindow) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](s.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
