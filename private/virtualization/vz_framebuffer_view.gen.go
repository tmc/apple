// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZFramebufferView] class.
var (
	_VZFramebufferViewClass     VZFramebufferViewClass
	_VZFramebufferViewClassOnce sync.Once
)

func getVZFramebufferViewClass() VZFramebufferViewClass {
	_VZFramebufferViewClassOnce.Do(func() {
		_VZFramebufferViewClass = VZFramebufferViewClass{class: objc.GetClass("_VZFramebufferView")}
	})
	return _VZFramebufferViewClass
}

// GetVZFramebufferViewClass returns the class object for _VZFramebufferView.
func GetVZFramebufferViewClass() VZFramebufferViewClass {
	return getVZFramebufferViewClass()
}

type VZFramebufferViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZFramebufferViewClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZFramebufferViewClass) Alloc() VZFramebufferView {
	rv := objc.SendIfResponds[VZFramebufferView](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZFramebufferView.ActionForLayerForKey]
//   - [VZFramebufferView.Cursor]
//   - [VZFramebufferView.SetCursor]
//   - [VZFramebufferView.DisplayProtectionOptions]
//   - [VZFramebufferView.FramebufferSize]
//   - [VZFramebufferView.PresenterDidUpdateContentHeadroom]
//   - [VZFramebufferView.PresenterDidUpdateHostDisplay]
//   - [VZFramebufferView.ShowsCursor]
//   - [VZFramebufferView.SetShowsCursor]
//   - [VZFramebufferView.SuppressFrameUpdates]
//   - [VZFramebufferView.SetSuppressFrameUpdates]
//   - [VZFramebufferView.DebugDescription]
//   - [VZFramebufferView.Description]
//   - [VZFramebufferView.Hash]
//   - [VZFramebufferView.Superclass]
type VZFramebufferView struct {
	appkit.NSView
}

// VZFramebufferViewFromID constructs a [VZFramebufferView] from an objc.ID.
func VZFramebufferViewFromID(id objc.ID) VZFramebufferView {
	return VZFramebufferView{NSView: appkit.NSViewFromID(id)}
}

// Ensure VZFramebufferView implements IVZFramebufferView.
var _ IVZFramebufferView = VZFramebufferView{}

// An interface definition for the [VZFramebufferView] class.
//
// # Methods
//
//   - [IVZFramebufferView.ActionForLayerForKey]
//   - [IVZFramebufferView.Cursor]
//   - [IVZFramebufferView.SetCursor]
//   - [IVZFramebufferView.DisplayProtectionOptions]
//   - [IVZFramebufferView.FramebufferSize]
//   - [IVZFramebufferView.PresenterDidUpdateContentHeadroom]
//   - [IVZFramebufferView.PresenterDidUpdateHostDisplay]
//   - [IVZFramebufferView.ShowsCursor]
//   - [IVZFramebufferView.SetShowsCursor]
//   - [IVZFramebufferView.SuppressFrameUpdates]
//   - [IVZFramebufferView.SetSuppressFrameUpdates]
//   - [IVZFramebufferView.DebugDescription]
//   - [IVZFramebufferView.Description]
//   - [IVZFramebufferView.Hash]
//   - [IVZFramebufferView.Superclass]
type IVZFramebufferView interface {
	appkit.INSView

	// Topic: Methods

	ActionForLayerForKey(layer objectivec.IObject, key objectivec.IObject) objectivec.IObject
	Cursor() appkit.NSCursor
	SetCursor(value appkit.NSCursor)
	DisplayProtectionOptions() foundation.NSNumber
	FramebufferSize() corefoundation.CGSize
	PresenterDidUpdateContentHeadroom(presenter objectivec.IObject, headroom float64)
	PresenterDidUpdateHostDisplay(presenter objectivec.IObject, display HostDisplayUpdate)
	ShowsCursor() bool
	SetShowsCursor(value bool)
	SuppressFrameUpdates() bool
	SetSuppressFrameUpdates(value bool)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZFramebufferView) Init() VZFramebufferView {
	rv := objc.SendIfResponds[VZFramebufferView](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZFramebufferView) Autorelease() VZFramebufferView {
	rv := objc.SendIfResponds[VZFramebufferView](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZFramebufferView creates a new VZFramebufferView instance.
func NewVZFramebufferView() VZFramebufferView {
	class := getVZFramebufferViewClass()
	rv := objc.SendIfResponds[VZFramebufferView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZFramebufferViewWithCoder(coder objectivec.IObject) VZFramebufferView {
	instance := getVZFramebufferViewClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return VZFramebufferViewFromID(rv)
}

func NewVZFramebufferViewWithFrame(frame corefoundation.CGRect) VZFramebufferView {
	instance := getVZFramebufferViewClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frame)
	return VZFramebufferViewFromID(rv)
}

func (v VZFramebufferView) ActionForLayerForKey(layer objectivec.IObject, key objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("actionForLayer:forKey:"), layer, key)
	return objectivec.Object{ID: rv}
}
func (v VZFramebufferView) PresenterDidUpdateContentHeadroom(presenter objectivec.IObject, headroom float64) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("presenter:didUpdateContentHeadroom:"), presenter, headroom)
}
func (v VZFramebufferView) PresenterDidUpdateHostDisplay(presenter objectivec.IObject, display HostDisplayUpdate) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("presenter:didUpdateHostDisplay:"), presenter, display)
}

func (v VZFramebufferView) Cursor() appkit.NSCursor {
	rv := objc.SendIfResponds[appkit.NSCursor](v.ID, objc.Sel("cursor"))
	return appkit.NSCursor(rv)
}
func (v VZFramebufferView) SetCursor(value appkit.NSCursor) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setCursor:"), value)
}
func (v VZFramebufferView) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZFramebufferView) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZFramebufferView) DisplayProtectionOptions() foundation.NSNumber {
	rv := objc.SendIfResponds[foundation.NSNumber](v.ID, objc.Sel("displayProtectionOptions"))
	return foundation.NSNumber(rv)
}
func (v VZFramebufferView) FramebufferSize() corefoundation.CGSize {
	rv := objc.SendIfResponds[corefoundation.CGSize](v.ID, objc.Sel("framebufferSize"))
	return corefoundation.CGSize(rv)
}
func (v VZFramebufferView) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZFramebufferView) ShowsCursor() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("showsCursor"))
	return rv
}
func (v VZFramebufferView) SetShowsCursor(value bool) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setShowsCursor:"), value)
}
func (v VZFramebufferView) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (v VZFramebufferView) SuppressFrameUpdates() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("suppressFrameUpdates"))
	return rv
}
func (v VZFramebufferView) SetSuppressFrameUpdates(value bool) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setSuppressFrameUpdates:"), value)
}
