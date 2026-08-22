// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[VZVirtualMachineView](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

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
//   - [VZVirtualMachineView.DisplayDidBeginReconfiguration]
//   - [VZVirtualMachineView.DisplayDidEndReconfiguration]
//   - [VZVirtualMachineView.DebugDescription]
//   - [VZVirtualMachineView.Description]
//   - [VZVirtualMachineView.Hash]
//   - [VZVirtualMachineView.Superclass]
type VZVirtualMachineView struct {
	appkit.NSView
}

// VZVirtualMachineViewFromID constructs a [VZVirtualMachineView] from an objc.ID.
func VZVirtualMachineViewFromID(id objc.ID) VZVirtualMachineView {
	return VZVirtualMachineView{NSView: appkit.NSViewFromID(id)}
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
//   - [IVZVirtualMachineView.DisplayDidBeginReconfiguration]
//   - [IVZVirtualMachineView.DisplayDidEndReconfiguration]
//   - [IVZVirtualMachineView.DebugDescription]
//   - [IVZVirtualMachineView.Description]
//   - [IVZVirtualMachineView.Hash]
//   - [IVZVirtualMachineView.Superclass]
type IVZVirtualMachineView interface {
	appkit.INSView

	// Topic: Methods

	_canGrabMouseInput() bool
	_canReleaseMouseInput() bool
	_delegate() unsafe.Pointer
	Set_delegate(value unsafe.Pointer)
	_grabMouseInput() bool
	_graphicsDisplay() IVZGraphicsDisplay
	Set_graphicsDisplay(value IVZGraphicsDisplay)
	_releaseMouseInput() bool
	_scaleMode() int64
	Set_scaleMode(value int64)
	_setDelegate(delegate objectivec.IObject)
	_setGraphicsDisplay(display objectivec.IObject)
	_setScaleMode(mode int64)
	DisplayDidBeginReconfiguration(reconfiguration objectivec.IObject)
	DisplayDidEndReconfiguration(reconfiguration objectivec.IObject)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZVirtualMachineView) Init() VZVirtualMachineView {
	rv := objc.SendIfResponds[VZVirtualMachineView](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtualMachineView) Autorelease() VZVirtualMachineView {
	rv := objc.SendIfResponds[VZVirtualMachineView](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachineView creates a new VZVirtualMachineView instance.
func NewVZVirtualMachineView() VZVirtualMachineView {
	class := getVZVirtualMachineViewClass()
	rv := objc.SendIfResponds[VZVirtualMachineView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVirtualMachineViewWithCoder(coder objectivec.IObject) VZVirtualMachineView {
	instance := getVZVirtualMachineViewClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return VZVirtualMachineViewFromID(rv)
}

func (v VZVirtualMachineView) _grabMouseInput() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("_grabMouseInput"))
	return rv
}

// GrabMouseInput is an exported wrapper for the private method _grabMouseInput.
func (v VZVirtualMachineView) GrabMouseInput() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_grabMouseInput")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_grabMouseInput"}
		return false, err
	}
	return v._grabMouseInput(), nil
}

// CanGrabMouseInput reports whether the receiver responds to the private selector _grabMouseInput.
func (v VZVirtualMachineView) CanGrabMouseInput() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_grabMouseInput"))
}
func (v VZVirtualMachineView) _releaseMouseInput() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("_releaseMouseInput"))
	return rv
}

// ReleaseMouseInput is an exported wrapper for the private method _releaseMouseInput.
func (v VZVirtualMachineView) ReleaseMouseInput() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_releaseMouseInput")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_releaseMouseInput"}
		return false, err
	}
	return v._releaseMouseInput(), nil
}

// CanReleaseMouseInput reports whether the receiver responds to the private selector _releaseMouseInput.
func (v VZVirtualMachineView) CanReleaseMouseInput() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_releaseMouseInput"))
}
func (v VZVirtualMachineView) _setDelegate(delegate objectivec.IObject) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_setDelegate:"), delegate)
}

// SetDelegate is an exported wrapper for the private method _setDelegate.
func (v VZVirtualMachineView) SetDelegate(delegate objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setDelegate:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setDelegate:"}
		return err
	}
	v._setDelegate(delegate)
	return nil
}

// CanSetDelegate reports whether the receiver responds to the private selector _setDelegate:.
func (v VZVirtualMachineView) CanSetDelegate() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setDelegate:"))
}
func (v VZVirtualMachineView) _setGraphicsDisplay(display objectivec.IObject) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_setGraphicsDisplay:"), display)
}

// SetGraphicsDisplay is an exported wrapper for the private method _setGraphicsDisplay.
func (v VZVirtualMachineView) SetGraphicsDisplay(display objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setGraphicsDisplay:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setGraphicsDisplay:"}
		return err
	}
	v._setGraphicsDisplay(display)
	return nil
}

// CanSetGraphicsDisplay reports whether the receiver responds to the private selector _setGraphicsDisplay:.
func (v VZVirtualMachineView) CanSetGraphicsDisplay() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setGraphicsDisplay:"))
}
func (v VZVirtualMachineView) _setScaleMode(mode int64) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_setScaleMode:"), mode)
}

// SetScaleMode is an exported wrapper for the private method _setScaleMode.
func (v VZVirtualMachineView) SetScaleMode(mode int64) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setScaleMode:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setScaleMode:"}
		return err
	}
	v._setScaleMode(mode)
	return nil
}

// CanSetScaleMode reports whether the receiver responds to the private selector _setScaleMode:.
func (v VZVirtualMachineView) CanSetScaleMode() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setScaleMode:"))
}
func (v VZVirtualMachineView) DisplayDidBeginReconfiguration(reconfiguration objectivec.IObject) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("displayDidBeginReconfiguration:"), reconfiguration)
}
func (v VZVirtualMachineView) DisplayDidEndReconfiguration(reconfiguration objectivec.IObject) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("displayDidEndReconfiguration:"), reconfiguration)
}

func (v VZVirtualMachineView) _canGrabMouseInput() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("_canGrabMouseInput"))
	return rv
}
func (v VZVirtualMachineView) _canReleaseMouseInput() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("_canReleaseMouseInput"))
	return rv
}
func (v VZVirtualMachineView) _delegate() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](v.ID, objc.Sel("_delegate"))
	return rv
}

// CanDelegate reports whether the receiver responds to the private selector _delegate.
func (v VZVirtualMachineView) CanDelegate() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_delegate"))
}

// Delegate is an exported wrapper for the private property _delegate.
func (v VZVirtualMachineView) Delegate() (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_delegate")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_delegate"}
	}
	return v._delegate(), nil
}
func (v VZVirtualMachineView) Set_delegate(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("set_delegate:"), value)
}
func (v VZVirtualMachineView) _graphicsDisplay() IVZGraphicsDisplay {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_graphicsDisplay"))
	return VZGraphicsDisplayFromID(objc.ID(rv))
}

// CanGraphicsDisplay reports whether the receiver responds to the private selector _graphicsDisplay.
func (v VZVirtualMachineView) CanGraphicsDisplay() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_graphicsDisplay"))
}

// GraphicsDisplay is an exported wrapper for the private property _graphicsDisplay.
func (v VZVirtualMachineView) GraphicsDisplay() (IVZGraphicsDisplay, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_graphicsDisplay")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_graphicsDisplay"}
	}
	return v._graphicsDisplay(), nil
}
func (v VZVirtualMachineView) Set_graphicsDisplay(value IVZGraphicsDisplay) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("set_graphicsDisplay:"), value)
}
func (v VZVirtualMachineView) _scaleMode() int64 {
	rv := objc.SendIfResponds[int64](v.ID, objc.Sel("_scaleMode"))
	return rv
}

// CanScaleMode reports whether the receiver responds to the private selector _scaleMode.
func (v VZVirtualMachineView) CanScaleMode() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_scaleMode"))
}

// ScaleMode is an exported wrapper for the private property _scaleMode.
func (v VZVirtualMachineView) ScaleMode() (int64, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_scaleMode")) {
		return 0, &objc.UnrecognizedSelectorError{Selector: "_scaleMode"}
	}
	return v._scaleMode(), nil
}
func (v VZVirtualMachineView) Set_scaleMode(value int64) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("set_scaleMode:"), value)
}
func (v VZVirtualMachineView) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZVirtualMachineView) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZVirtualMachineView) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZVirtualMachineView) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
