// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtualMachineView] class.
var (
	_VZVirtualMachineViewClass     VZVirtualMachineViewClass
	_VZVirtualMachineViewClassOnce sync.Once
)

func getVZVirtualMachineViewClass() VZVirtualMachineViewClass {
	_VZVirtualMachineViewClassOnce.Do(func() {
		_VZVirtualMachineViewClass = VZVirtualMachineViewClass{class: objc.GetClass("VZVirtualMachineView")}
	})
	return _VZVirtualMachineViewClass
}

// GetVZVirtualMachineViewClass returns the class object for VZVirtualMachineView.
func GetVZVirtualMachineViewClass() VZVirtualMachineViewClass {
	return getVZVirtualMachineViewClass()
}

type VZVirtualMachineViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtualMachineViewClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtualMachineViewClass) Alloc() VZVirtualMachineView {
	rv := objc.Send[VZVirtualMachineView](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZVirtualMachineView._canGrabMouseInput]
//   - [VZVirtualMachineView._canReleaseMouseInput]
//   - [VZVirtualMachineView._delegate]
//   - [VZVirtualMachineView.Set_delegate]
//   - [VZVirtualMachineView._grabMouseInput]
//   - [VZVirtualMachineView._graphicsDisplay]
//   - [VZVirtualMachineView.Set_graphicsDisplay]
//   - [VZVirtualMachineView._releaseMouseInput]
//   - [VZVirtualMachineView._scaleMode]
//   - [VZVirtualMachineView.Set_scaleMode]
//   - [VZVirtualMachineView._setDelegate]
//   - [VZVirtualMachineView._setGraphicsDisplay]
//   - [VZVirtualMachineView._setScaleMode]
//   - [VZVirtualMachineView.AcceptsFirstResponder]
//   - [VZVirtualMachineView.BecomeFirstResponder]
//   - [VZVirtualMachineView.DisplayDidBeginReconfiguration]
//   - [VZVirtualMachineView.DisplayDidEndReconfiguration]
//   - [VZVirtualMachineView.FlagsChanged]
//   - [VZVirtualMachineView.KeyDown]
//   - [VZVirtualMachineView.KeyUp]
//   - [VZVirtualMachineView.Layout]
//   - [VZVirtualMachineView.MagnifyWithEvent]
//   - [VZVirtualMachineView.MouseDown]
//   - [VZVirtualMachineView.MouseDragged]
//   - [VZVirtualMachineView.MouseEntered]
//   - [VZVirtualMachineView.MouseExited]
//   - [VZVirtualMachineView.MouseMoved]
//   - [VZVirtualMachineView.MouseUp]
//   - [VZVirtualMachineView.OtherMouseDown]
//   - [VZVirtualMachineView.OtherMouseDragged]
//   - [VZVirtualMachineView.OtherMouseUp]
//   - [VZVirtualMachineView.QuickLookWithEvent]
//   - [VZVirtualMachineView.ResignFirstResponder]
//   - [VZVirtualMachineView.RightMouseDown]
//   - [VZVirtualMachineView.RightMouseDragged]
//   - [VZVirtualMachineView.RightMouseUp]
//   - [VZVirtualMachineView.RotateWithEvent]
//   - [VZVirtualMachineView.ScrollWheel]
//   - [VZVirtualMachineView.SmartMagnifyWithEvent]
//   - [VZVirtualMachineView.UpdateTrackingAreas]
//   - [VZVirtualMachineView.ViewDidEndLiveResize]
//   - [VZVirtualMachineView.ViewDidMoveToWindow]
//   - [VZVirtualMachineView.ViewWillMoveToWindow]
//   - [VZVirtualMachineView.ViewWillStartLiveResize]
//   - [VZVirtualMachineView.DebugDescription]
//   - [VZVirtualMachineView.Description]
//   - [VZVirtualMachineView.Hash]
//   - [VZVirtualMachineView.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView
type VZVirtualMachineView struct {
	objectivec.Object
}

// VZVirtualMachineViewFromID constructs a [VZVirtualMachineView] from an objc.ID.
func VZVirtualMachineViewFromID(id objc.ID) VZVirtualMachineView {
	return VZVirtualMachineView{Object: objectivec.Object{ID: id}}
}
// Ensure VZVirtualMachineView implements IVZVirtualMachineView.
var _ IVZVirtualMachineView = VZVirtualMachineView{}

// An interface definition for the [VZVirtualMachineView] class.
//
// # Methods
//
//   - [IVZVirtualMachineView._canGrabMouseInput]
//   - [IVZVirtualMachineView._canReleaseMouseInput]
//   - [IVZVirtualMachineView._delegate]
//   - [IVZVirtualMachineView.Set_delegate]
//   - [IVZVirtualMachineView._grabMouseInput]
//   - [IVZVirtualMachineView._graphicsDisplay]
//   - [IVZVirtualMachineView.Set_graphicsDisplay]
//   - [IVZVirtualMachineView._releaseMouseInput]
//   - [IVZVirtualMachineView._scaleMode]
//   - [IVZVirtualMachineView.Set_scaleMode]
//   - [IVZVirtualMachineView._setDelegate]
//   - [IVZVirtualMachineView._setGraphicsDisplay]
//   - [IVZVirtualMachineView._setScaleMode]
//   - [IVZVirtualMachineView.AcceptsFirstResponder]
//   - [IVZVirtualMachineView.BecomeFirstResponder]
//   - [IVZVirtualMachineView.DisplayDidBeginReconfiguration]
//   - [IVZVirtualMachineView.DisplayDidEndReconfiguration]
//   - [IVZVirtualMachineView.FlagsChanged]
//   - [IVZVirtualMachineView.KeyDown]
//   - [IVZVirtualMachineView.KeyUp]
//   - [IVZVirtualMachineView.Layout]
//   - [IVZVirtualMachineView.MagnifyWithEvent]
//   - [IVZVirtualMachineView.MouseDown]
//   - [IVZVirtualMachineView.MouseDragged]
//   - [IVZVirtualMachineView.MouseEntered]
//   - [IVZVirtualMachineView.MouseExited]
//   - [IVZVirtualMachineView.MouseMoved]
//   - [IVZVirtualMachineView.MouseUp]
//   - [IVZVirtualMachineView.OtherMouseDown]
//   - [IVZVirtualMachineView.OtherMouseDragged]
//   - [IVZVirtualMachineView.OtherMouseUp]
//   - [IVZVirtualMachineView.QuickLookWithEvent]
//   - [IVZVirtualMachineView.ResignFirstResponder]
//   - [IVZVirtualMachineView.RightMouseDown]
//   - [IVZVirtualMachineView.RightMouseDragged]
//   - [IVZVirtualMachineView.RightMouseUp]
//   - [IVZVirtualMachineView.RotateWithEvent]
//   - [IVZVirtualMachineView.ScrollWheel]
//   - [IVZVirtualMachineView.SmartMagnifyWithEvent]
//   - [IVZVirtualMachineView.UpdateTrackingAreas]
//   - [IVZVirtualMachineView.ViewDidEndLiveResize]
//   - [IVZVirtualMachineView.ViewDidMoveToWindow]
//   - [IVZVirtualMachineView.ViewWillMoveToWindow]
//   - [IVZVirtualMachineView.ViewWillStartLiveResize]
//   - [IVZVirtualMachineView.DebugDescription]
//   - [IVZVirtualMachineView.Description]
//   - [IVZVirtualMachineView.Hash]
//   - [IVZVirtualMachineView.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView
type IVZVirtualMachineView interface {
	INSView

	// Topic: Methods

	_canGrabMouseInput() bool
	_canReleaseMouseInput() bool
	_delegate() objectivec.IObject
	Set_delegate(value objectivec.IObject)
	_grabMouseInput() bool
	_graphicsDisplay() IVZGraphicsDisplay
	Set_graphicsDisplay(value IVZGraphicsDisplay)
	_releaseMouseInput() bool
	_scaleMode() int64
	Set_scaleMode(value int64)
	_setDelegate(delegate objectivec.IObject)
	_setGraphicsDisplay(display objectivec.IObject)
	_setScaleMode(mode int64)
	AcceptsFirstResponder() bool
	BecomeFirstResponder() bool
	DisplayDidBeginReconfiguration(reconfiguration objectivec.IObject)
	DisplayDidEndReconfiguration(reconfiguration objectivec.IObject)
	FlagsChanged(changed objectivec.IObject)
	KeyDown(down objectivec.IObject)
	KeyUp(up objectivec.IObject)
	Layout()
	MagnifyWithEvent(event objectivec.IObject)
	MouseDown(down objectivec.IObject)
	MouseDragged(dragged objectivec.IObject)
	MouseEntered(entered objectivec.IObject)
	MouseExited(exited objectivec.IObject)
	MouseMoved(moved objectivec.IObject)
	MouseUp(up objectivec.IObject)
	OtherMouseDown(down objectivec.IObject)
	OtherMouseDragged(dragged objectivec.IObject)
	OtherMouseUp(up objectivec.IObject)
	QuickLookWithEvent(event objectivec.IObject)
	ResignFirstResponder() bool
	RightMouseDown(down objectivec.IObject)
	RightMouseDragged(dragged objectivec.IObject)
	RightMouseUp(up objectivec.IObject)
	RotateWithEvent(event objectivec.IObject)
	ScrollWheel(wheel objectivec.IObject)
	SmartMagnifyWithEvent(event objectivec.IObject)
	UpdateTrackingAreas()
	ViewDidEndLiveResize()
	ViewDidMoveToWindow()
	ViewWillMoveToWindow(window objectivec.IObject)
	ViewWillStartLiveResize()
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZVirtualMachineView) Init() VZVirtualMachineView {
	rv := objc.Send[VZVirtualMachineView](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtualMachineView) Autorelease() VZVirtualMachineView {
	rv := objc.Send[VZVirtualMachineView](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachineView creates a new VZVirtualMachineView instance.
func NewVZVirtualMachineView() VZVirtualMachineView {
	class := getVZVirtualMachineViewClass()
	rv := objc.Send[VZVirtualMachineView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/initWithCoder:
func NewVirtualMachineViewWithCoder(coder objectivec.IObject) VZVirtualMachineView {
	instance := getVZVirtualMachineViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return VZVirtualMachineViewFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/_grabMouseInput
func (v VZVirtualMachineView) _grabMouseInput() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_grabMouseInput"))
	return rv
}

// GrabMouseInput is an exported wrapper for the private method _grabMouseInput.
func (v VZVirtualMachineView) GrabMouseInput() bool {
	return v._grabMouseInput()
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/_releaseMouseInput
func (v VZVirtualMachineView) _releaseMouseInput() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_releaseMouseInput"))
	return rv
}

// ReleaseMouseInput is an exported wrapper for the private method _releaseMouseInput.
func (v VZVirtualMachineView) ReleaseMouseInput() bool {
	return v._releaseMouseInput()
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/_setDelegate:
func (v VZVirtualMachineView) _setDelegate(delegate objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setDelegate:"), delegate)
}

// SetDelegate is an exported wrapper for the private method _setDelegate.
func (v VZVirtualMachineView) SetDelegate(delegate objectivec.IObject) {
	v._setDelegate(delegate)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/_setGraphicsDisplay:
func (v VZVirtualMachineView) _setGraphicsDisplay(display objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setGraphicsDisplay:"), display)
}

// SetGraphicsDisplay is an exported wrapper for the private method _setGraphicsDisplay.
func (v VZVirtualMachineView) SetGraphicsDisplay(display objectivec.IObject) {
	v._setGraphicsDisplay(display)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/_setScaleMode:
func (v VZVirtualMachineView) _setScaleMode(mode int64) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setScaleMode:"), mode)
}

// SetScaleMode is an exported wrapper for the private method _setScaleMode.
func (v VZVirtualMachineView) SetScaleMode(mode int64) {
	v._setScaleMode(mode)
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/acceptsFirstResponder
func (v VZVirtualMachineView) AcceptsFirstResponder() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("acceptsFirstResponder"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/becomeFirstResponder
func (v VZVirtualMachineView) BecomeFirstResponder() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("becomeFirstResponder"))
	return rv
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/displayDidBeginReconfiguration:
func (v VZVirtualMachineView) DisplayDidBeginReconfiguration(reconfiguration objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("displayDidBeginReconfiguration:"), reconfiguration)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/displayDidEndReconfiguration:
func (v VZVirtualMachineView) DisplayDidEndReconfiguration(reconfiguration objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("displayDidEndReconfiguration:"), reconfiguration)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/flagsChanged:
func (v VZVirtualMachineView) FlagsChanged(changed objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("flagsChanged:"), changed)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/keyDown:
func (v VZVirtualMachineView) KeyDown(down objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("keyDown:"), down)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/keyUp:
func (v VZVirtualMachineView) KeyUp(up objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("keyUp:"), up)
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/layout
func (v VZVirtualMachineView) Layout() {
	objc.Send[objc.ID](v.ID, objc.Sel("layout"))
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/magnifyWithEvent:
func (v VZVirtualMachineView) MagnifyWithEvent(event objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("magnifyWithEvent:"), event)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/mouseDown:
func (v VZVirtualMachineView) MouseDown(down objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("mouseDown:"), down)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/mouseDragged:
func (v VZVirtualMachineView) MouseDragged(dragged objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("mouseDragged:"), dragged)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/mouseEntered:
func (v VZVirtualMachineView) MouseEntered(entered objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("mouseEntered:"), entered)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/mouseExited:
func (v VZVirtualMachineView) MouseExited(exited objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("mouseExited:"), exited)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/mouseMoved:
func (v VZVirtualMachineView) MouseMoved(moved objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("mouseMoved:"), moved)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/mouseUp:
func (v VZVirtualMachineView) MouseUp(up objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("mouseUp:"), up)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/otherMouseDown:
func (v VZVirtualMachineView) OtherMouseDown(down objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("otherMouseDown:"), down)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/otherMouseDragged:
func (v VZVirtualMachineView) OtherMouseDragged(dragged objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("otherMouseDragged:"), dragged)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/otherMouseUp:
func (v VZVirtualMachineView) OtherMouseUp(up objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("otherMouseUp:"), up)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/quickLookWithEvent:
func (v VZVirtualMachineView) QuickLookWithEvent(event objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("quickLookWithEvent:"), event)
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/resignFirstResponder
func (v VZVirtualMachineView) ResignFirstResponder() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("resignFirstResponder"))
	return rv
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/rightMouseDown:
func (v VZVirtualMachineView) RightMouseDown(down objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("rightMouseDown:"), down)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/rightMouseDragged:
func (v VZVirtualMachineView) RightMouseDragged(dragged objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("rightMouseDragged:"), dragged)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/rightMouseUp:
func (v VZVirtualMachineView) RightMouseUp(up objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("rightMouseUp:"), up)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/rotateWithEvent:
func (v VZVirtualMachineView) RotateWithEvent(event objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("rotateWithEvent:"), event)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/scrollWheel:
func (v VZVirtualMachineView) ScrollWheel(wheel objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("scrollWheel:"), wheel)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/smartMagnifyWithEvent:
func (v VZVirtualMachineView) SmartMagnifyWithEvent(event objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("smartMagnifyWithEvent:"), event)
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/updateTrackingAreas
func (v VZVirtualMachineView) UpdateTrackingAreas() {
	objc.Send[objc.ID](v.ID, objc.Sel("updateTrackingAreas"))
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/viewDidEndLiveResize
func (v VZVirtualMachineView) ViewDidEndLiveResize() {
	objc.Send[objc.ID](v.ID, objc.Sel("viewDidEndLiveResize"))
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/viewDidMoveToWindow
func (v VZVirtualMachineView) ViewDidMoveToWindow() {
	objc.Send[objc.ID](v.ID, objc.Sel("viewDidMoveToWindow"))
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/viewWillMoveToWindow:
func (v VZVirtualMachineView) ViewWillMoveToWindow(window objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("viewWillMoveToWindow:"), window)
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/viewWillStartLiveResize
func (v VZVirtualMachineView) ViewWillStartLiveResize() {
	objc.Send[objc.ID](v.ID, objc.Sel("viewWillStartLiveResize"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/_canGrabMouseInput
func (v VZVirtualMachineView) _canGrabMouseInput() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_canGrabMouseInput"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/_canReleaseMouseInput
func (v VZVirtualMachineView) _canReleaseMouseInput() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_canReleaseMouseInput"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/_delegate
func (v VZVirtualMachineView) _delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_delegate"))
	return objectivec.Object{ID: rv}
}
func (v VZVirtualMachineView) Set_delegate(value objectivec.IObject) {
	objc.Send[struct{}](v.ID, objc.Sel("set_delegate:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/_graphicsDisplay
func (v VZVirtualMachineView) _graphicsDisplay() IVZGraphicsDisplay {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_graphicsDisplay"))
	return VZGraphicsDisplayFromID(objc.ID(rv))
}
func (v VZVirtualMachineView) Set_graphicsDisplay(value IVZGraphicsDisplay) {
	objc.Send[struct{}](v.ID, objc.Sel("set_graphicsDisplay:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/_scaleMode
func (v VZVirtualMachineView) _scaleMode() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("_scaleMode"))
	return rv
}
func (v VZVirtualMachineView) Set_scaleMode(value int64) {
	objc.Send[struct{}](v.ID, objc.Sel("set_scaleMode:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/debugDescription
func (v VZVirtualMachineView) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/description
func (v VZVirtualMachineView) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/hash
func (v VZVirtualMachineView) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/superclass
func (v VZVirtualMachineView) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}

