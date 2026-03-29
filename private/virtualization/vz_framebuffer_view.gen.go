// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
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
	rv := objc.Send[VZFramebufferView](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZFramebufferView.ActionForLayerForKey]
//   - [VZFramebufferView.Cursor]
//   - [VZFramebufferView.SetCursor]
//   - [VZFramebufferView.DisplayProtectionOptions]
//   - [VZFramebufferView.Framebuffer]
//   - [VZFramebufferView.SetFramebuffer]
//   - [VZFramebufferView.FramebufferDidUpdateCursor]
//   - [VZFramebufferView.FramebufferDidUpdateFrame]
//   - [VZFramebufferView.FramebufferDidUpdateGraphicsOrientation]
//   - [VZFramebufferView.FramebufferDidUpdateColorSpace]
//   - [VZFramebufferView.FramebufferSize]
//   - [VZFramebufferView.GetDisplayProtectionOptions]
//   - [VZFramebufferView.ShowsCursor]
//   - [VZFramebufferView.SetShowsCursor]
//   - [VZFramebufferView.SuppressFrameUpdates]
//   - [VZFramebufferView.SetSuppressFrameUpdates]
//   - [VZFramebufferView.DebugDescription]
//   - [VZFramebufferView.Description]
//   - [VZFramebufferView.Hash]
//   - [VZFramebufferView.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferView
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
//   - [IVZFramebufferView.Framebuffer]
//   - [IVZFramebufferView.SetFramebuffer]
//   - [IVZFramebufferView.FramebufferDidUpdateCursor]
//   - [IVZFramebufferView.FramebufferDidUpdateFrame]
//   - [IVZFramebufferView.FramebufferDidUpdateGraphicsOrientation]
//   - [IVZFramebufferView.FramebufferDidUpdateColorSpace]
//   - [IVZFramebufferView.FramebufferSize]
//   - [IVZFramebufferView.GetDisplayProtectionOptions]
//   - [IVZFramebufferView.ShowsCursor]
//   - [IVZFramebufferView.SetShowsCursor]
//   - [IVZFramebufferView.SuppressFrameUpdates]
//   - [IVZFramebufferView.SetSuppressFrameUpdates]
//   - [IVZFramebufferView.DebugDescription]
//   - [IVZFramebufferView.Description]
//   - [IVZFramebufferView.Hash]
//   - [IVZFramebufferView.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferView
type IVZFramebufferView interface {
	appkit.INSView

	// Topic: Methods

	ActionForLayerForKey(layer objectivec.IObject, key objectivec.IObject) objectivec.IObject
	Cursor() appkit.NSCursor
	SetCursor(value appkit.NSCursor)
	DisplayProtectionOptions() foundation.NSNumber
	Framebuffer() *VZFramebuffer
	SetFramebuffer(value *VZFramebuffer)
	FramebufferDidUpdateCursor(framebuffer objectivec.IObject, cursor objectivec.IObject)
	FramebufferDidUpdateFrame(framebuffer objectivec.IObject, frame objectivec.IObject)
	FramebufferDidUpdateGraphicsOrientation(framebuffer objectivec.IObject, orientation int64)
	FramebufferDidUpdateColorSpace(space objectivec.IObject)
	FramebufferSize() corefoundation.CGSize
	GetDisplayProtectionOptions() objectivec.IObject
	ShowsCursor() bool
	SetShowsCursor(value bool)
	SuppressFrameUpdates() bool
	SetSuppressFrameUpdates(value bool)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZFramebufferView) Init() VZFramebufferView {
	rv := objc.Send[VZFramebufferView](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZFramebufferView) Autorelease() VZFramebufferView {
	rv := objc.Send[VZFramebufferView](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZFramebufferView creates a new VZFramebufferView instance.
func NewVZFramebufferView() VZFramebufferView {
	class := getVZFramebufferViewClass()
	rv := objc.Send[VZFramebufferView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferView/initWithCoder:
func NewVZFramebufferViewWithCoder(coder objectivec.IObject) VZFramebufferView {
	instance := getVZFramebufferViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return VZFramebufferViewFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferView/initWithFrame:
func NewVZFramebufferViewWithFrame(frame corefoundation.CGRect) VZFramebufferView {
	instance := getVZFramebufferViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frame)
	return VZFramebufferViewFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferView/actionForLayer:forKey:
func (v VZFramebufferView) ActionForLayerForKey(layer objectivec.IObject, key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("actionForLayer:forKey:"), layer, key)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferView/framebuffer:didUpdateCursor:
func (v VZFramebufferView) FramebufferDidUpdateCursor(framebuffer objectivec.IObject, cursor objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("framebuffer:didUpdateCursor:"), framebuffer, cursor)
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferView/framebuffer:didUpdateFrame:
func (v VZFramebufferView) FramebufferDidUpdateFrame(framebuffer objectivec.IObject, frame objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("framebuffer:didUpdateFrame:"), framebuffer, frame)
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferView/framebuffer:didUpdateGraphicsOrientation:
func (v VZFramebufferView) FramebufferDidUpdateGraphicsOrientation(framebuffer objectivec.IObject, orientation int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("framebuffer:didUpdateGraphicsOrientation:"), framebuffer, orientation)
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferView/framebufferDidUpdateColorSpace:
func (v VZFramebufferView) FramebufferDidUpdateColorSpace(space objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("framebufferDidUpdateColorSpace:"), space)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferView/getDisplayProtectionOptions
func (v VZFramebufferView) GetDisplayProtectionOptions() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("getDisplayProtectionOptions"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferView/cursor
func (v VZFramebufferView) Cursor() appkit.NSCursor {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("cursor"))
	return appkit.NSCursorFromID(objc.ID(rv))
}
func (v VZFramebufferView) SetCursor(value appkit.NSCursor) {
	objc.Send[struct{}](v.ID, objc.Sel("setCursor:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferView/debugDescription
func (v VZFramebufferView) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferView/description
func (v VZFramebufferView) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferView/displayProtectionOptions
func (v VZFramebufferView) DisplayProtectionOptions() foundation.NSNumber {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("displayProtectionOptions"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferView/framebuffer
func (v VZFramebufferView) Framebuffer() *VZFramebuffer {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("framebuffer"))
	if rv == 0 {
		return nil
	}
	val := VZFramebufferFromID(objc.ID(rv))
	return &val
}
func (v VZFramebufferView) SetFramebuffer(value *VZFramebuffer) {
	if value == nil {
		objc.Send[struct{}](v.ID, objc.Sel("setFramebuffer:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](v.ID, objc.Sel("setFramebuffer:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferView/framebufferSize
func (v VZFramebufferView) FramebufferSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("framebufferSize"))
	return corefoundation.CGSize(rv)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferView/hash
func (v VZFramebufferView) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferView/showsCursor
func (v VZFramebufferView) ShowsCursor() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("showsCursor"))
	return rv
}
func (v VZFramebufferView) SetShowsCursor(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setShowsCursor:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferView/superclass
func (v VZFramebufferView) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZFramebufferView/suppressFrameUpdates
func (v VZFramebufferView) SuppressFrameUpdates() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("suppressFrameUpdates"))
	return rv
}
func (v VZFramebufferView) SetSuppressFrameUpdates(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setSuppressFrameUpdates:"), value)
}

