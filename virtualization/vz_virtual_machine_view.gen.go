// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/foundation"
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

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtualMachineViewClass) Alloc() VZVirtualMachineView {
	rv := objc.Send[VZVirtualMachineView](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// A view that allows user interaction with a VM.
//
// # Overview
// 
// The [VZVirtualMachineView] is a UI element that shows the contents of the
// VM frame buffer that you can optionally configure to respond to changes in
// the host’s display settings. If the VM configuration includes a keyboard
// and a pointing device, the view forwards keyboard and mouse events to the
// VM through those devices.
//
// # Configuring the VM
//
//   - [VZVirtualMachineView.AutomaticallyReconfiguresDisplay]: A Boolean value that indicates whether the graphics display associated with this view automatically reconfigures with respect to view changes.
//   - [VZVirtualMachineView.SetAutomaticallyReconfiguresDisplay]
//   - [VZVirtualMachineView.CapturesSystemKeys]: A Boolean value that determines whether the system should send certain system keyboard shortcuts to the guest instead of the host.
//   - [VZVirtualMachineView.SetCapturesSystemKeys]
//   - [VZVirtualMachineView.VirtualMachine]: The VM to display in the view.
//   - [VZVirtualMachineView.SetVirtualMachine]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView
type VZVirtualMachineView struct {
	appkit.NSView
}

// VZVirtualMachineViewFromID constructs a [VZVirtualMachineView] from an objc.ID.
//
// A view that allows user interaction with a VM.
func VZVirtualMachineViewFromID(id objc.ID) VZVirtualMachineView {
	return VZVirtualMachineView{NSView: appkit.NSViewFromID(id)}
}
// NOTE: VZVirtualMachineView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZVirtualMachineView] class.
//
// # Configuring the VM
//
//   - [IVZVirtualMachineView.AutomaticallyReconfiguresDisplay]: A Boolean value that indicates whether the graphics display associated with this view automatically reconfigures with respect to view changes.
//   - [IVZVirtualMachineView.SetAutomaticallyReconfiguresDisplay]
//   - [IVZVirtualMachineView.CapturesSystemKeys]: A Boolean value that determines whether the system should send certain system keyboard shortcuts to the guest instead of the host.
//   - [IVZVirtualMachineView.SetCapturesSystemKeys]
//   - [IVZVirtualMachineView.VirtualMachine]: The VM to display in the view.
//   - [IVZVirtualMachineView.SetVirtualMachine]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView
type IVZVirtualMachineView interface {
	appkit.INSView

	// Topic: Configuring the VM

	// A Boolean value that indicates whether the graphics display associated with this view automatically reconfigures with respect to view changes.
	AutomaticallyReconfiguresDisplay() bool
	SetAutomaticallyReconfiguresDisplay(value bool)
	// A Boolean value that determines whether the system should send certain system keyboard shortcuts to the guest instead of the host.
	CapturesSystemKeys() bool
	SetCapturesSystemKeys(value bool)
	// The VM to display in the view.
	VirtualMachine() IVZVirtualMachine
	SetVirtualMachine(value IVZVirtualMachine)
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









func (v VZVirtualMachineView) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](v.ID, objc.Sel("encodeWithCoder:"), coder)
}












// A Boolean value that indicates whether the graphics display associated with
// this view automatically reconfigures with respect to view changes.
//
// # Discussion
// 
// Set this property to [true] to automatically resize or reconfigure this
// graphics display when the view properties update. For example, resizing the
// display when the view has a live resize operation. When enabled, the
// graphics display automatically reconfigures to match the host display
// environment.
// 
// You can set this property on only a single [VZVirtualMachineView] targeting
// a particular [VZGraphicsDisplay] at a time. If multiple
// [VZVirtualMachineView] views targeting the same VZGraphicsDisplay enable
// this property, only one view respects the property, and the framework
// disables the property on the other views.
// 
// The default is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/automaticallyReconfiguresDisplay
func (v VZVirtualMachineView) AutomaticallyReconfiguresDisplay() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("automaticallyReconfiguresDisplay"))
	return rv
}
func (v VZVirtualMachineView) SetAutomaticallyReconfiguresDisplay(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setAutomaticallyReconfiguresDisplay:"), value)
}



// A Boolean value that determines whether the system should send certain
// system keyboard shortcuts to the guest instead of the host.
//
// # Discussion
// 
// Defaults to `false.`
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/capturesSystemKeys
func (v VZVirtualMachineView) CapturesSystemKeys() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("capturesSystemKeys"))
	return rv
}
func (v VZVirtualMachineView) SetCapturesSystemKeys(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setCapturesSystemKeys:"), value)
}



// The VM to display in the view.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtualMachineView/virtualMachine
func (v VZVirtualMachineView) VirtualMachine() IVZVirtualMachine {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("virtualMachine"))
	return VZVirtualMachineFromID(objc.ID(rv))
}
func (v VZVirtualMachineView) SetVirtualMachine(value IVZVirtualMachine) {
	objc.Send[struct{}](v.ID, objc.Sel("setVirtualMachine:"), value)
}



































